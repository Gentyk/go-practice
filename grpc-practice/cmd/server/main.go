package main

import (
	"example/web-service-gin/go-practice/grpc-practice/pkg/api"
	"example/web-service-gin/go-practice/grpc-practice/pkg/api/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &server.GRPCServer{}
	api.RegisterSenderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}