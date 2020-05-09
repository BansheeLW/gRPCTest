package main

import (
	"gRPCTest/proto/server-hello-proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type TigerServer struct{}

func (t *TigerServer) FeedTiger(ctx context.Context, req *server_hello_proto.FeedRequest) (*server_hello_proto.FeedResponse, error) {
	resp := &server_hello_proto.FeedResponse{
		HowlCount: req.Weight * 2,
	}
	return resp, nil
}

func (t *TigerServer) Channel(stream server_hello_proto.TigerService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		resp := &server_hello_proto.FeedResponse{
			HowlCount: args.Weight * 2,
		}
		err = stream.Send(resp)
		if err != nil {
			return err
		}
	}
}

func (t *TigerServer) HelloTiger(ctx context.Context, req *server_hello_proto.HelloRequest) (*server_hello_proto.HelloResponse, error) {
	resp := &server_hello_proto.HelloResponse{}
	resp.Name = req.Name
	resp.Age = req.Age + 11
	return resp, nil
}

func main() {
	grpcServer := grpc.NewServer()
	server_hello_proto.RegisterTigerServiceServer(grpcServer, new(TigerServer))

	listener, err := net.Listen("tcp", "127.0.0.1:1235")
	if err != nil {
		log.Fatal(err)
	}
	_ = grpcServer.Serve(listener)
}
