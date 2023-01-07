package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/ishidawataru/sctp"
	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// We start a server echoing data on the first stream the client opens,
// then connect with a client, send the message, and wait for its receipt.
func main() {
	transferProtocol := flag.String("transportProtocol", "quic", "tcp/quic")
	addr := flag.String("serverAddr", "https://localhost:8090", "server address")
	concurrency := flag.Int("concurrency", 1, "number > 0")
	flag.Parse()

	// Get the SystemCertPool, continue with an empty pool on error
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := os.ReadFile(certFile)
	if err != nil {
		log.Fatalf("Failed to append %q to RootCAs: %v", certFile, err)
	}

	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		log.Println("No certs appended, using system certs only")
	}

	var qconf quic.Config
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			RootCAs:            rootCAs,
		},
		QuicConfig: &qconf,
	}
	defer roundTripper.Close()
	hclient := &http.Client{
		Transport: roundTripper,
	}

	ipString := strings.Split(*addr, ":")[0]
	port, err := strconv.Atoi(strings.Split(*addr, ":")[1])
	if err != nil {
		log.Fatal(err)
	}
	ip, err := net.ResolveIPAddr("ip", ipString)
	if err != nil {
		log.Fatalf("invalid IP: %v", err)
	}
	sctpAddr := &sctp.SCTPAddr{
		IPAddrs: []net.IPAddr{*ip},
		Port:    port,
	}
	now := time.Now()
	if *transferProtocol == "quic" {
		log.Infof("Received msg of length %d", len(getWithQuic("https://"+*addr, *hclient)))
	} else if *transferProtocol == "http" || *transferProtocol == "fasthttp" {
		log.Infof("Received msg of length %d", len(getWithHttp("http://"+*addr)))
	} else if *transferProtocol == "grpc" {
		log.Infof("Received msg of length %d", len(getWithGrpc(*addr)))
	} else if *transferProtocol == "kcp" {
		log.Infof("Received msg of length %d", len(getWithKcp(*addr)))
	} else if *transferProtocol == "tcp" {
		log.Infof("Received msg of length %d", len(getWithTcp(*addr)))
	} else if *transferProtocol == "tcpmarshall" {
		log.Infof("Received msg of length %d", len(getWithTcpMarshal(*addr)))
	} else if *transferProtocol == "sctp" {
		log.Infof("Received msg of length %d", len(getWithSctp(sctpAddr)))
	} else if *transferProtocol == "capnp" {
		var channel = make(chan []byte)
		for i := 0; i < *concurrency; i += 1 {
			go func() {
				channel <- getWithCapnp(*addr)
			}()
		}
		for i := 0; i < *concurrency; i += 1 {
			log.Infof("Received msg of length %d", len(<-channel))
		}
	} else if *transferProtocol == "flatbuf" {
		log.Infof("Received msg of length %d", len(getWithFlatbuf(*addr)))
	} else {
		log.Fatal("Invalid transfer type")
	}

	log.Infof("%v", time.Since(now))
}
