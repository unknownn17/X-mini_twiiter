package main

import (
	"fmt"
	"log"
	"net"
	"tweets/internal/config"
	"tweets/internal/connections"
	tweets17 "tweets/internal/protos/tweets"

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
	tweets17.RegisterTweetsServer(s, server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
