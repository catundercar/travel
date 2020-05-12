package main

import (
	"context"
	"log"
	"net"

	pb "xgrpc.v1/hello"

	"google.golang.org/grpc"
)

const (
	port = ":50000"
)

type HelloServer struct{}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Reply: "你好" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &HelloServer{})
	s.Serve(lis)
}
