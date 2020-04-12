package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	pb "xgrpc.v1/hello"
)

func NewHelloClient(addr string) pb.HelloServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("New HelloServiceClient error: %v", err)
	}

	hc := pb.NewHelloServiceClient(conn)
	return hc
}

func main() {
	//conn, err := grpc.Dial("127.0.0.1:50000", grpc.WithInsecure())
	//if err != nil {
	//	log.Fatalf("connect error: %v", err)
	//}
	c := NewHelloClient("127.0.0.1:50000")
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Test",
	})

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println(resp.GetReply())
}
