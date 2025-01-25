package greetservice

import (
	"context"
	"log"
	"time"

	pb "github.com/Clement-Jean/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const LONG_OPERATION_DURATION = 3 * time.Second

func (*Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	resC := make(chan string)
	go func() {
		time.Sleep(LONG_OPERATION_DURATION)
		resC <- "Hello " + in.FirstName
	}()

	// todo: try creating a "handleWithDeadline" decorator
	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}
	case res := <-resC:
		return &pb.GreetResponse{
			Result: res,
		}, nil
	}

	return nil, status.Error(codes.Internal, "Unknown")
}
