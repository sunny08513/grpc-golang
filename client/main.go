// client/main.go
package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-golang/calculator"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

// LoggingUnaryClientInterceptor is a unary client interceptor for logging
func LoggingUnaryClientInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	log.Printf("Unary RPC: %s", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}
	return err
}

// LoggingStreamClientInterceptor is a stream client interceptor for logging
func LoggingStreamClientInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	log.Printf("Streaming RPC: %s", method)
	cs, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		st, _ := status.FromError(err)
		log.Printf("RPC failed with status: %v", st.Code())
	}
	return cs, err
}

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(LoggingUnaryClientInterceptor), grpc.WithStreamInterceptor(LoggingStreamClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	calClient := pb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := calClient.Sum(ctx, &pb.Request{A: 10, B: 30})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("sum: %d", r.GetResult())

	r, err = calClient.Multiply(ctx, &pb.Request{A: 10, B: 30})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	log.Printf("multiply: %d", r.GetResult())

	primeStream, err := calClient.GetPrimes(ctx, &pb.PrimeRequest{N: -1})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			log.Fatalf("gRPC error: %v, %v", st.Code(), st.Message())
		} else {
			log.Fatalf("could not get primes: %v", err)
		}
		return
	}

	for {
		prime, err := primeStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				log.Fatalf("streaming error: %v, %v", st.Code(), st.Message())
			} else {
				log.Fatalf("error receiving prime: %v", err)
			}
		}
		log.Printf("Prime: %d\n", prime.GetResult())
	}
}
