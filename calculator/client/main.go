package main

import (
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

	// sumServiceClient := pb.NewSumServiceClient(conn)
	// primesServiceClient := pb.NewPrimesServiceClient(conn)
	// avgServiceClient := pb.NewAvgServiceClient(conn)
	maxServiceClient := pb.NewMaxServiceClient(conn)

	// getSum(sumServiceClient)
	// getPrimes(primesServiceClient)
	// getAverage(avgServiceClient)
	getMaxSync(maxServiceClient, []int32{-5, -7, 13, 13, 17})
	getMaxAsync(maxServiceClient, []int32{-5, -7, 13, 13, 17})
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
