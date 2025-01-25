package calcservice

import (
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Sum: req.A + req.B,
	}, nil
}

func (*Server) Primes(req *pb.PrimesRequest, stream grpc.ServerStreamingServer[pb.PrimesResponse]) error {
	log.Printf("Primes endpoint hit, req.Input == %d", req.Input)
	currentNum := req.Input
	divider := int32(2)

	for currentNum > 1 {
		if currentNum%divider == 0 {
			stream.Send(&pb.PrimesResponse{
				Factor: divider,
			})
			currentNum /= divider
			continue
		}

		divider++
	}

	return nil
}

func (*Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Println("Avg endpoint hit")

	var sum int64
	var operandsQty int64

	for {
		val, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: float32(sum) / float32(operandsQty),
			})
		}
		if err != nil {
			log.Printf("Failed to get avg stream input value: %v", err)
			break
		}

		sum += val.Input
		operandsQty++
	}

	return nil
}

func (*Server) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	log.Println("Max endpoint hit")

	var currentMax int32 = 0
	firstValIsReceived := false

	for {
		nextVal, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Failed to get nextVal from Max stream: %v", err)
		}

		if !firstValIsReceived {
			currentMax = nextVal.Input
			firstValIsReceived = true
		} else {
			if currentMax < nextVal.Input {
				currentMax = nextVal.Input
			}
		}

		time.Sleep(time.Duration(rand.Intn(101)) * time.Millisecond)
		err = stream.Send(&pb.MaxResponse{Max: currentMax})
		if err != nil {
			log.Printf("Failed to send Max value: %v", err)
		}
	}
}

func (*Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	if req.Input < 0 {
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative argument: %d", req.Input),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(req.Input)),
	}, nil
}
