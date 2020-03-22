package main

import (
	"context"
	"fmt"
	calcularpb "grpc_test/calculator/calculatorpb"

	"google.golang.org/grpc"

	// "grpc_test/greet/greetpb"
	"log"
	"net"
)

type server struct{}

// Greet ddd
// func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
// 	fmt.Printf("Server Greet function was invoked %v\n", req)
// 	rep := &greetpb.GreetResponse{
// 		Result: "hola " + req.GetGreeting().GetFirstName() + " " + req.GetGreeting().GetLastName(),
// 	}
// 	return rep, nil
// }

func (*server) Sum(ctx context.Context, req *calcularpb.SumRequest) (*calcularpb.SumResponse, error) {
	fmt.Printf("Server Calcular function was invoked %v\n", req)
	// res := &calcularpb.SumResponse{Result: req.GetNumber1() + req.GetNumber2()}
	res := &calcularpb.SumResponse{Result: req.GetNumber1() - req.GetNumber2()}
	return res, nil
}

func main() {
	fmt.Println("Hola Mundo yo soy el Servidor GRPC!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Error fatal listener v", err)
	}

	s := grpc.NewServer()
	calcularpb.RegisterSumServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error fatal server v", err)
	}
}
