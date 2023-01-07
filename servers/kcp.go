package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go/v5"
)

func serveWithKcp() {
	if listener, err := kcp.Listen(addr); err == nil {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal(err)
			}
			buffer := make([]byte, 10)
			conn.Read(buffer)
			_, err = conn.Write(payloadData)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal(err)
	}
}
