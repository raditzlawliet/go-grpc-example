package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"raditzlawliet/go-grpc-example/protobuf"
	"time"

	"google.golang.org/grpc"
)

var (
	host = flag.String("host", "localhost", "host")
	port = flag.String("port", "9995", "port")
)

func init() {
	flag.Parse()
}

func main() {
	log.Println("Client Alpha!")

	client := Client{Host: *host, Port: *port}
	client.Connect()

	// try to sending message ~
	r, e := client.GetMessage("Resistance")
	if e != nil {
		log.Println("Error GetMessage, ", e.Error())
	} else {
		log.Println("Message Receive frmo ZERO: ", r.Message)
	}
}

type Client struct {
	Host          string
	Port          string
	Connection    *grpc.ClientConn
	ClientService protobuf.MessageServiceClient
}

func (c *Client) Connect() {
	serverAddress := fmt.Sprintf("%s:%s", c.Host, c.Port)
	log.Println("Connecting Server ZERO in ", serverAddress)

	// dialing
	var err error
	c.Connection, err = grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect: %v", err)
		log.Println("try to reconnecting after 5s")
		time.Sleep(5 * time.Second)
		defer c.Connect()
		return
	}

	// Creating client service
	c.ClientService = protobuf.NewMessageServiceClient(c.Connection)
}
func (c *Client) GetMessage(name string) (*protobuf.Response, error) {
	// creating context with timeout
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(10*time.Second))

	// request
	response, err := c.ClientService.Get(ctx, &protobuf.Request{
		Name: name,
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}
