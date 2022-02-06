package kafka

import (
	"github.com/segmentio/kafka-go"
)

func InitKafka() *kafka.Writer {
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic", 0)
	//if err != nil {
	//	panic(err.Error())
	//}
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "topic",
	}
	return w
}
