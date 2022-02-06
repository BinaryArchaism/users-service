package service

import (
	"context"
	"github.com/BinaryArchaism/users-service/users-service/cache"
	"github.com/BinaryArchaism/users-service/users-service/kafka"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/BinaryArchaism/users-service/users-service/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

type usersService struct {
	models.UnimplementedUsersServiceServer
	repo   repository.IRepository
	ch     cache.ICache
	report kafka.IReports
}

func (u *usersService) AddUser(ctx context.Context, user *models.UserToAdd) (*emptypb.Empty, error) {
	id, err := u.repo.AddUser(ctx, user)
	if err != nil {
		logrus.Debug(err)
	}
	err = u.report.AddUser(id, user)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, err
}

func (u *usersService) DeleteUser(ctx context.Context, id *models.UserId) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, u.repo.DeleteUser(ctx, id)
}

func (u usersService) GetUsers(ctx context.Context, _ *emptypb.Empty) (*models.Users, error) {
	return u.ch.GetUsers(ctx)
}

func NewUsersService(repo repository.IRepository, ch cache.ICache, report kafka.IReports) *usersService {
	return &usersService{repo: repo, ch: ch, report: report}
}
