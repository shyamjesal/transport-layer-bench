FROM golang:1.18-alpine as server
WORKDIR /app

COPY servers ./servers
COPY certs/localhost.crt ./certs/localhost.crt
COPY certs/localhost.key ./certs/localhost.key
COPY proto ./proto
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

RUN go build servers/main.go servers/http.go servers/grpc.go servers/kcp.go

CMD [ "main" ]

FROM golang:1.18-alpine as client
WORKDIR /app

COPY clients ./clients
COPY certs/localhost.crt ./certs/localhost.crt
COPY certs/localhost.key ./certs/localhost.key
COPY proto ./proto
COPY go.mod ./go.mod
COPY go.sum ./go.sum
RUN go mod download

RUN go build clients/main.go clients/quic.go clients/http.go clients/grpc.go clients/kcp.go

CMD [ "main" ]
