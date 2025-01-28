package main

import (
	"fmt"
	"like/internal/config"
	"like/internal/connections"
	like_retweet17 "like/internal/protos/like_retweet"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Configuration()
	ls, err := net.Listen(c.User.Host, c.User.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connections.NewService()
	like_retweet17.RegisterLike_Retweet_CommentServer(s, server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
