package main

import (
	"context"
	"log"
	"time"

	pb "github.com/srivapra/grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	log.Printf("Response from gRPC server's SayHello function: %s", res.GetMessage())
}
