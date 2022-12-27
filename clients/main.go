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
	buf := make([]byte, 10240)
	now := time.Now()
	if *transferProtocol == "quic" {
		log.Infof("Received msg of length %d", len(getWithQuic(*addr, *hclient)))
	} else if *transferProtocol == "http" {
		log.Infof("Received msg of length %d", len(getWithHttp(*addr)))
	} else if *transferProtocol == "grpc" {
		log.Infof("Received msg of length %d", len(getWithGrpc(*addr)))
	} else if *transferProtocol == "kcp" {
		log.Infof("Received msg of length %d", len(getWithKcp(*addr)))
	} else if *transferProtocol == "tcp" {
		conn, err := net.Dial("tcp", *addr)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		buf := make([]byte, 10240)

		_, err = conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
	} else if *transferProtocol == "sctp" {
		var laddr *sctp.SCTPAddr
		conn, err := sctp.DialSCTP("sctp", laddr, sctpAddr)
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		defer conn.Close()
		//sndbuf, err := conn.GetWriteBuffer()
		//if err != nil {
		//	log.Fatalf("failed to get write buf: %v", err)
		//}
		//rcvbuf, err := conn.GetReadBuffer()
		//if err != nil {
		//	log.Fatalf("failed to get read buf: %v", err)
		//}
		//log.Printf("SndBufSize: %d, RcvBufSize: %d", sndbuf, rcvbuf)

		n, _, err := conn.SCTPRead(buf)
		if err != nil {
			log.Fatalf("failed to read: %v", err)
		}
		log.Printf("read: len %d", n)
	}
	log.Infof("%v", time.Since(now))
}
