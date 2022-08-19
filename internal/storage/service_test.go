package storage_test

import (
	"context"
	"errors"
	"io"
	"net/url"
	"testing"

	"github.com/go-logr/logr"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	temporalsdk_api_enums "go.temporal.io/api/enums/v1"
	temporalsdk_client "go.temporal.io/sdk/client"
	temporalsdk_mocks "go.temporal.io/sdk/mocks"
	goa "goa.design/goa/v3/pkg"
	"gocloud.dev/blob"
	"gocloud.dev/blob/fileblob"
	"gocloud.dev/blob/memblob"
	"gotest.tools/v3/assert"

	goastorage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	"github.com/artefactual-sdps/enduro/internal/ref"
	"github.com/artefactual-sdps/enduro/internal/storage"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence"
	"github.com/artefactual-sdps/enduro/internal/storage/persistence/fake"
	"github.com/artefactual-sdps/enduro/internal/storage/status"
)

type setUpAttrs struct {
	logger         *logr.Logger
	config         *storage.Config
	persistence    *persistence.Storage
	temporalClient *temporalsdk_client.Client

	persistenceMock    *fake.MockStorage
	temporalClientMock *temporalsdk_mocks.Client
}

func setUpService(t *testing.T, attrs *setUpAttrs) storage.Service {
	t.Helper()

	psMock := fake.NewMockStorage(gomock.NewController(t))
	var ps persistence.Storage = psMock

	tcMock := &temporalsdk_mocks.Client{}
	var tc temporalsdk_client.Client = tcMock

	params := setUpAttrs{
		logger: ref.New(logr.Discard()),
		config: ref.New(storage.Config{
			Internal: storage.LocationConfig{
				Name:   "",
				Bucket: "internal",
				Region: "eu-west-2",
			},
			Locations: []storage.LocationConfig{
				{
					Name:   "perma-aips-1",
					Bucket: "perma-aips-1",
					Region: "eu-west-2",
				},
			},
		}),
		persistence:        &ps,
		persistenceMock:    psMock,
		temporalClient:     &tc,
		temporalClientMock: tcMock,
	}
	if attrs.logger != nil {
		params.logger = attrs.logger
	}
	if attrs.config != nil {
		params.config = attrs.config
	}
	if attrs.persistence != nil {
		params.persistence = attrs.persistence
	}
	if attrs.temporalClient != nil {
		params.temporalClient = attrs.temporalClient
	}

	*attrs = params

	s, err := storage.NewService(
		*params.logger,
		*params.config,
		*params.persistence,
		*params.temporalClient,
	)
	assert.NilError(t, err)

	return s
}

func fakeInternalLocation(t *testing.T, svc storage.Service, b *blob.Bucket) {
	t.Helper()

	t.Cleanup(func() { b.Close() })

	l, err := svc.Location("")
	assert.NilError(t, err)

	if b == nil {
		b = memblob.OpenBucket(nil)
	}

	l.SetBucket(b)
}

func fakeLocation(t *testing.T, svc storage.Service, name, objectKey, contents string) {
	t.Helper()

	l, err := svc.Location(name)
	assert.NilError(t, err)

	mb := memblob.OpenBucket(nil)
	t.Cleanup(func() { mb.Close() })
	l.SetBucket(mb)

	mb.WriteAll(context.Background(), objectKey, []byte(contents), nil)
}

func TestNewService(t *testing.T) {
	t.Parallel()

	_, err := storage.NewService(
		logr.Discard(),
		storage.Config{},
		nil,
		nil,
	)

	assert.ErrorContains(t, err, "s3blob.OpenBucket: bucketName is required")
}

func TestServiceSubmit(t *testing.T) {
	t.Parallel()

	t.Run("Returns not_valid if AIPID is invalid", func(t *testing.T) {
		t.Parallel()

		AIPID := "12345"
		attrs := &setUpAttrs{}
		svc := setUpService(t, attrs)

		ret, err := svc.Submit(context.Background(), &goastorage.SubmitPayload{
			AipID: AIPID,
		})
		assert.Assert(t, ret == nil)
		assert.Equal(t, err.(*goa.ServiceError).Name, "not_valid")
		assert.ErrorContains(t, err, "invalid UUID length: 5")
	})

	t.Run("Returns not_available if workflow cannot be executed", func(t *testing.T) {
		t.Parallel()

		AIPID := "5ab42bc3-acc2-420b-bbd0-76efdef94828"
		attrs := &setUpAttrs{}
		svc := setUpService(t, attrs)

		attrs.temporalClientMock.
			On(
				"ExecuteWorkflow",
				mock.AnythingOfType("*context.timerCtx"),
				temporalsdk_client.StartWorkflowOptions{
					ID:                    "storage-upload-workflow-" + AIPID,
					TaskQueue:             "global",
					WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
				},
				"storage-upload-workflow",
				&storage.StorageUploadWorkflowRequest{AIPID: AIPID},
			).
			Return(
				nil,
				errors.New("something went wrong"),
			).
			Times(1)

		ret, err := svc.Submit(context.Background(), &goastorage.SubmitPayload{
			AipID: AIPID,
		})
		assert.Assert(t, ret == nil)
		assert.Equal(t, err.(*goa.ServiceError).Name, "not_available")
		assert.ErrorContains(t, err, "cannot perform operation")
	})

	t.Run("Returns not_valid if package cannot be persisted", func(t *testing.T) {
		t.Parallel()

		AIPID := "5ab42bc3-acc2-420b-bbd0-76efdef94828"
		attrs := &setUpAttrs{}
		svc := setUpService(t, attrs)
		ctx := context.Background()

		attrs.temporalClientMock.
			On(
				"ExecuteWorkflow",
				mock.AnythingOfType("*context.timerCtx"),
				temporalsdk_client.StartWorkflowOptions{
					ID:                    "storage-upload-workflow-" + AIPID,
					TaskQueue:             "global",
					WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
				},
				"storage-upload-workflow",
				&storage.StorageUploadWorkflowRequest{AIPID: AIPID},
			).
			Return(
				&temporalsdk_mocks.WorkflowRun{},
				nil,
			).
			Times(1)

		attrs.persistenceMock.
			EXPECT().
			CreatePackage(
				gomock.AssignableToTypeOf(ctx),
				"package",
				uuid.MustParse(AIPID),
				gomock.Any(),
			).
			Return(
				nil,
				errors.New("database server error"),
			).
			Times(1)

		ret, err := svc.Submit(ctx, &goastorage.SubmitPayload{
			Name:  "package",
			AipID: AIPID,
		})
		assert.Assert(t, ret == nil)
		assert.Equal(t, err.(*goa.ServiceError).Name, "not_valid")
		assert.ErrorContains(t, err, "cannot persist package")
	})

	t.Run("Returns not_valid if signed URL cannot be generated", func(t *testing.T) {
		t.Parallel()

		AIPID := "5ab42bc3-acc2-420b-bbd0-76efdef94828"
		attrs := &setUpAttrs{}
		svc := setUpService(t, attrs)
		ctx := context.Background()

		attrs.temporalClientMock.
			On(
				"ExecuteWorkflow",
				mock.AnythingOfType("*context.timerCtx"),
				temporalsdk_client.StartWorkflowOptions{
					ID:                    "storage-upload-workflow-" + AIPID,
					TaskQueue:             "global",
					WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
				},
				"storage-upload-workflow",
				&storage.StorageUploadWorkflowRequest{AIPID: AIPID},
			).
			Return(
				&temporalsdk_mocks.WorkflowRun{},
				nil,
			).
			Times(1)

		attrs.persistenceMock.
			EXPECT().
			CreatePackage(
				gomock.AssignableToTypeOf(ctx),
				"package",
				uuid.MustParse(AIPID),
				func() gomock.Matcher {
					return gomock.Any()
				}(),
			).
			Return(
				&goastorage.StoredStoragePackage{},
				nil,
			).
			Times(1)

		fakeInternalLocation(t, svc, nil)

		ret, err := svc.Submit(ctx, &goastorage.SubmitPayload{
			Name:  "package",
			AipID: AIPID,
		})
		assert.Assert(t, ret == nil)
		assert.Equal(t, err.(*goa.ServiceError).Name, "not_valid")
		assert.ErrorContains(t, err, "cannot persist package")
	})

	t.Run("Returns signed URL", func(t *testing.T) {
		t.Parallel()

		AIPID := "5ab42bc3-acc2-420b-bbd0-76efdef94828"
		attrs := &setUpAttrs{}
		svc := setUpService(t, attrs)
		ctx := context.Background()

		attrs.temporalClientMock.
			On(
				"ExecuteWorkflow",
				mock.AnythingOfType("*context.timerCtx"),
				temporalsdk_client.StartWorkflowOptions{
					ID:                    "storage-upload-workflow-" + AIPID,
					TaskQueue:             "global",
					WorkflowIDReusePolicy: temporalsdk_api_enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
				},
				"storage-upload-workflow",
				&storage.StorageUploadWorkflowRequest{AIPID: AIPID},
			).
			Return(
				&temporalsdk_mocks.WorkflowRun{},
				nil,
			).
			Times(1)

		attrs.persistenceMock.
			EXPECT().
			CreatePackage(
				gomock.AssignableToTypeOf(ctx),
				"package",
				uuid.MustParse(AIPID),
				func() gomock.Matcher {
					return gomock.Any()
				}(),
			).
			Return(
				&goastorage.StoredStoragePackage{},
				nil,
			).
			Times(1)

		// Fake internal location, using fileblob because it can generate signed URLs.
		furl, err := url.Parse("file:///tmp/dir")
		assert.NilError(t, err)
		b, err := fileblob.OpenBucket("/tmp", &fileblob.Options{URLSigner: fileblob.NewURLSignerHMAC(furl, []byte("1234"))})
		assert.NilError(t, err)
		fakeInternalLocation(t, svc, b)

		ret, err := svc.Submit(ctx, &goastorage.SubmitPayload{
			Name:  "package",
			AipID: AIPID,
		})
		assert.Equal(t, ret.URL[0:15], "file:///tmp/dir")
		assert.NilError(t, err)
	})
}

func TestServiceLocation(t *testing.T) {
	t.Parallel()

	svc := setUpService(t, &setUpAttrs{})

	testCases := map[string]struct {
		name string
		err  error
	}{
		"Returns internal location": {
			"",
			nil,
		},
		"Returns location": {
			"perma-aips-1",
			nil,
		},
		"Returns error when location cannot be found": {
			"perma-aips-999",
			errors.New("error loading location: unknown location perma-aips-999"),
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			loc, err := svc.Location(tc.name)

			if tc.err == nil {
				assert.NilError(t, err)
				assert.Equal(t, loc.Name(), tc.name)
			} else {
				assert.Error(t, err, tc.err.Error())
			}
		})
	}
}

func TestServiceDownload(t *testing.T) {
	t.Parallel()

	svc := setUpService(t, &setUpAttrs{})

	blob, err := svc.Download(context.Background(), &goastorage.DownloadPayload{})
	assert.NilError(t, err)
	assert.DeepEqual(t, blob, []byte{}) // Not implemented.
}

func TestServiceList(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	t.Run("Returns defined locations", func(t *testing.T) {
		t.Parallel()

		svc := setUpService(t, &setUpAttrs{})

		res, err := svc.Locations(ctx)
		assert.NilError(t, err)
		assert.DeepEqual(t, res, goastorage.StoredLocationCollection{
			{
				ID:          1,
				Name:        "perma-aips-1",
				Description: ref.New(""),
				Source:      "minio",
				Purpose:     "aip_store",
				UUID:        ref.New(""),
			},
		})
	})
}

func TestReject(t *testing.T) {
	t.Run("Fails when passing an invalid UUID", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		attrs.persistenceMock.
			EXPECT().
			UpdatePackageStatus(
				ctx,
				status.StatusRejected,
				uuid.MustParse(AIPID),
			).
			Return(nil).
			Times(1)

		err := svc.Reject(ctx, &goastorage.RejectPayload{AipID: AIPID})
		assert.NilError(t, err)
	})
}

func TestServiceReadPackage(t *testing.T) {
	t.Parallel()

	t.Run("Fails when passing an invalid UUID", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()

		pkg, err := svc.ReadPackage(ctx, "<invalid-uuid>")
		assert.Error(t, err, "invalid UUID length: 14")
		assert.Assert(t, pkg == nil)
	})
}

func TestServiceUpdatePackageStatus(t *testing.T) {
	t.Parallel()

	t.Run("Fails when passing an invalid UUID", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()

		err := svc.UpdatePackageStatus(ctx, status.StatusStored, "<invalid-uuid>")
		assert.Error(t, err, "invalid UUID length: 14")
	})

	t.Run("Returns if persistence failed", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		attrs.persistenceMock.
			EXPECT().
			UpdatePackageStatus(
				ctx,
				status.StatusStored,
				uuid.MustParse(AIPID),
			).
			Return(errors.New("something is wrong")).
			Times(1)

		err := svc.UpdatePackageStatus(ctx, status.StatusStored, AIPID)
		assert.Error(t, err, "something is wrong")
	})
}

func TestServiceUpdatePackageLocation(t *testing.T) {
	t.Parallel()

	t.Run("Fails when passing an invalid UUID", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()

		err := svc.UpdatePackageLocation(ctx, "perma-aips-1", "<invalid-uuid>")
		assert.Error(t, err, "invalid UUID length: 14")
	})

	t.Run("Returns if persistence failed", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"
		loc := "perma-aips-1"

		attrs.persistenceMock.
			EXPECT().
			UpdatePackageLocation(
				ctx,
				loc,
				uuid.MustParse(AIPID),
			).
			Return(errors.New("something is wrong")).
			Times(1)

		err := svc.UpdatePackageLocation(ctx, "perma-aips-1", AIPID)
		assert.Error(t, err, "something is wrong")
	})
}

func TestServiceDelete(t *testing.T) {
	t.Parallel()

	t.Run("From internal location", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		attrs.persistenceMock.
			EXPECT().
			ReadPackage(
				ctx,
				uuid.MustParse(AIPID),
			).
			Return(
				&goastorage.StoredStoragePackage{
					ID:        1,
					AipID:     AIPID,
					ObjectKey: AIPID,
					Location:  nil,
				},
				nil,
			).
			Times(1)

		fakeLocation(t, svc, "", AIPID, "foobar")

		err := svc.Delete(ctx, AIPID)
		assert.NilError(t, err)
	})

	t.Run("From perma location", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "76a654ad-dccc-4dd3-a398-e84cd9f96415"

		attrs.persistenceMock.
			EXPECT().
			ReadPackage(
				ctx,
				uuid.MustParse(AIPID),
			).
			Return(
				&goastorage.StoredStoragePackage{
					ID:        1,
					AipID:     AIPID,
					ObjectKey: AIPID,
					Location:  ref.New("perma-aips-1"),
				},
				nil,
			).
			Times(1)

		fakeLocation(t, svc, "perma-aips-1", AIPID, "foobar")

		err := svc.Delete(ctx, AIPID)
		assert.NilError(t, err)
	})

	t.Run("Fails if object does not exist", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "76a654ad-dccc-4dd3-a398-e84cd9f96415"

		attrs.persistenceMock.
			EXPECT().
			ReadPackage(
				ctx,
				uuid.MustParse(AIPID),
			).
			Return(
				&goastorage.StoredStoragePackage{
					ID:        1,
					AipID:     AIPID,
					ObjectKey: AIPID,
					Location:  ref.New("perma-aips-1"),
				},
				nil,
			).
			Times(1)

		// Fake empty location.
		l, err := svc.Location("perma-aips-1")
		assert.NilError(t, err)
		mb := memblob.OpenBucket(nil)
		l.SetBucket(mb)

		err = svc.Delete(ctx, AIPID)
		assert.Error(t, err, "blob (key \"76a654ad-dccc-4dd3-a398-e84cd9f96415\") (code=NotFound): blob not found")
	})

	t.Run("Fails if location is not available", func(t *testing.T) {
		t.Parallel()

		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "76a654ad-dccc-4dd3-a398-e84cd9f96415"

		attrs.persistenceMock.
			EXPECT().
			ReadPackage(
				ctx,
				uuid.MustParse(AIPID),
			).
			Return(
				&goastorage.StoredStoragePackage{
					ID:        1,
					AipID:     AIPID,
					ObjectKey: AIPID,
					Location:  ref.New("perma-aips-99"),
				},
				nil,
			).
			Times(1)

		err := svc.Delete(ctx, AIPID)
		assert.Error(t, err, "error loading location: unknown location perma-aips-99")
	})
}

func TestPackageReader(t *testing.T) {
	t.Parallel()

	t.Run("Provides a valid reader", func(t *testing.T) {
		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		fakeLocation(t, svc, "perma-aips-1", AIPID, "contents")

		reader, err := svc.PackageReader(ctx, &goastorage.StoredStoragePackage{
			ID:        1,
			AipID:     AIPID,
			ObjectKey: AIPID,
			Location:  ref.New("perma-aips-1"),
		})
		assert.NilError(t, err)

		blob, err := io.ReadAll(reader)
		assert.NilError(t, err)
		assert.Equal(t, string(blob), "contents")
	})

	t.Run("Fails if the location is not available", func(t *testing.T) {
		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		_, err := svc.PackageReader(ctx, &goastorage.StoredStoragePackage{
			ID:        1,
			AipID:     AIPID,
			ObjectKey: AIPID,
			Location:  ref.New("perma-aips-99"),
		})
		assert.Error(t, err, "error loading location: unknown location perma-aips-99")
	})

	t.Run("Fails if the reader cannot be created", func(t *testing.T) {
		attrs := setUpAttrs{}
		svc := setUpService(t, &attrs)
		ctx := context.Background()
		AIPID := "7c09fa45-cdac-4874-90af-56dc86a6e73c"

		// Close the bucker beforehand to force the error.
		l, err := svc.Location("perma-aips-1")
		assert.NilError(t, err)
		l.Bucket().Close()

		_, err = svc.PackageReader(ctx, &goastorage.StoredStoragePackage{
			ID:        1,
			AipID:     AIPID,
			ObjectKey: AIPID,
			Location:  ref.New("perma-aips-1"),
		})
		assert.Error(t, err, "blob: Bucket has been closed (code=FailedPrecondition)")
	})
}