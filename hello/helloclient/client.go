package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hello/hellopb"
	"google.golang.org/grpc"
)

func doSomething(s string) {
	fmt.Println("", s)
}

func startPolling1() {
	for {
		time.Sleep(5 * time.Second)
		go doSomething("hello")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Path[1:])
}
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
	go startPolling1()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
