package main

import (
	"context"
	flatbuffers "github.com/google/flatbuffers/go"
	flatbuf "github.com/shyamjesal/transfer-bench/flatbuf"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type greeterServer struct {
	flatbuf.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *flatbuf.HelloRequest) (*flatbuffers.Builder, error) {
	log.Info("received %s", req.Key())
	b := flatbuffers.NewBuilder(0)
	idx := b.CreateByteString(payloadData)
	flatbuf.HelloReplyStart(b)
	flatbuf.HelloReplyAddMessage(b, idx)
	b.Finish(flatbuf.HelloReplyEnd(b))
	return b, nil
}

func serveWithFlatbuf() {
	codec := &flatbuffers.FlatbuffersCodec{}
	grpcServer := grpc.NewServer(grpc.ForceServerCodec(codec))
	flatbuf.RegisterGreeterServer(grpcServer, &greeterServer{})
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("[producer] failed to listen: %v", err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("[producer] failed to serve: %s", err)
	}
}
