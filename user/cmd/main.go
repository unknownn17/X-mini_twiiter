package main

import (
	"cmd/main.go/internal/config"
	"cmd/main.go/internal/connections"
	user17 "cmd/main.go/internal/protos"
	"fmt"
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
	server := connections.Newservice()
	user17.RegisterUserServer(s, server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}
