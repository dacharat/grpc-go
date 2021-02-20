package service

import (
	"context"
	"fmt"

	"github.com/dacharat/grpc-go/cmd/proto"
)

// Server server
type Server struct {
	proto.AddServiceServer
}

// Add add
func (s *Server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	fmt.Println(a, b)
	result := a + b

	return &proto.Response{Result: result}, nil
}

// Subtract substract
func (s *Server) Subtract(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b

	return &proto.Response{Result: result}, nil
}
