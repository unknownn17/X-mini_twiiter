package likeretweet

import (
	like_retweet17 "api/internal/protos/like_retweet"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserClinet() like_retweet17.Like_Retweet_CommentClient {
	conn, err := grpc.NewClient("lrc_service:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}
	client := like_retweet17.NewLike_Retweet_CommentClient(conn)
	return client
}
