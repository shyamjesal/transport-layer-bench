package main

import (
	"context"
	flatbuffers "github.com/google/flatbuffers/go"
	flatbuf "github.com/shyamjesal/transfer-bench/flatbuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getWithFlatbuf(addr string) []byte {
	conn, err := grpc.Dial(addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})))
	defer conn.Close()
	if err != nil {
		log.Fatalf("[producer] fail to dial: %s", err)
	}
	client := flatbuf.NewGreeterClient(conn)
	b := flatbuffers.NewBuilder(0)
	i := b.CreateString("bla")
	flatbuf.HelloRequestStart(b)
	flatbuf.HelloRequestAddKey(b, i)
	b.Finish(flatbuf.HelloRequestEnd(b))

	ack, err := client.SayHello(context.Background(), b, grpc.CallContentSubtype("flatbuffers"))
	if err != nil {
		log.Fatalf("[producer] client error in string consumption: %s", err)
	}
	return ack.Message()
}
