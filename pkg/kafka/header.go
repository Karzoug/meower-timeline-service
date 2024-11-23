//nolint:gosec // false positive on md5 - we use it only for hashing
package kafka

import (
	"crypto/md5"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/protobuf/proto"
)

const MessageTypeHeaderKey = "fngpnt"

func MessageTypeHeaderValue(msg proto.Message) string {
	res := md5.Sum([]byte(proto.MessageName(msg)))
	return fmt.Sprintf("%x", res)
}

func LookupHeaderValue(headers []kafka.Header, key string) ([]byte, bool) {
	for _, header := range headers {
		if header.Key == key {
			return header.Value, true
		}
	}
	return nil, false
}
