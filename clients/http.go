package main

import (
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func getWithHttp(addr string) []byte {
	rsp, err := http.Get(addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Infof("Got response for %s: %#v", addr, rsp)

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
