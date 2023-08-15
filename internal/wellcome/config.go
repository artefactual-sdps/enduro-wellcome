package wellcome

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws/session"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
)

type Config struct {
	Upload  *BucketConfig
	Oauth   *OAuth2Config
	Ingests *IngestsConfig
}

type BucketConfig struct {
	// Enabled toggles the Wellcome upload functionality on (true) or off (false).
	Enabled bool

	// URL specifies the location's driver and address by URL (e.g.
	// "s3://my-bucket?region=us-west-1", "file:///tmp/my-bucket").
	URL string

	// S3 compatible location configuration. If URL has a value then these
	// fields are ignored.
	Name      string
	Region    string
	Endpoint  string
	PathStyle bool
	Profile   string
	Key       string
	Secret    string
	Token     string
	Bucket    string
}

type OAuth2Config struct {
	TokenURL  string // e.g. "http://localhost:8002"
	GrantType string // e.g. "client_credentials"
	ClientID  string // e.g. "enduro"
	SecretID  string // e.g. "d18a98dc7241e4e4542bf79296c87146f064df50ba02d885e38c2bd982f2b67d"
}

type IngestsConfig struct {
	// The URL of the Wellcome ingests API endpoint, e.g.
	// http://localhost/ingests
	ApiURL string

	// API docs: "ingestType -- If this is the first bag with this external
	// identifier in this space, use create. Otherwise, use update."
	Type string

	// API docs: "The broad category of a bag, e.g. digitised, born-digital."
	SpaceID string
}

// ToURL returns a URL representation of a bucket location.
func (c BucketConfig) ToURL() (string, error) {
	if c.URL != "" {
		return c.URL, nil
	}

	// TODO: The URL probably needs more information (e.g. region, user?)
	l, err := url.JoinPath(c.Endpoint, c.Name)
	if err != nil {
		return "", fmt.Errorf("invalid bucket config: endpoint: %q, name: %q", c.Endpoint, c.Name)
	}

	return l, nil
}

func (c BucketConfig) OpenBucket(ctx context.Context) (*blob.Bucket, error) {
	if c.URL != "" {
		return c.openURLBucket(ctx)
	}

	return c.openS3Bucket(ctx)
}

func (c BucketConfig) openURLBucket(ctx context.Context) (*blob.Bucket, error) {
	b, err := blob.OpenBucket(ctx, c.URL)
	if err != nil {
		return nil, fmt.Errorf("open bucket by URL: %v", err)
	}

	return b, nil
}

func (c BucketConfig) openS3Bucket(ctx context.Context) (*blob.Bucket, error) {
	opts := session.Options{
		Profile: c.Profile,
		// Force enable Shared Config support
		SharedConfigState: session.SharedConfigEnable,
	}

	opts.Config.WithRegion(c.Region)
	opts.Config.WithEndpoint(c.Endpoint)
	opts.Config.WithS3ForcePathStyle(c.PathStyle)
	sess, err := session.NewSessionWithOptions(opts)
	if err != nil {
		return nil, err
	}

	// Create a *blob.Bucket.
	bucket, err := s3blob.OpenBucket(ctx, sess, c.Bucket, nil)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
