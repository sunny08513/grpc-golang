// client/main.go
package main

import (
	"context"
	"log"
	"time"

	pb "grpc-golang/calculator"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Sum(ctx, &pb.SumRequest{A: 10, B: 20})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Printf("Sum: %d", r.GetResult())
}
