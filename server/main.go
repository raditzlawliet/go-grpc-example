package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/raditzlawliet/go-grpc-example/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "9091", "port")
)

func init() {
	flag.Parse()
}

type greeterServer struct {
	proto.UnimplementedGreeterServer
}

// SayHello rpc from proto/greeter.proto
func (s *greeterServer) SayHello(c context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	log.Println("Incoming Greeter/SayHello request with data:", req)
	response := &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s! Welcome back!", req.Name),
	}
	return response, nil
}

// SayHelloAgain rpc from proto/greeter
func (s *greeterServer) SayHelloAgain(c context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	log.Println("Incoming Greeter/SayHelloAgain request with data:", req)
	response := &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello again %s! Welcome back!", req.Name),
	}
	return response, nil
}

var store = map[string]string{}

type storeServer struct {
	proto.UnimplementedStoreServer
}

// Set rpc from proto/store
func (s *storeServer) Set(c context.Context, req *proto.SetRequest) (*proto.SetResponse, error) {
	log.Println("Incoming Store/Set request with data:", req)
	store[req.Key] = req.Value
	return &proto.SetResponse{}, nil
}

// Get rpc from proto/store
func (s *storeServer) Get(c context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	log.Println("Incoming Store/Get request with data:", req)
	if value, ok := store[req.Key]; ok {
		return &proto.GetResponse{Value: value}, nil
	}
	return &proto.GetResponse{Value: ""}, nil
}

func main() {
	log.Println("gRPC Server listening in 0.0.0.0:", *port)
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	greeterServer := &greeterServer{}
	storeServer := &storeServer{}

	proto.RegisterGreeterServer(grpcServer, greeterServer)
	proto.RegisterStoreServer(grpcServer, storeServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
