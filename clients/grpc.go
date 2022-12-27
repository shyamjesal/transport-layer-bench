package main

import (
	"context"
	pb_client "github.com/shyamjesal/transfer-bench/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getWithGrpc(addr string) []byte {
	conn, err := grpc.Dial(addr, grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	if err != nil {
		log.Fatalf("[producer] fail to dial: %s", err)
	}
	client := pb_client.NewProducerConsumerClient(conn)
	ack, err := client.FetchBytes(context.Background(), &pb_client.Empty{})
	if err != nil {
		log.Fatalf("[producer] client error in string consumption: %s", err)
	}
	return ack.Value
}
