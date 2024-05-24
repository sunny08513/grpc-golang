// calculator/server.go
package main

import (
	"context"
	"log"
	"math"
	"net"

	pb "grpc-golang/calculator"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
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

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
