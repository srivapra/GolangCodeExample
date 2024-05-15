package main

import (
	"log"

	pb "github.com/srivapra/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Prashant", "Alice", "Bob"},
	}

	callSayHello(client)
	callSayHelloServerStreaming(client, names)
	callSayHelloClientStreaming(client, names)
	callSayHelloBidirectionalStreaming(client, names)

}
