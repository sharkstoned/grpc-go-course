package main

import (
	"log"
	"net"
	"os"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	sum "github.com/Clement-Jean/grpc-go-course/calculator/server/sum"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:27519"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen for tcp: %v", err)
		os.Exit(1)
	}
	defer listener.Close()
	log.Printf("Listening at %s\n", addr)

	server := grpc.NewServer()
	registerServices(server)

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
		os.Exit(1)
	}
}

func registerServices(server *grpc.Server) {
	pb.RegisterSumServiceServer(server, &sum.Server{})
}
