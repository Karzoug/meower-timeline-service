package kafka

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/protobuf/proto"

	timelineApi "github.com/Karzoug/meower-timeline-service/internal/timeline/client/kafka/gen/timeline/v1"
)

const timelineTopic = "timeline"

var topic = timelineTopic

type producer struct {
	p  *kafka.Producer
	wg *sync.WaitGroup
}

func NewProducer(ctx context.Context, cfg Config) (producer, func(context.Context) error, error) {
	const op = "create kafka producer:"

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
	})
	if err != nil {
		return producer{}, nil, fmt.Errorf("%s failed to create producer: %w", op, err)
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
	_, err = p.GetMetadata(&topic, false, timeout)
	if err != nil {
		return producer{}, nil, fmt.Errorf("%s failed to get metadata: %w", op, err)
	}

	closeFn := func(ctx context.Context) error {
		p.Flush(15 * 1000)
		p.Close()

		return nil
	}

	return producer{p: p}, closeFn, nil
}

func (p producer) SendFollowersGotNewPostEvent(followerIDs []string, authorID, postID string) error {
	if len(followerIDs) == 0 {
		return nil
	}

	delivChan := make(chan kafka.Event, 1)
	ev := &timelineApi.FollowerGotNewPostEvent{
		PostId:   postID,
		AuthorId: authorID,
	}

	p.wg.Add(1)
	defer p.wg.Done()

	for i := range followerIDs {
		ev.FollowerId = followerIDs[i]
		val, err := proto.Marshal(ev)
		if err != nil {
			return fmt.Errorf("failed to marshal follower got new post event: %w", err)
		}

		var ch chan kafka.Event
		if i == len(followerIDs)-1 {
			ch = delivChan
		}
		if err := p.p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(followerIDs[i]),
			Value: val,
		}, ch); err != nil {
			return fmt.Errorf("failed to produce follower got new post event: %w", err)
		}
	}

	delivEvent := <-delivChan
	switch ev := delivEvent.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			return fmt.Errorf("failed to deliver message: %w", ev.TopicPartition.Error)
		}
	case kafka.Error:
		return fmt.Errorf("kafka producer error: %w", ev)
	}

	return nil
}
