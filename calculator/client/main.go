package main

import (
	"context"
	"io"
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
	primesServiceClient := pb.NewPrimesServiceClient(conn)
	avgServiceClient := pb.NewAvgServiceClient(conn)

	sumResp, err := sumServiceClient.Sum(
		context.Background(),
		&pb.SumRequest{A: 3, B: 5},
	)
	if err != nil {
		log.Printf("Request failed - sum(): %v", err)
		return
	}

	log.Printf("Sum: %d", sumResp.Sum)

	// ---

	primesStream, err := primesServiceClient.Primes(context.Background(), &pb.PrimesRequest{Input: 120})
	if err != nil {
		log.Printf("Failed to acquire primes stream: %v", err)
	}
	for {
		result, err := primesStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Failed to get value from primes stream: %v", err)
			break
		}

		log.Println(result.Factor)
	}

	// ---

	stream, err := avgServiceClient.Avg(context.Background())
	if err != nil {
		log.Printf("Failed to acquire avg stream: %v", err)
	}
	for _, operand := range []int64{5, 12, 43, 6} {
		log.Printf("Sending operand %d to avg\n", operand)
		err := stream.Send(&pb.AvgRequest{
			Input: operand,
		})
		if err != nil {
			log.Printf("Failed to send operand to avg: %v", err)
			return
		}
	}

	avgResult, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Failed to get avg result: %v", err)
		return
	}

	log.Printf("Avg: %f", avgResult.Avg)
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
