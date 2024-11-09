package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/raditzlawliet/go-grpc-example/proto"
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
	client, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	greeterClient := proto.NewGreeterClient(client)
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := greeterClient.SayHello(ctx, &proto.SayHelloRequest{
			Name: "Radit",
		})
		fmt.Println("Greeter/SayHello", response, err)
		cancel()
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := greeterClient.SayHelloAgain(ctx, &proto.SayHelloRequest{
			Name: "Radit",
		})
		fmt.Println("Greeter/SayHelloAgain", response, err)
		cancel()
	}

	storeClient := proto.NewStoreClient(client)
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := storeClient.Set(ctx, &proto.SetRequest{
			Key:   "key-something-1",
			Value: "some-value",
		})
		fmt.Println("Store/Set", response, err)
		cancel()
	}
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := storeClient.Get(ctx, &proto.GetRequest{
			Key: "key-something-1",
		})
		fmt.Println("Store/Get", response, err)
		cancel()
	}

}
