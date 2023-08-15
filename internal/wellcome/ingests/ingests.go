package ingests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/artefactual-sdps/enduro/internal/wellcome"
	"github.com/artefactual-sdps/enduro/internal/wellcome/auth"
)

type IngestID string

type AIPlocation struct {
	Bucket wellcome.BucketConfig
	Key    string
}

type Service struct {
	cfg *wellcome.IngestsConfig
}

func NewService(cfg *wellcome.IngestsConfig) *Service {
	return &Service{cfg: cfg}
}

// Example Wellcome storage service request for creating a new ingest
//
// curl -X POST "$API_URL/ingests" \
// --header "Authorization: $ACCESS_TOKEN" \
// --header "Content-Type: application/json" \
// --data "{
//     \"type\": \"Ingest\",
//     \"ingestType\": {\"id\": \"$INGEST_TYPE\", \"type\": \"IngestType\"},
//     \"space\": {\"id\": \"$SPACE\", \"type\": \"Space\"},
//     \"sourceLocation\": {
//         \"provider\": {\"id\": \"amazon-s3\", \"type\": \"Provider\"},
//         \"bucket\": \"$UPLOADS_BUCKET\",
//         \"path\": \"$UPLOADED_BAG_KEY\",
//         \"type\": \"Location\"
//     },
//     \"bag\": {
//         \"info\": {
//             \"externalIdentifier\": \"$EXTERNAL_IDENTIFIER\",
//             \"type\": \"BagInfo\"
//         },
//         \"type\": \"Bag\"
//     }
// }"

// IngestData represents the ingests POST request data.
type IngestData struct {
	Type           string         `json:"type"`       // e.g. "Ingest",
	IngestType     IDType         `json:"ingestType"` // e.g. {"id": "$INGEST_TYPE", "type": "IngestType"}
	Space          IDType         `json:"space"`      // e.g. {"id": "$SPACE", "type": "Space"}
	SourceLocation SourceLocation `json:"sourceLocation"`
	Bag            BagData        `json:"bag"`
}

type IDType struct {
	ID   string `json:"id"`
	Type string `json:"type"` // e.g. "IngestType"
}

type SourceLocation struct {
	Provider IDType `json:"provider"` // e.g. {"id": "amazon-s3", "type": "Provider"}
	Bucket   string `json:"bucket"`   // S3 bucket URL
	Path     string `json:"path"`     // blob key (aka filename)
	Type     string `json:"type"`     // e.g. "Location"
}

type BagData struct {
	Info BagInfo `json:"info"`
	Type string  `json:"type"` // e.g. "Bag"
}

type BagInfo struct {
	ExternalID string `json:"externalIdentifier"` // AIP UUID from a3m
	Type       string `json:"type"`               // e.g. "BagInfo"
}

func (svc *Service) Submit(token *auth.AccessToken, aipLocation *AIPlocation, bagExternalID string) (IngestID, error) {
	bucketURL, err := aipLocation.Bucket.ToURL()
	if err != nil {
		return "", err
	}

	data := &IngestData{
		Type: "Ingest",
		IngestType: IDType{
			ID:   svc.cfg.Type,
			Type: "IngestType",
		},
		Space: IDType{
			ID:   svc.cfg.SpaceID,
			Type: "Space",
		},
		SourceLocation: SourceLocation{
			Provider: IDType{
				ID:   "amazon-s3",
				Type: "Provider",
			},
			Bucket: bucketURL,
			Path:   aipLocation.Key,
			Type:   "Location",
		},
		Bag: BagData{
			Info: BagInfo{
				ExternalID: bagExternalID,
				Type:       "BagInfo",
			},
			Type: "Bag",
		},
	}

	// Encode the request data as JSON.
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		return "", fmt.Errorf("ingest bag: couldn't encode ingest data:  %v", err)
	}

	// Create a POST /ingests request.
	req, err := http.NewRequest(http.MethodPost, svc.cfg.ApiURL, buf)
	if err != nil {
		return "", fmt.Errorf("ingest bag: couldn't create an HTTP request:  %v", err)
	}

	// Set the request headers.
	req.Header.Set("Authorization", token.Value)
	req.Header.Set("Content-Type", "application/json")

	// Do the POST /ingests request.
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("ingest bag: couldn't send request: %v\n", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("ingest bag: couldn't read response: %v\n", err)
	}
	res.Body.Close()

	// Expect a "201 Created" response status.
	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("ingest bag: unexpected response: status: %q, body: %q", res.Status, body)
	}

	return ingestID(res.Header)
}

func ingestID(h http.Header) (IngestID, error) {
	loc := h.Get("Location")
	if loc == "" {
		return "", errors.New(`ingest bag: missing "Location" response header`)
	}

	id, ok := strings.CutPrefix(loc, "/ingests/")
	if !ok {
		return "", errors.New("ingest bag: missing ingest ID in response")
	}

	return IngestID(id), nil
}
