package main

import (
	"context"
	"fmt"
	"grpc_test/greet/greetpb"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

// Greet ddd
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Server Greet function was invoked %v\n", req)
	rep := &greetpb.GreetResponse{
		Result: "hola " + req.GetGreeting().GetFirstName() + " " + req.GetGreeting().GetLastName(),
	}
	return rep, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {

	log.Printf("Server Many Greet functiuon was invoked %v\n", req)

	// firstName := req.GetGreeting().GetFirstName
	// fmt.Println(firstName)
	for i := 0; i < 10; i++ {
		result := "Hello " + req.GetGreeting().GetFirstName() + "number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{Result: result}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func main() {
	fmt.Println("Hello World I am the Server!!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Error fatal listener v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error fatal server v", err)
	}

}
