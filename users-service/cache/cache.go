package cache

import (
	"context"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/BinaryArchaism/users-service/users-service/repository"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
	"time"
)

type ICache interface {
	GetUsers(ctx context.Context) (*models.Users, error)
}

type cache struct {
	rd             *redis.Client
	repo           repository.IRepository
	lastUpdateTime time.Time
	countOfItems   int //TODO использовать другой способ
}

func (ch *cache) updateCache(ctx context.Context) error {
	ch.countOfItems = 0
	users, err := ch.repo.GetUsers(ctx)
	if err != nil {
		logrus.Debug(err)
		return err
	}
	for _, user := range users.Users {
		strUser, err := proto.Marshal(user)
		if err != nil {
			return err
		}
		ch.rd.Set(string(ch.countOfItems), strUser, time.Minute)
		ch.countOfItems++
	}
	return nil
}

func (ch *cache) GetUsers(ctx context.Context) (*models.Users, error) {
	if time.Since(ch.lastUpdateTime).Seconds() > 60 {
		ch.updateCache(ctx)
		ch.lastUpdateTime = time.Now()
	}
	var users models.Users
	for i := 0; i < ch.countOfItems; i++ {
		strUser, err := ch.rd.Get(string(i)).Result()
		if err != nil {
			return nil, err
		}
		var user models.FullUser
		err = proto.Unmarshal([]byte(strUser), &user)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}
	return &users, nil
}

func NewCache(rd *redis.Client, repo repository.IRepository) ICache {
	return &cache{rd: rd, repo: repo}
}
