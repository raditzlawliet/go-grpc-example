package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/raditzlawliet/go-grpc-example/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "World"
)

var (
	addr = flag.String("addr", "localhost:9091", "the server address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("Client gRPC!")

	// connect to server
	client := grpcClient{TargetAddress: *addr}
	client.Connect()

	// try to sending message ~
	{
		r, e := client.SayHello(*name)
		if e != nil {
			log.Println("Error GetMessage, ", e.Error())
		} else {
			log.Println("Message Receive from Server: ", r.Message)
		}
	}

	// try to sending another message ~
	{
		r, e := client.SayHelloAgain(*name)
		if e != nil {
			log.Println("Error GetMessage, ", e.Error())
		} else {
			log.Println("Message Receive from Server: ", r.Message)
		}
	}
}

type grpcClient struct {
	TargetAddress string
	Connection    *grpc.ClientConn
	GreeterClient greeter.GreeterClient
}

func (c *grpcClient) Connect() {
	log.Println("Connecting Server in", c.TargetAddress)

	// connect
	var err error
	c.Connection, err = grpc.NewClient(c.TargetAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(fmt.Sprintf("did not connect: %v", err))
		log.Println("try to reconnecting after 5s")
		time.Sleep(5 * time.Second)
		defer c.Connect()
		return
	}

	// Creating client service
	c.GreeterClient = greeter.NewGreeterClient(c.Connection)
}

func (c *grpcClient) SayHello(name string) (*greeter.SayHelloResponse, error) {
	// creating context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	// request
	response, err := c.GreeterClient.SayHello(ctx, &greeter.SayHelloRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *grpcClient) SayHelloAgain(name string) (*greeter.SayHelloResponse, error) {
	// creating context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(1*time.Second))
	defer cancel()

	// request
	response, err := c.GreeterClient.SayHelloAgain(ctx, &greeter.SayHelloRequest{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
