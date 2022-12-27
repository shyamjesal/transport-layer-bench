package main

import (
	"context"
	pb "github.com/shyamjesal/transfer-bench/proto"
)

func (s *consumerServer) FetchBytes(ctx context.Context, str *pb.Empty) (*pb.Reply, error) {
	return &pb.Reply{Value: payloadData}, nil
}
