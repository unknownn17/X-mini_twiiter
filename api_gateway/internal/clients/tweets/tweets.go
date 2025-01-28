package tweets

import (
	tweets17 "api/internal/protos/tweets"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() tweets17.TweetsClient {
	conn, err := grpc.NewClient("tweets_service:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := tweets17.NewTweetsClient(conn)
	return client
}
