// calculator/server.go
package main

import (
	"context"
	"log"
	"math"
	"net"

	pb "grpc-golang/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func LoggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("Unary RPC: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}
	return resp, err
}

func LoggingStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Printf("Streaming RPC: %s", info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}
	return err
}

func (s *server) Sum(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := req.GetA() + req.GetB()
	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	result := req.GetA() * req.GetB()
	return &pb.Response{Result: result}, nil
}

func (s *server) GetPrimes(req *pb.PrimeRequest, stream pb.Calculator_GetPrimesServer) error {
	number := req.GetN()
	if number < 0 {
		return status.Errorf(codes.InvalidArgument, "The number must be greater than zero")
	}
	for i := 2; i <= int(number); i++ {
		if isPrime(i) {
			stream.Send(&pb.PrimeResponse{Result: int32(i)})
		}
	}
	return nil
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(LoggingUnaryInterceptor), grpc.StreamInterceptor(LoggingStreamInterceptor))
	pb.RegisterCalculatorServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
