package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	certFile = "../certs/localhost.crt"
	keyFile  = "../certs/localhost.key"
)

func getWithQuic(addr string, hclient http.Client) []byte {
	rsp, err := hclient.Get(addr)
	if err != nil {
		log.Fatal(err)
	}
	//log.Infof("Got response for %s: %#v", addr, rsp)

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
