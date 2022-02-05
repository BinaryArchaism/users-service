package service

import (
	"context"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/BinaryArchaism/users-service/users-service/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type usersService struct {
	models.UnimplementedUsersServiceServer
	repo repository.IRepository
}

func (u *usersService) AddUser(ctx context.Context, user *models.UserToAdd) (*emptypb.Empty, error) {
	_, err := u.repo.AddUser(ctx, user)
	if err != nil {
		logrus.Debug(err)
	}
	//TODO clickhouse send user with id
	return nil, err
}

func (u *usersService) DeleteUser(ctx context.Context, id *models.UserId) (*emptypb.Empty, error) {
	return nil, u.repo.DeleteUser(ctx, id)
}

func (u usersService) GetUsers(ctx context.Context, _ *emptypb.Empty) (*models.Users, error) {
	return u.repo.GetUsers(ctx)
}

func NewUsersService(repo repository.IRepository) *usersService {
	return &usersService{repo: repo}
}
