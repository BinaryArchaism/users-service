package main

import (
	"context"
	"github.com/BinaryArchaism/users-service/test-client/models"
	"google.golang.org/grpc"
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

	deletetest(client, ctx)
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
	feature, err := client.GetUsers(ctx, nil)
	if err != nil {
		log.Fatalf("%v.getall(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}

func deletetest(client models.UsersServiceClient, ctx context.Context) {
	feature, err := client.DeleteUser(ctx, &models.UserId{Id: uint64(1)})
	if err != nil {
		log.Fatalf("%v.Delete(_) = _, %v: ", client, err)
	}
	log.Println(feature)
}
