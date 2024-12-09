package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	pkfk "github.com/Karzoug/meower-common-go/kafka"
	"github.com/Karzoug/meower-common-go/trace/otlp"

	zerologHook "github.com/Karzoug/meower-timeline-service/internal/delivery/kafka/zerolog"
	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	timelineApi "github.com/Karzoug/meower-timeline-service/pkg/proto/kafka/timeline/v1"
)

const (
	timelineTopic = "timelines"
)

var (
	topics                = []string{timelineTopic}
	changeTaskEventFngpnt = pkfk.MessageTypeHeaderValue(&timelineApi.ChangeTaskEvent{})
)

type consumer struct {
	c               *kafka.Consumer
	timelineService service.TimelineService
	tracer          trace.Tracer
	logger          zerolog.Logger
}

func NewConsumer(ctx context.Context, cfg Config, service service.TimelineService, tracer trace.Tracer, logger zerolog.Logger) (consumer, error) {
	const op = "create kafka consumer"

	logger = logger.With().
		Str("component", "kafka consumer").
		Logger()
	tracedLogger := logger.Hook(zerologHook.TraceIDHook())

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.Brokers,
		"group.id":                 cfg.GroupID,
		"auto.offset.reset":        "latest",
		"auto.commit.interval.ms":  cfg.CommitIntervalMilliseconds,
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return consumer{}, fmt.Errorf("%s: failed to create consumer: %w", op, err)
	}

	var (
		timeout int
		topic   = timelineTopic
	)
	if t, ok := ctx.Deadline(); ok {
		timeout = int(time.Until(t).Milliseconds())
	} else {
		timeout = 500
	}

	// analog PING here
	_, err = c.GetMetadata(&topic, false, timeout)
	if err != nil {
		return consumer{}, fmt.Errorf("%s: failed to get metadata: %w", op, err)
	}

	return consumer{
		c:               c,
		timelineService: service,
		tracer:          tracer,
		logger:          tracedLogger,
	}, nil
}

func (c consumer) Run(ctx context.Context) (err error) {
	const op = "run kafka consumer"

	defer func() {
		if defErr := c.c.Close(); defErr != nil {
			err = errors.Join(err,
				fmt.Errorf("%s: failed to close consumer: %w", op, defErr))
		}
	}()

	if err := c.c.SubscribeTopics(topics, nil); err != nil {
		return err
	}

	run := true
	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			msg, err := c.c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				var kafkaErr kafka.Error
				if errors.As(err, &kafkaErr) {
					if kafkaErr.IsFatal() {
						return fmt.Errorf("%s: fatal error while read message: %w", op, err)
					}
					if !kafkaErr.IsTimeout() {
						c.logger.Error().
							Err(err).
							Msg("failed to read message")
					}
				}
				continue
			}

			if len(msg.Headers) == 0 {
				c.storeOffset(msg)
				continue
			}
			eventType, ok := lookupHeaderValue(msg.Headers, pkfk.MessageTypeHeaderKey)
			if !ok {
				c.storeOffset(msg)
				continue
			}
			eventTypeFngpnt := string(eventType)

			ctx = otlp.InjectTracing(ctx, c.tracer)
			hlogger := c.logger.With().
				Str("topic", *msg.TopicPartition.Topic).
				Str("key", string(msg.Key)).
				Str("event fingerprint", eventTypeFngpnt).
				Ctx(ctx).
				Logger()

			if eventTypeFngpnt == changeTaskEventFngpnt {
				err = c.handler(ctx, msg, hlogger)
			}

			if err != nil {
				// log, not store offset, return from consumer with error
				return err
			}

			c.storeOffset(msg)
		}
	}

	return nil
}

func (c consumer) storeOffset(msg *kafka.Message) {
	_, err := c.c.StoreMessage(msg)
	if err != nil {
		c.logger.Error().
			Err(err).
			Str("topic", *msg.TopicPartition.Topic).
			Str("key", string(msg.Key)).
			Msg("failed to store offset after message")
	}
}

func lookupHeaderValue(headers []kafka.Header, key string) ([]byte, bool) {
	for _, header := range headers {
		if header.Key == key {
			return header.Value, true
		}
	}
	return nil, false
}
