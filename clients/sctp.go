package main

import (
	"github.com/ishidawataru/sctp"
	log "github.com/sirupsen/logrus"
)

func getWithSctp(sctpAddr *sctp.SCTPAddr) []byte {
	var laddr *sctp.SCTPAddr
	buf := make([]byte, 10240)
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
	return buf
}
