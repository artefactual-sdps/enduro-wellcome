package activities

import (
	"context"
	"io"
	"os"

	"github.com/go-logr/logr"

	"github.com/artefactual-sdps/enduro/internal/wellcome"
	"github.com/artefactual-sdps/enduro/internal/wellcome/auth"
	"github.com/artefactual-sdps/enduro/internal/wellcome/ingests"
)

const IngestActivityName = "wellcome-ingest-activity"

type IngestActivityParams struct {
	AIPID   string
	AIPPath string
	Key     string
}

type IngestActivity struct {
	logger logr.Logger
	cfg    wellcome.Config
}

func NewIngestActivity(logger logr.Logger, cfg wellcome.Config) *IngestActivity {
	return &IngestActivity{
		logger: logger,
		cfg:    cfg,
	}
}

func (a *IngestActivity) Execute(ctx context.Context, params *IngestActivityParams) error {
	if a.cfg.Upload == nil || !a.cfg.Upload.Enabled {
		a.logger.V(1).Info("Wellcome ingest is disabled, skipping.")
		return nil
	}

	a.logger.V(1).Info("Executing Wellcome ingest activity.")

	loc, err := copyAIP(ctx, params.AIPPath, params.Key, *a.cfg.Upload)
	if err != nil {
		return err
	}

	token, err := auth.GetAccessToken(
		a.cfg.Oauth.TokenURL,
		a.cfg.Oauth.GrantType,
		a.cfg.Oauth.ClientID,
		a.cfg.Oauth.SecretID,
	)
	if err != nil {
		return err
	}

	a.logger.V(1).Info("Got an access token.")

	ingest := ingests.NewService(a.cfg.Ingests)
	ingestID, err := ingest.Submit(token, loc, params.AIPID)
	if err != nil {
		return err
	}

	a.logger.V(1).Info("Started ingest", "ID", ingestID)

	return nil
}

func copyAIP(ctx context.Context, source, key string, dest wellcome.BucketConfig) (*ingests.AIPlocation, error) {
	reader, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer reader.Close() //#nosec G307 -- Errors returned by Close() here do not require specific handling.

	b, err := dest.OpenBucket(ctx)
	if err != nil {
		return nil, err
	}
	defer b.Close()

	writer, err := b.NewWriter(ctx, key, nil)
	if err != nil {
		return nil, err
	}

	_, copyErr := io.Copy(writer, reader)
	closeErr := writer.Close()

	if copyErr != nil {
		return nil, copyErr
	}
	if closeErr != nil {
		return nil, closeErr
	}

	loc := &ingests.AIPlocation{
		Bucket: dest,
		Key:    key,
	}

	return loc, nil
}
