package main

import (
	"flag"
	"github.com/JoeReid/fastTCP"
	"github.com/ishidawataru/sctp"
	"github.com/lucas-clemente/quic-go/http3"
	pb "github.com/shyamjesal/transfer-bench/proto"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"github.com/xtaci/kcp-go/v5"
	"google.golang.org/grpc"
	"io"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"strings"
)

var transferSizeKB = 10
var payloadData []byte

const (
	addr     = "0.0.0.0:8090"
	certFile = "certs/localhost.crt"
	keyFile  = "certs/localhost.key"
)

type consumerServer struct {
	pb.UnimplementedProducerConsumerServer
}

func main() {

	transferProtocol := flag.String("transportProtocol", "quic", "tcp/quic")
	flag.Parse()
	payloadData = make([]byte, transferSizeKB*1024)
	if _, err := rand.Read(payloadData); err != nil {
		log.Fatal(err)
	}
	ip := strings.Split(addr, ":")[0]
	port, err := strconv.Atoi(strings.Split(addr, ":")[1])
	if err != nil {
		log.Fatal(err)
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/", getWithHttp)

	log.Infof("using %s to serve %d bytes", *transferProtocol, len(payloadData))
	if *transferProtocol == "http" {
		http.ListenAndServe(addr, handler)
	} else if *transferProtocol == "quic" {
		log.Fatal(http3.ListenAndServe(addr, certFile, keyFile, handler))
	} else if *transferProtocol == "grpc" {
		grpcServer := grpc.NewServer()
		pb.RegisterProducerConsumerServer(grpcServer, &consumerServer{})
		lis, err := net.Listen("tcp", ":8090")
		if err != nil {
			log.Fatalf("[producer] failed to listen: %v", err)
		}
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("[producer] failed to serve: %s", err)
		}
	} else if *transferProtocol == "fasthttp" {
		log.Fatal(fasthttp.ListenAndServe(addr, fastHTTPHandler))
	} else if *transferProtocol == "kcp" {
		if listener, err := kcp.Listen(addr); err == nil {
			for {
				conn, err := listener.Accept()
				if err != nil {
					log.Fatal(err)
				}
				_, err = conn.Write(payloadData)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			log.Fatal(err)
		}
	} else if *transferProtocol == "tcp" {
		server, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatal(err)
		}
		for {
			// accept connection
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}
			// send file to client
			_, err = conn.Write(payloadData)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else if *transferProtocol == "fasttcp" {
		handler := func(tcpFile io.ReadWriter) {
			log.Info("Processing using fasttcp")
			_, err := tcpFile.Write(payloadData)
			if err != nil {
				log.Fatal(err)
			}
		}
		server := fastTCP.NewServer(addr, handler, fastTCP.TCPOptions{FastOpen: true})

		// Start the TCP server
		err := server.ListenTCP()
		if err != nil {
			panic(err)
		}
	} else if *transferProtocol == "sctp" {
		//not working due to some reason
		ip, err := net.ResolveIPAddr("ip", ip)
		if err != nil {
			log.Fatalf("invalid IP: %v", err)
		}
		sctpAddr := &sctp.SCTPAddr{
			IPAddrs: []net.IPAddr{*ip},
			Port:    port,
		}
		ln, err := sctp.ListenSCTP("sctp", sctpAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v IP: %v", err, sctpAddr)
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatalf("failed to accept: %v", err)
			}

			//log.Printf("Accepted Connection from RemoteAddr: %s", conn.RemoteAddr())
			//wconn := sctp.NewSCTPSndRcvInfoWrappedConn(conn.(*sctp.SCTPConn))
			//sndbuf, err := wconn.GetWriteBuffer()
			//if err != nil {
			//	log.Fatalf("failed to get write buf: %v", err)
			//}
			//rcvbuf, err := wconn.GetReadBuffer()
			//if err != nil {
			//	log.Fatalf("failed to get read buf: %v", err)
			//}
			//log.Printf("SndBufSize: %d, RcvBufSize: %d", sndbuf, rcvbuf)
			//// send file to client
			_, err = conn.Write(payloadData)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
