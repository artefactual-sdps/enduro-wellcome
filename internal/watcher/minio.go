package watcher

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-logr/logr"
	"github.com/redis/go-redis/v9"
	"gocloud.dev/blob"
	"gocloud.dev/blob/s3blob"
)

// minioWatcher implements a Watcher for watching lists in Redis.
type minioWatcher struct {
	client     redis.UniversalClient
	sess       *session.Session
	logger     logr.Logger
	listName   string
	failedList string
	bucket     string
	*commonWatcherImpl
}

type MinioEventSet struct {
	Event     []MinioEvent
	EventTime string
}

var _ Watcher = (*minioWatcher)(nil)

const redisPopTimeout = time.Second * 2

func NewMinioWatcher(ctx context.Context, logger logr.Logger, config *MinioConfig) (*minioWatcher, error) {
	opts, err := redis.ParseURL(config.RedisAddress)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	sessOpts := session.Options{}
	sessOpts.Config.WithRegion(config.Region)
	if config.Profile != "" {
		sessOpts.Profile = config.Profile
	}
	if config.Endpoint != "" {
		sessOpts.Config.WithEndpoint(config.Endpoint)
	}
	if config.PathStyle {
		sessOpts.Config.WithS3ForcePathStyle(config.PathStyle)
	}
	if config.Key != "" && config.Secret != "" {
		sessOpts.Config.WithCredentials(credentials.NewStaticCredentials(config.Key, config.Secret, config.Token))
	}
	sess, err := session.NewSessionWithOptions(sessOpts)
	if err != nil {
		return nil, err
	}
	if config.RedisFailedList == "" {
		config.RedisFailedList = config.RedisList + "-failed"
	}

	return &minioWatcher{
		client:     client,
		sess:       sess,
		listName:   config.RedisList,
		failedList: config.RedisFailedList,
		logger:     logger,
		bucket:     config.Bucket,
		commonWatcherImpl: &commonWatcherImpl{
			name:             config.Name,
			retentionPeriod:  config.RetentionPeriod,
			stripTopLevelDir: config.StripTopLevelDir,
		},
	}, nil
}

func (w *minioWatcher) Watch(ctx context.Context) (*BlobEvent, Cleanup, error) {
	event, val, err := w.pop(ctx)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = ErrWatchTimeout
		}
		return nil, noopCleanup, err
	}
	if event.Bucket != w.bucket {
		return nil, noopCleanup, ErrBucketMismatch
	}
	cleanup := w.rem(val)

	return event, cleanup, nil
}

func (w *minioWatcher) Path() string {
	return ""
}

// rem return a function that allows items from the failed list to be safely removed
// in the event of a workflow or some other type of redis error.
func (w *minioWatcher) rem(val string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		logger := w.logger.WithValues("list", w.failedList)
		if _, err := w.client.LRem(ctx, w.failedList, 1, val).Result(); err != nil {
			logger.Error(err, "Error removing message from failed list.")
			return err
		}
		logger.V(2).Info("Successfully removed message(s).")

		return nil
	}
}

func (w *minioWatcher) pop(ctx context.Context) (*BlobEvent, string, error) {
	val, err := w.client.BLMove(ctx, w.listName, w.failedList, "RIGHT", "LEFT", redisPopTimeout).Result()
	if err != nil {
		return nil, "", fmt.Errorf("error retrieving from Redis list: %w", err)
	}

	event, err := w.event(val)
	if err != nil {
		return nil, "", fmt.Errorf("error processing item received: %w", err)
	}

	return event, val, nil
}

// event processes Minio-specific events delivered via Redis. We expect a
// single item array containing a map of {Event: ..., EvenTime: ...}
func (w *minioWatcher) event(blob string) (*BlobEvent, error) {
	container := []json.RawMessage{}
	if err := json.Unmarshal([]byte(blob), &container); err != nil {
		return nil, err
	}

	var set MinioEventSet
	if err := json.Unmarshal(container[0], &set); err != nil {
		return nil, fmt.Errorf("error procesing item received from Redis list: %w", err)
	}
	if len(set.Event) == 0 {
		return nil, fmt.Errorf("error processing item received from Redis list: empty event list")
	}

	key, err := url.QueryUnescape(set.Event[0].S3.Object.Key)
	if err != nil {
		return nil, fmt.Errorf("error processing item received from Redis list: %w", err)
	}

	return NewBlobEventWithBucket(w, set.Event[0].S3.Bucket.Name, key), nil
}

func (w *minioWatcher) OpenBucket(ctx context.Context) (*blob.Bucket, error) {
	return s3blob.OpenBucket(ctx, w.sess, w.bucket, nil)
}
