package main

import (
	"github.com/ishidawataru/sctp"
	log "github.com/sirupsen/logrus"
	"net"
)

func serveWithSctp(ipString string, port int) {
	ip, err := net.ResolveIPAddr("ip", ipString)
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
