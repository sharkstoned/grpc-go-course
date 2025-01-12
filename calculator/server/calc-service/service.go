package calcservice

import (
	"context"

	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
)

func (*Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{
		Sum: req.A + req.B,
	}, nil
}
