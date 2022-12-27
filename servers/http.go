package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
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
