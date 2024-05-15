package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/srivapra/grpc-demo/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional streaming started ")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Error whiling calling the SayHelloBidirectionalStreaming function %v ", err)
	}
	waitC := make(chan chan struct{})

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v ", err)
			}
			log.Println(res)
		}
		close(waitC)

	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending the stream %v ", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitC
	log.Printf("Bidirectional Streaming finished")
}
