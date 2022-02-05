package main

import (
	"context"
	"fmt"
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

	repo := repository.NewRepository(db)

	s := grpc.NewServer()
	srv := service.NewUsersService(repo)

	srv.AddUser(context.Background(), &models.UserToAdd{
		FirstName: "asdf",
		LastName:  "fasd",
		Age:       nil,
		Email:     "fasdf",
	})
	srv.AddUser(context.Background(), &models.UserToAdd{
		FirstName: "qwer",
		LastName:  "qwer",
		Age:       nil,
		Email:     "qwer",
	})
	fmt.Println(srv.GetUsers(context.Background(), nil))
	srv.DeleteUser(context.Background(), &models.UserId{Id: 1})
	fmt.Println(srv.GetUsers(context.Background(), nil))

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
