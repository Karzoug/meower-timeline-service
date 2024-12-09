package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/cenkalti/backoff/v4"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/codes"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/entity"
	timelineApi "github.com/Karzoug/meower-timeline-service/pkg/proto/kafka/timeline/v1"

	"google.golang.org/protobuf/proto"
)

const (
	defaultOperationTimeout   = 5 * time.Second
	maxRetryTimeoutBeforeExit = 120 * time.Second

	preffixSpanName = "TimelineService.KafkaConsumer/"
)

func (c consumer) handler(ctx context.Context, msg *kafka.Message, logger zerolog.Logger) error {
	logger.Info().Msg("received message")

	event := &timelineApi.ChangeTaskEvent{}
	if err := proto.Unmarshal(msg.Value, event); err != nil {
		return fmt.Errorf("failed to deserialize payload: %w", err)
	}

	spanMethodName := preffixSpanName
	switch event.ChangeType {
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_POST_INSERT:
		spanMethodName += "postInsert"
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_POST_DELETE:
		spanMethodName += "postDelete"
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_DELETE:
		spanMethodName += "userDelete"
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_SUBSCRIBE:
		spanMethodName += "userSubscribe"
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_UNSUBSCRIBE:
		spanMethodName += "userUnsubscribe"
	default:
		return nil
	}

	ctx, span := c.tracer.Start(ctx, spanMethodName)
	defer span.End()

	var operation func() error
	switch event.ChangeType {
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_POST_INSERT:
		operation = c.buildPostInsertOperation(ctx, event, logger)
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_POST_DELETE:
		operation = c.buildPostDeleteOperation(ctx, event, logger)
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_DELETE:
		operation = c.buildUserDeleteOperation(ctx, event, logger)
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_SUBSCRIBE:
		operation = c.buildUserSubscribeOperation(ctx, event, logger)
	case timelineApi.ChangeTaskType_CHANGE_TASK_TYPE_USER_UNSUBSCRIBE:
		operation = c.buildUserUnsubscribeOperation(ctx, event, logger)
	}

	if err := backoff.Retry(operation,
		backoff.NewExponentialBackOff(
			backoff.WithMaxElapsedTime(maxRetryTimeoutBeforeExit),
		),
	); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "all operation retries failed")
		return err
	}

	logger.Info().Msg("processed message")

	return nil
}

func (c consumer) buildPostInsertOperation(ctx context.Context, event *timelineApi.ChangeTaskEvent, logger zerolog.Logger) func() error {
	const op = "build post insert operation"

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		targetUserID, err := xid.FromString(event.TargetUserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid target user id: %w", op, err))
		}
		userID, err := xid.FromString(event.UserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid user id: %w", op, err))
		}
		postID, err := xid.FromString(event.PostId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid post id: %w", op, err))
		}

		if err := c.timelineService.PushTimelinePost(ctx,
			targetUserID,
			entity.Post{
				AuthorID: userID,
				PostID:   postID,
			}); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Str("post_id", postID.String()).
					Err(serr.Unwrap()).
					Msg("push post to timeline failed")
			} else {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Str("post_id", postID.String()).
					Err(err).
					Msg("push post to timeline failed")
			}

			return err
		}

		return nil
	}
}

func (c consumer) buildPostDeleteOperation(ctx context.Context, event *timelineApi.ChangeTaskEvent, logger zerolog.Logger) func() error {
	const op = "build post delete operation"

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		targetUserID, err := xid.FromString(event.TargetUserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid target user id: %w", op, err))
		}
		userID, err := xid.FromString(event.UserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid user id: %w", op, err))
		}
		postID, err := xid.FromString(event.PostId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid post id: %w", op, err))
		}

		if err := c.timelineService.DeleteTimelinePost(ctx,
			targetUserID,
			entity.Post{
				AuthorID: userID,
				PostID:   postID,
			}); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Str("post_id", postID.String()).
					Err(serr.Unwrap()).
					Msg("delete post from timeline failed")
			} else {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Str("post_id", postID.String()).
					Err(err).
					Msg("delete post from timeline failed")
			}

			return err
		}

		return nil
	}
}

func (c consumer) buildUserDeleteOperation(ctx context.Context, event *timelineApi.ChangeTaskEvent, logger zerolog.Logger) func() error {
	const op = "build user delete operation"

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		targetUserID, err := xid.FromString(event.TargetUserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid target user id: %w", op, err))
		}

		if err := c.timelineService.DeleteTimeline(ctx, targetUserID); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Err(serr.Unwrap()).
					Msg("user home timeline delete failed")
			} else {
				logger.Warn().
					Str("target_user_id", event.TargetUserId).
					Err(err).
					Msg("user home timeline delete failed")
			}

			return err
		}

		return nil
	}
}

func (c consumer) buildUserSubscribeOperation(ctx context.Context, event *timelineApi.ChangeTaskEvent, logger zerolog.Logger) func() error {
	const op = "build user subscribe operation"

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		userID, err := xid.FromString(event.UserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid user id: %w", op, err))
		}
		targetUserID, err := xid.FromString(event.TargetUserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid target user id: %w", op, err))
		}

		if err := c.timelineService.SubscribeOnUser(ctx, userID, targetUserID); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				logger.Warn().
					Str("user_id", event.UserId).
					Str("target_user_id", event.TargetUserId).
					Err(serr.Unwrap()).
					Msg("subscribe on user failed")
			} else {
				logger.Warn().
					Str("user_id", event.UserId).
					Str("target_user_id", event.TargetUserId).
					Err(err).
					Msg("subscribe on user failed")
			}

			return err
		}

		return nil
	}
}

func (c consumer) buildUserUnsubscribeOperation(ctx context.Context, event *timelineApi.ChangeTaskEvent, logger zerolog.Logger) func() error {
	const op = "build user unsubscribe operation"

	return func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		userID, err := xid.FromString(event.UserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid user id: %w", op, err))
		}
		targetUserID, err := xid.FromString(event.TargetUserId)
		if err != nil {
			return backoff.Permanent(fmt.Errorf("%s: invalid target user id: %w", op, err))
		}

		if err := c.timelineService.UnsubscribeFromUser(ctx, userID, targetUserID); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				logger.Warn().
					Str("user_id", event.UserId).
					Str("target_user_id", event.TargetUserId).
					Err(serr.Unwrap()).
					Msg("unsubscribe from user failed")
			} else {
				logger.Warn().
					Str("user_id", event.UserId).
					Str("target_user_id", event.TargetUserId).
					Err(err).
					Msg("unsubscribe from user failed")
			}

			return err
		}

		return nil
	}
}
