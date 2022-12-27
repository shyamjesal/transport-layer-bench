package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/xtaci/kcp-go/v5"
)

//
//func getWithKcp(addr string) []byte {
//	// dial to the echo server
//	buf := make([]byte, 10240)
//	if sess, err := kcp.Dial(addr); err == nil {
//		log.Println("sent: hi")
//		if _, err := sess.Write([]byte("hi")); err == nil {
//			// read back the data
//			if _, err := io.ReadFull(sess, buf); err == nil {
//				//log.Println("recv:", string(buf))
//			} else {
//				log.Fatal(err)
//			}
//		} else {
//			log.Fatal(err)
//		}
//	} else {
//		log.Fatal(err)
//	}
//	return buf
//}

func getWithKcp(addr string) []byte {
	// dial to the echo server
	buf := make([]byte, 10240)
	if sess, err := kcp.Dial(addr); err == nil {
		// read back the data
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
