package kafka

import (
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"log"
)

type IReports interface {
	AddUser(id uint64, user *models.UserToAdd) error
}

type reports struct {
	kfka *kafka.Conn
}

func (r *reports) AddUser(id uint64, user *models.UserToAdd) error {
	strUser, err := proto.Marshal(user)
	if err != nil {
		return err
	}
	_, err = r.kfka.WriteMessages(kafka.Message{
		Partition: 0,
		Key:       []byte(string(id)),
		Value:     strUser,
	})
	if err != nil {
		return err
	}

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	return nil
}

func NewReports() IReports {
	return &reports{InitKafka()}
}
