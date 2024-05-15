package main

import (
	"log"
	"net"

	pb "github.com/srivapra/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listner, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("gRPC server listening at %v", listner.Addr())
	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
