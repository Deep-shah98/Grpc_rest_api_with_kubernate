package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hello/hellopb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello client running ...")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	client := hellopb.NewHelloServiceClient(cc)
	request := &hellopb.HelloRequest{Name: "Kloudone Learner"}

	resp, _ := client.Hello(context.Background(), request)
	fmt.Printf("Receive response => [%v]", resp.Greeting)
}
