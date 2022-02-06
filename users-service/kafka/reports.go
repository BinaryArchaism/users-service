package kafka

import (
	"context"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"log"
)

type IReports interface {
	AddUser(id uint64, user *models.UserToAdd) error
}

type reports struct {
	kfka *kafka.Writer
}

func (r *reports) AddUser(id uint64, user *models.UserToAdd) error {
	strUser, err := proto.Marshal(user)
	if err != nil {
		return err
	}
	err = r.kfka.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(string(id)),
			Value: strUser,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := r.kfka.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
	return nil
}

func NewReports() IReports {
	return &reports{InitKafka()}
}
