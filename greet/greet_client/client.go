package main

import (
	"context"
	"grpc_test/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello World I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c, "Cesar ", "Alvarado")
	doUnary(c, "Kaisy ", "Paipay")
	doUnary(c, "Stacy ", "Alvarado")

}

func doUnary(c greetpb.GreetServiceClient, firstName, lastName string) {
	// fmt.Printf("Creating client %f", c)

	respond, err := c.Greet(context.Background(), &greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: firstName, LastName: lastName}})
	if err != nil {
		log.Printf("error %v", err)
	}

	log.Printf("%v", respond.Result)
}
