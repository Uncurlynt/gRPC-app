package main

import (
	"context"
	pb "github.com/Uncurlynt/gRPC-app/proto"
	"io"
	"log"
	"time"
)

func CallHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming started")
	stream, err := client.SayHelloBidirectionalStreaming((context.Background()))
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	ch := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err != io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(ch)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while streaming %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-ch
	log.Printf("Bidirectional streaming finished")
}
