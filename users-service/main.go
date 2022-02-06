package main

import (
	"github.com/BinaryArchaism/users-service/users-service/cache"
	"github.com/BinaryArchaism/users-service/users-service/kafka"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/BinaryArchaism/users-service/users-service/repository"
	"github.com/BinaryArchaism/users-service/users-service/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
	db, err := repository.EstablishPSQLConnection()
	defer repository.CloseConn(db)
	if err != nil {
		logrus.Fatal(err)
	}
	rd, err := cache.InitRedis()
	defer cache.CloseRedis(rd)
	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)

	ch := cache.NewCache(rd, repo)

	report := kafka.NewReports()

	s := grpc.NewServer()
	srv := service.NewUsersService(repo, ch, report)

	models.RegisterUsersServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal(err)
	} else {
		logrus.Infoln("Server listening on :8080")
	}

	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
