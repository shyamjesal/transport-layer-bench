package main

import (
	log "github.com/sirupsen/logrus"
	"net"
)

// handleEcho send back everything it received
func handleKcp(conn net.Conn) {
	for {
		_, err := conn.Write(payloadData)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
