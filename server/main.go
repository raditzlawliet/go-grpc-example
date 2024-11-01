package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/raditzlawliet/go-grpc-example/greeter"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "9091", "port")
)

func init() {
	flag.Parse()
}

type server struct {
	greeter.UnimplementedGreeterServer
}

func NewGreeterService() *server {
	return &server{}
}

func (s *server) SayHello(c context.Context, req *greeter.SayHelloRequest) (*greeter.SayHelloResponse, error) {
	log.Println("Incoming SayHello from:", req.Name)
	response := &greeter.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s! Welcome back!", req.Name),
	}
	return response, nil
}

func (s *server) SayHelloAgain(c context.Context, req *greeter.SayHelloRequest) (*greeter.SayHelloResponse, error) {
	log.Println("Incoming SayHelloAgain from:", req.Name)
	response := &greeter.SayHelloResponse{
		Message: fmt.Sprintf("Hello again %s! Welcome back!", req.Name),
	}
	return response, nil
}

func main() {
	log.Println("gRPC Server listening in 0.0.0.0:", *port)
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeterService := NewGreeterService()

	greeter.RegisterGreeterServer(grpcServer, greeterService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
