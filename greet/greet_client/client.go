package main

import (
	"fmt"
	"log"

	"github.com/alfaroqi/golang-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Printf("Hello, I'm a client\n")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
