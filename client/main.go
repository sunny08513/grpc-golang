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
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

	primeStream, err := calClient.GetPrimes(ctx, &pb.PrimeRequest{N: 10})
	if err != nil {
		log.Fatalf("error receiving primeStream: %v", err)
	}

	for {
		prime, err := primeStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving prime: %v", err)
		}
		log.Printf("Prime: %d\n", prime.GetResult())
	}
}
