package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"

	"github.com/Karzoug/meower-timeline-service/internal/timeline/service"
	pkfk "github.com/Karzoug/meower-timeline-service/pkg/kafka"
)

type consumer struct {
	c               *kafka.Consumer
	topics          []string
	timelineService service.TimelineService
	logger          zerolog.Logger
}

func NewConsumer(ctx context.Context, cfg Config, service service.TimelineService, logger zerolog.Logger) (consumer, error) {
	const op = "create kafka consumer:"

	logger = logger.With().
		Str("component", "kafka consumer").
		Logger()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.Brokers,
		"group.id":                 cfg.GroupID,
		"auto.offset.reset":        "earliest",
		"auto.commit.interval.ms":  cfg.CommitIntervalMilliseconds,
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return consumer{}, fmt.Errorf("%s failed to create consumer: %w", op, err)
	}

	var (
		topic   string
		timeout int
	)
	if t, ok := ctx.Deadline(); ok {
		timeout = int(time.Until(t).Milliseconds())
	} else {
		timeout = 500
	}

	// analog PING here
	_, err = c.GetMetadata(&topic, false, timeout)
	if err != nil {
		return consumer{}, fmt.Errorf("%s failed to get metadata: %w", op, err)
	}

	return consumer{
		c:               c,
		timelineService: service,
		logger:          logger,
	}, nil
}

func (c consumer) Run(ctx context.Context) (err error) {
	defer func() {
		if defErr := c.c.Close(); defErr != nil {
			err = errors.Join(err,
				fmt.Errorf("failed to close consumer: %w", defErr))
		}
	}()

	if err := c.c.SubscribeTopics(c.topics, nil); err != nil {
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
						return fmt.Errorf("fatal error while read message: %w", err)
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
			_, ok := pkfk.LookupHeaderValue(msg.Headers, pkfk.MessageTypeHeaderKey)
			if !ok {
				c.storeOffset(msg)
				continue
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
