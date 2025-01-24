package calcservice

import (
	"context"
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
