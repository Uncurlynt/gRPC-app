package main

import (
	pb "github.com/Uncurlynt/gRPC-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	port = ":8008"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	name := &pb.NameList{
		Names: []string{"Artemii", "Andreev", "Alekseevich", "Aviasales"},
	}

	//callSayHelloServerStream(client, name)
	//CallSayHelloClientStream(client, name)
	CallHelloBidirectionalStream(client, name)
}
