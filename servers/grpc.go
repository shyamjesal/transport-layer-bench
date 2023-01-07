package main

import (
	"context"
	pb "github.com/shyamjesal/transfer-bench/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type consumerServer struct {
	pb.UnimplementedProducerConsumerServer
}

func (s *consumerServer) FetchBytes(ctx context.Context, str *pb.Empty) (*pb.Reply, error) {
	return &pb.Reply{Value: payloadData}, nil
}

func serveUsingGrpc() {
	grpcServer := grpc.NewServer()
	pb.RegisterProducerConsumerServer(grpcServer, &consumerServer{})
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("[producer] failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[producer] failed to serve: %s", err)
	}
}
