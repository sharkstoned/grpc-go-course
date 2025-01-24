package main

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	conn := establishGrpcConnection()
	defer conn.Close()

	greetServiceClient := pb.NewGreetServiceClient(conn)

	program(greetServiceClient)
}

func establishGrpcConnection() *grpc.ClientConn {
	transportCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(addr, transportCredentials)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
		os.Exit(1)
	}

	return conn
}

func program(client pb.GreetServiceClient) {
	doGreet(client)
	doGreetManyTimes(client)
	doLongGreet(client)
}

// func main() {
// 	tls := false // change that to true if needed
// 	opts := []grpc.DialOption{}

// 	if tls {
// 		certFile := "ssl/ca.crt"
// 		creds, err := credentials.NewClientTLSFromFile(certFile, "")

// 		if err != nil {
// 			log.Fatalf("Error while loading CA trust certificate: %v\n", err)
// 		}
// 		opts = append(opts, grpc.WithTransportCredentials(creds))
// 	} else {
// 		creds := grpc.WithTransportCredentials(insecure.NewCredentials())
// 		opts = append(opts, creds)
// 	}

// 	opts = append(opts, grpc.WithChainUnaryInterceptor(LogInterceptor(), AddHeaderInterceptor()))

// 	conn, err := grpc.Dial(addr, opts...)

// 	if err != nil {
// 		log.Fatalf("Did not connect: %v", err)
// 	}

// 	defer conn.Close()
// 	c := pb.NewGreetServiceClient(conn)

// 	doGreet(c)
// 	// doGreetManyTimes(c)
// 	// doLongGreet(c)
// 	// doGreetEveryone(c)
// 	// doGreetWithDeadline(c, 5*time.Second)
// 	// doGreetWithDeadline(c, 1*time.Second)
// }
