package main

import (
	pb "github.com/Uncurlynt/gRPC-app/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
