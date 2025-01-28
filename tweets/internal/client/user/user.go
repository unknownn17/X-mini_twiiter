package userservice

import (
	"log"
	user17 "tweets/internal/protos/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() user17.UserClient {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := user17.NewUserClient(conn)
	return client
}
