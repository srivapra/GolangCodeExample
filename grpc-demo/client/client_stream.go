package main

import (
	"context"
	"log"
	"time"

	pb "github.com/srivapra/grpc-demo/proto"
)

func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client streaming started ")

	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the SayHelloClientStreaming function %v ", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v ", err)
		}
		log.Printf("Sent the request with name %s ", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving the response %v ", err)
	}
	log.Printf("Response recieved %s ", res.Message)
}
