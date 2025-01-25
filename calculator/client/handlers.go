package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func getSum(client pb.SumServiceClient) {
	sumResp, err := client.Sum(
		context.Background(),
		&pb.SumRequest{A: 3, B: 5},
	)
	if err != nil {
		log.Printf("Request failed - sum(): %v", err)
		return
	}

	log.Printf("Sum: %d", sumResp.Sum)
}

func getPrimes(client pb.PrimesServiceClient) {
	primesStream, err := client.Primes(context.Background(), &pb.PrimesRequest{Input: 120})
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

}

func getAverage(client pb.AvgServiceClient) {
	stream, err := client.Avg(context.Background())
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

func getMaxSync(client pb.MaxServiceClient, inputValues []int32) {
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Printf("Failed to create stream: %v", err)
	}

	for _, val := range inputValues {
		log.Printf("Sending value %d", val)
		stream.Send(&pb.MaxRequest{Input: val})

		val, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Failed to receive value from stream: %v", err)
		}

		log.Printf("Received new max value: %d", val.Max)
	}
	stream.CloseSend()
}

func getMaxAsync(client pb.MaxServiceClient, inputValues []int32) {
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Printf("Failed to create stream: %v", err)
	}

	waitChannel := make(chan struct{})
	go func() {
		for _, val := range inputValues {
			log.Printf("Sending value %d", val)
			time.Sleep(time.Duration(rand.Intn(1001)) * time.Millisecond)
			stream.Send(&pb.MaxRequest{Input: val})
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			val, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Failed to receive value from stream: %v", err)
			}

			log.Printf("Received new max value: %d", val.Max)
		}

		close(waitChannel)
	}()

	<-waitChannel
}
