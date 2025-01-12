package greetservice

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
)

func (*Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet was invoked with %v\n", req)
	return &pb.GreetResponse{Result: "Hello " + req.FirstName}, nil
}
