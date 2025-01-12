//go:build !test
// +build !test

package main

import (
	"log"
	"net"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

// var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:50051"

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

	s := grpc.NewServer()
	// Server{} is a collection of endpoints defined with protobuf
	// At this point we bind together grpc server s with the endpoints definition
	pb.RegisterGreetServiceServer(s, &Server{})

	// Make grpc server listen on tcp via the listener
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
