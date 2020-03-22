package main

import (
	"context"
	"fmt"
	calcularpb "grpc_test/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello World I'm a Cliente")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error %v", err)
	}

	defer conn.Close()

	srvSumClient := calcularpb.NewSumServiceClient(conn)
	// doUnarySuma(srvSumClient, 10, 30)
	// doUnarySuma(srvSumClient, 20, 50)
	doUnarySuma(srvSumClient, 50, 30)
}

func doUnarySuma(c calcularpb.SumServiceClient, Number1, Number2 int32) {
	res, err := c.Sum(context.Background(), &calcularpb.SumRequest{Number1: Number1, Number2: Number2})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println(res.GetResult())
	return
}
