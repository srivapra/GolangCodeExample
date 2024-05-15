package main

import (
	"context"
	"io"
	"log"

	pb "github.com/srivapra/grpc-demo/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming Started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("error calling function SayHelloServerStreaming: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming %v ", err)
		}
		log.Println(message)
	}
	log.Printf("Streaming Finished ")
}
