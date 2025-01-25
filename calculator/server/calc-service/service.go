package calcservice

import (
	"context"
	"io"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	grpc "google.golang.org/grpc"
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
