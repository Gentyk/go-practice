package server

import (
	"context"
	"example/web-service-gin/go-practice/grpc-practice/pkg/api"
)

type GRPCServer struct {
	// Встраиваем generated UnimplementedSenderServer для forward-compatibility
	// и чтобы GRPCServer удовлетворял интерфейсу api.SenderServer.
	api.UnimplementedSenderServer
}

func (s *GRPCServer) Send(ctx context.Context, req *api.Request) (*api.Response, error) {
	return &api.Response{Result: req.GetField1()}, nil
}