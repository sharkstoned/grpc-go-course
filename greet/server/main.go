//go:build !test
// +build !test

package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	greetservice "github.com/Clement-Jean/grpc-go-course/greet/server/greet-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr = "0.0.0.0:50051"
var tls = true

// func main() {
// 	lis, err := net.Listen("tcp", addr)

// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v\n", err)
// 	}

// 	defer lis.Close()
// 	log.Printf("Listening at %s\n", addr)

// 	opts := []grpc.ServerOption{}

// 	tls := false // change that to true if needed
// 	if tls {
// 		certFile := "ssl/server.crt"
// 		keyFile := "ssl/server.pem"
// 		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

// 		if err != nil {
// 			log.Fatalf("Failed loading certificates: %v\n", err)
// 		}
// 		opts = append(opts, grpc.Creds(creds))
// 	}

// 	opts = append(opts, grpc.ChainUnaryInterceptor(LogInterceptor(), CheckHeaderInterceptor()))

// 	s := grpc.NewServer(opts...)
// 	pb.RegisterGreetServiceServer(s, &Server{})

// 	defer s.Stop()
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("Failed to serve: %v\n", err)
// 	}
// }

func main() {
	// Create a tcp listener
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening at %s\n", addr)

	serverOpts := []grpc.ServerOption{}
	if tls {
		creds, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			log.Fatalf("Failed to load tls cert %v\n", err)
		}

		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	s := grpc.NewServer(serverOpts...)
	// Server{} is a collection of endpoints defined with protobuf
	// At this point we bind together grpc server s with the endpoints definition
	pb.RegisterGreetServiceServer(s, &greetservice.Server{})

	// Make grpc server listen on tcp via the listener
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
