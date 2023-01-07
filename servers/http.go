package main

import (
	"github.com/JoeReid/fastTCP"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
)

func getWithHttp(w http.ResponseWriter, req *http.Request) {
	log.Info("Server received a request")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write(payloadData)
}
func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	ctx.Write(payloadData)
}

func serveWithFastHttp() {
	handler := func(tcpFile io.ReadWriter) {
		log.Info("Processing using fasttcp")
		_, err := tcpFile.Write(payloadData)
		if err != nil {
			log.Fatal(err)
		}
	}
	server := fastTCP.NewServer(addr, handler, fastTCP.TCPOptions{FastOpen: true})

	// Start the TCP server
	err := server.ListenTCP()
	if err != nil {
		panic(err)
	}
}
