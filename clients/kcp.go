package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go/v5"
)

func getWithKcp(addr string) []byte {
	// dial to the echo server
	buf := make([]byte, 10240)
	if sess, err := kcp.Dial(addr); err == nil {
		// read back the data
		sess.Write([]byte("bla"))
		if _, err := sess.Read(buf); err == nil {
			//log.Println("recv:", string(buf))
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
	return buf
}
