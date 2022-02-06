package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func InitKafka() *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic", 0)
	if err != nil {
		panic(err.Error())
	}
	return conn
}
