package integration_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/artefactual-sdps/enduro/internal/a3m"
	"github.com/artefactual-sdps/enduro/internal/api"
	"github.com/artefactual-sdps/enduro/internal/api/auth"
	goastorage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	"github.com/artefactual-sdps/enduro/internal/config"
	"github.com/artefactual-sdps/enduro/internal/db"
	"github.com/artefactual-sdps/enduro/internal/event"
	"github.com/artefactual-sdps/enduro/internal/package_"
	"github.com/artefactual-sdps/enduro/internal/ref"
	"github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/temporal"
	"github.com/artefactual-sdps/enduro/internal/upload"
	"github.com/artefactual-sdps/enduro/internal/watcher"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/google/uuid"
	"github.com/otiai10/copy"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/history/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
	"golang.org/x/exp/slices"
	"gotest.tools/assert"
	"gotest.tools/fs"
)

var (
	// Path of the repository root.
	rootPath string

	// Path to sampledata
	sampledataPath string

	// Temporal namespace.
	temporalNamespace = getEnv("TEMPORAL_NAMESPACE", "default")

	// Temporal address.
	temporalAddress = getEnv("TEMPORAL_ADDRESS", "")
)

func getEnv(name, fallback string) string {
	const prefix = "ENDURO_INTEGRATION"
	v := os.Getenv(fmt.Sprintf("%s_%s", prefix, name))
	if v == "" {
		return fallback
	}
	return v
}

func init() {
	var err error
	rootPath, err = filepath.Abs("../")
	if err != nil {
		panic(err)
	}
}

type temporalInstance struct {
	client client.Client
	addr   string // Used when we're connected to a user-provided instance.
}

func setUpTemporal(t *testing.T) *temporalInstance {
	t.Helper()

	addr := temporalAddress

	// Dial user-provided Temporal server.
	if addr != "" {
		client, err := client.Dial(client.Options{
			HostPort:  addr,
			Namespace: temporalNamespace,
		})
		assert.NilError(t, err, "Failed to connect to user-provided Temporal server.")
		t.Logf("Using Temporal server listening on %q.", addr)
		return &temporalInstance{
			client: client,
			addr:   addr,
		}
	}

	// Fallback to development server provided by the Temporal GO SDK.
	ctx := context.Background()
	s, err := testsuite.StartDevServer(ctx, testsuite.DevServerOptions{
		LogLevel: "fatal",
	})
	assert.NilError(t, err, "Failed to start Temporal development server.")
	t.Cleanup(func() {
		s.Stop()
	})

	c := s.Client()
	t.Cleanup(func() {
		c.Close()
	})

	return &temporalInstance{
		client: c,
		addr:   s.FrontendHostPort(),
	}
}

func defaultLogger(t *testing.T) logr.Logger {
	t.Helper()

	return testr.NewWithOptions(t, testr.Options{
		LogTimestamp: false,
		Verbosity:    0,
	})
}

func defaultConfig(t *testing.T) *config.Configuration {
	t.Helper()

	// dbFile := fs.NewDir(t, "enduro_test").Join("test.db")

	return &config.Configuration{
		Verbosity: 0,
		Debug:     false,
		API: api.Config{
			Listen: "127.0.0.1:22333",
			Auth: auth.Config{
				Enabled: false,
			},
		},
		Event: event.Config{
			RedisAddress: "",
			RedisChannel: "",
		},
		Database: db.Config{
			DSN:     "",
			Migrate: true,
		},
		Temporal: temporal.Config{
			Address:   temporalAddress,
			Namespace: temporalNamespace,
			TaskQueue: "enduro",
		},
		Watcher: watcher.Config{
			Filesystem: []*watcher.FilesystemConfig{
				{
					Name: "SIP",
					Path: "",
				},
			},
		},
		Storage: storage.Config{
			Database: storage.Database{
				DSN: "",
			},
		},
		Upload: upload.Config{},
		A3m: a3m.Config{
			Address:  "",
			ShareDir: "",
		},
	}
}

type testEnvironment struct {
	cfg             *config.Configuration
	temporalClient  client.Client
	ingestBucketDir *fs.Dir
	dipsBucketDir   *fs.Dir
	aipsBucketDir   *fs.Dir
	token           string
}

func (e *testEnvironment) startWorkflow(t *testing.T) {
	t.Helper()

	ctx := context.Background()
	tdir := fs.NewDir(t, "enduro_integration_test")

	req := package_.ProcessingWorkflowRequest{
		WatcherName:                "integration_test",
		RetentionPeriod:            ref.New(time.Minute * 10),
		CompletedDir:               tdir.Join("completed/"),
		StripTopLevelDir:           false,
		Key:                        "test1234",
		IsDir:                      true,
		AutoApproveAIP:             true,
		DefaultPermanentLocationID: ref.New(uuid.New()),
	}
	err := package_.InitProcessingWorkflow(ctx, e.temporalClient, &req)
	assert.NilError(t, err, "Error initializing processing workflow.")
}

func (e *testEnvironment) submitSampleTransfer(t *testing.T, path string) string {
	t.Helper()

	path = filepath.Join(sampledataPath, path)
	items, err := os.ReadDir(path)
	assert.NilError(t, err)

	prefix := ""
	for _, item := range items {
		if strings.HasSuffix(item.Name(), ".isr.json") {
			prefix = strings.TrimSuffix(item.Name(), ".isr.json")
			break
		}
	}
	if prefix == "" {
		t.Fatalf("Cannot find ISR document in %q", path)
	}

	err = copy.Copy(path, e.ingestBucketDir.Path(), copy.Options{})
	assert.NilError(t, err, "Sample transfer %q could not be copied.", path)

	return prefix
}

func (e *testEnvironment) workflowMustCompleteWithErrors(t *testing.T, workflowID string) {
	ctx := context.Background()
	events := []*history.HistoryEvent{}
	iter := e.temporalClient.GetWorkflowHistory(ctx, workflowID, "", true, enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
	for iter.HasNext() {
		event, err := iter.Next()
		assert.NilError(t, err)
		events = append(events, event)
	}
	lastEvent := events[len(events)-1]

	if lastEvent.EventType != enums.EVENT_TYPE_WORKFLOW_EXECUTION_FAILED {
		t.Log("Workflow did not fail!")
		t.Fatal()
	}
}

func (e *testEnvironment) workflowMustCompleteWithoutErrors(t *testing.T, workflowID string) {
	ctx := context.Background()
	events := []*history.HistoryEvent{}
	iter := e.temporalClient.GetWorkflowHistory(ctx, workflowID, "", true, enums.HISTORY_EVENT_FILTER_TYPE_ALL_EVENT)
	for iter.HasNext() {
		event, err := iter.Next()
		assert.NilError(t, err)
		events = append(events, event)
	}
	lastEvent := events[len(events)-1]

	if lastEvent.EventType == enums.EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED {
		return
	}

	switch lastEvent.EventType {
	case enums.EVENT_TYPE_WORKFLOW_EXECUTION_FAILED:
		attrs := lastEvent.GetWorkflowExecutionFailedEventAttributes()
		t.Log("Workflow failed!")
		t.Logf("  - Failure: %s", attrs.Failure.Message)
		t.Logf("  - Info: %s", attrs.Failure.FailureInfo)
		t.Fatal()
	case enums.EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT:
		attrs := lastEvent.GetWorkflowExecutionTerminatedEventAttributes()
		t.Fatalf("Workflow was terminated: %s - %s", attrs.Details, attrs.Reason)
	default:
		attrs := lastEvent.Attributes
		t.Fatalf("Workflow completed unexpectedly: %s - %s", lastEvent.EventType, attrs)
	}
}

func (e *testEnvironment) expect(t *testing.T, method, path string, query map[string][]string, token string) *apitest.Response {
	url := fmt.Sprintf("http://%s%s", e.cfg.API.Listen, path)

	at := apitest.New().
		EnableNetworking(&http.Client{Timeout: time.Second * 1}).
		Intercept(func(r *http.Request) {
			r.Header.Set("Authorization", "Bearer "+token)
		})

	return at.Method(method).URL(url).QueryCollection(query).Expect(t)
}

func setUpMain(t *testing.T, temporalInstance *temporalInstance) *testEnvironment {
	t.Helper()

	env := &testEnvironment{
		cfg:             defaultConfig(t),
		ingestBucketDir: fs.NewDir(t, "ingest-bucket"),
		dipsBucketDir:   fs.NewDir(t, "dips-bucket"),
		aipsBucketDir:   fs.NewDir(t, "aips-bucket"),
		temporalClient:  temporalInstance.client,
	}

	env.cfg.Temporal.Address = temporalInstance.addr
	env.cfg.Watcher.Filesystem = []*watcher.FilesystemConfig{
		{
			Name: "",
			Path: "file://" + env.ingestBucketDir.Path(),
		},
	}
	env.cfg.API.CORSOrigin = "*"
	// env.token = token

	// ctx, cancel := context.WithCancel(context.Background())
	// m := workercmd.NewMain(defaultLogger(t), *env.cfg)

	return env
}

func requireSampledata(t *testing.T) {
	if sampledataPath == "" {
		t.Skip("sdps_postintegration_sampledata not indicated")
	}

	if !filepath.IsAbs(sampledataPath) {
		sampledataPath = filepath.Join(rootPath, sampledataPath)
	}
	if _, err := os.Stat(sampledataPath); os.IsNotExist(err) {
		t.Fatalf("sdps_postintegration_sampledata not found in: %q", sampledataPath)
	}
}

type storageService struct {
	t    testing.TB
	pkgs []*storageObject
	mu   sync.RWMutex
}

type storageObject struct {
	ID uuid.UUID
}

func newFakeStorageService(t testing.TB) *storageService {
	return &storageService{
		t: t,
	}
}

func (s *storageService) PackageRead(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())

	id, err := uuid.Parse(params["id"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	i := slices.IndexFunc(s.pkgs, func(p *storageObject) bool {
		return p.ID == id
	})
	if i == -1 {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	p := s.pkgs[i]

	blob, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(blob)
}

func (s *storageService) Submit(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	assert.NilError(s.t, err)

	var payload *goastorage.SubmitPayload
	assert.NilError(s.t, json.Unmarshal(body, &payload))
	assert.Assert(s.t, len(payload.AipID) > 0)

	s.mu.Lock()
	defer s.mu.Unlock()
	p := storageObject{
		ID: uuid.New(),
	}
	s.pkgs = append(s.pkgs, &p)
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"package_id": ` + p.ID.String() + `}`))
}

func fakeStorageService(t *testing.T) *httptest.Server {
	t.Helper()

	svc := newFakeStorageService(t)

	router := httptreemux.NewContextMux()
	router.GET("/ingests/:id", svc.PackageRead)
	router.POST("/ingests", svc.Submit)

	srv := httptest.NewServer(router)
	t.Cleanup(srv.Close)

	return srv
}

func TestIntegration(t *testing.T) {
	requireSampledata(t)
	temporalInstance := setUpTemporal(t)

	t.Run("Processes Bag successfully.", func(t *testing.T) {
		// fss := fakeStorageService(t)
		env := setUpMain(t, temporalInstance)

		// path := "ZippedBag.tgz"
		// transferPrefix := env.submitSampleTransfer(t, path)

		env.startWorkflow(t)

		// env.workflowMustCompleteWithoutErrors(t, workflowID)

		// Expect one batch.
		params := map[string][]string{"number": {"28-standard-EAS-A-non-encrypted"}}
		env.expect(t, http.MethodGet, "/contents/batches", params, env.token).
			Status(http.StatusOK).
			Assert(jsonpath.Equal("$.total_count", float64(1))).
			Assert(jsonpath.Len("$.items", 1)).
			End()
	})
}
