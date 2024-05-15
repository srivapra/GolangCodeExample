package main

import (
	"io"
	"log"

	pb "github.com/srivapra/grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesLists{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got Request with names %v ", req.Message)
		messages = append(messages, "Hello ", req.Message)
	}
}
