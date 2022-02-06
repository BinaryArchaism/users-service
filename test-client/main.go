package main

import (
	"context"
	"github.com/BinaryArchaism/users-service/test-client/models"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := models.NewUsersServiceClient(conn)
	ctx := context.Background()

	//getalltest(client, ctx)
	//deletetest(client, ctx)
	addtest(client, ctx)
	getalltest(client, ctx)
}

func addtest(client models.UsersServiceClient, ctx context.Context) {
	feature, err := client.AddUser(ctx, &models.UserToAdd{
		FirstName: "Test",
		LastName:  "Test",
		Age:       nil,
		Email:     "RRRR",
	})
	if err != nil {
		log.Fatalf("%v.add(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func getalltest(client models.UsersServiceClient, ctx context.Context) {
	feature, err := client.GetUsers(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("%v.getall(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func deletetest(client models.UsersServiceClient, ctx context.Context) {
	_, err := client.DeleteUser(ctx, &models.UserId{Id: uint64(3)})
	if err != nil {
		log.Fatalf("%v.Delete(_) = _, %v: ", client, err)
	}
}
