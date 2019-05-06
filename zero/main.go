package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/raditzlawliet/go-grpc-example/protobuf"

	"google.golang.org/grpc"
)

var (
	port = flag.String("port", "9995", "port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("ZERO Server listening in :", *port)
	lis, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	messageService := MessageService()
	protobuf.RegisterMessageServiceServer(grpcServer, messageService)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	name string
}

func MessageService() *server {
	return &server{
		name: "Server",
	}
}

func (s *server) Get(c context.Context, req *protobuf.Request) (*protobuf.Response, error) {
	log.Println("Incoming Request with name: ", req.Name)
	response := &protobuf.Response{
		Message: fmt.Sprintf("Hello %s! Welcome back!", req.Name),
	}
	return response, nil
}
