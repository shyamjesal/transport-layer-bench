package main

import (
	"flag"
	"github.com/lucas-clemente/quic-go/http3"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

var transferSizeKB = 10
var payloadData []byte

const (
	addr     = "0.0.0.0:8090"
	certFile = "../certs/localhost.crt"
	keyFile  = "../certs/localhost.key"
)

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
		serveUsingGrpc()
	} else if *transferProtocol == "fasthttp" {
		log.Fatal(fasthttp.ListenAndServe(addr, fastHTTPHandler))
	} else if *transferProtocol == "kcp" {
		serveWithKcp()
	} else if *transferProtocol == "tcp" {
		serveWithTcp()
	} else if *transferProtocol == "tcpmarshall" {
		serveWithTcpMarshal()
	} else if *transferProtocol == "fasttcp" {
		serveWithFastHttp()
	} else if *transferProtocol == "sctp" {
		serveWithSctp(ip, port)
	} else if *transferProtocol == "capnp" {
		serveWithCapnp()
	} else if *transferProtocol == "flatbuf" {
		serveWithFlatbuf()
	} else {
		log.Fatal("Wrong type")
	}
}
