package main

import (
	pb "github.com/Uncurlynt/gRPC-app/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":8008"
)

type helloServer struct {
	pb.GreetServiceServer
	//pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
