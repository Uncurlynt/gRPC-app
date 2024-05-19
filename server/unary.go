package main

//46:20
import (
	"context"
	pb "github.com/Uncurlynt/gRPC-app/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
