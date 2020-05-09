package main

import (
	"context"
	"fmt"
	server_hello_proto "gRPCTest/proto/server-hello-proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1235", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	client := server_hello_proto.NewTigerServiceClient(conn)
	response, err := client.HelloTiger(context.Background(), &server_hello_proto.HelloRequest{
		Name: "ban",
		Age:  11,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)

	feedResponse, err := client.FeedTiger(context.Background(), &server_hello_proto.FeedRequest{
		Meat:   "pork",
		Weight: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(feedResponse)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := stream.Send(&server_hello_proto.FeedRequest{
				Meat:   "pork",
				Weight: 10,
			})
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(reply)
	}

}
