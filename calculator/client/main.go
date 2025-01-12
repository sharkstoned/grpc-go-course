package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

var addr string = "0.0.0.0:27519"

func main() {
	conn := establishConnection()
	defer conn.Close()

	sumServiceClient := pb.NewSumServiceClient(conn)

	sumResp, err := sumServiceClient.Sum(
		context.Background(),
		&pb.SumRequest{A: 3, B: 5},
	)
	if err != nil {
		log.Printf("Request failed - sum(): %v", err)
	}

	log.Printf("Sum: %d", sumResp.Sum)

}

func establishConnection() *grpc.ClientConn {
	transportCredentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(addr, transportCredentials)
	if err != nil {
		log.Fatalf("Failed to establish connection: %v", err)
		os.Exit(1)
	}

	return conn
}
