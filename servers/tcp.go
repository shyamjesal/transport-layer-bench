package main

import (
	"bufio"
	"capnproto.org/go/capnp/v3"
	capnp_proto "github.com/shyamjesal/transfer-bench/capnp"
	log "github.com/sirupsen/logrus"
	"net"
)

func serveWithTcp() {
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// accept connection
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	buffer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Info(buffer)
	// send file to client
	_, err = conn.Write(payloadData)
	if err != nil {
		log.Fatal(err)
	}
}

func serveWithTcpMarshal() {
	server, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		// accept connection
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(conn)
		buffer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		log.Info(buffer)
		//prepare marshaled packet
		arena := capnp.SingleSegment(nil)
		msg, seg, err := capnp.NewMessage(arena)
		if err != nil {
			panic(err)
		}
		book, err := capnp_proto.NewRootBook(seg)
		_ = book.SetTitle(payloadData)

		// Then, we set the page count.
		book.SetPageCount(1440)
		if err != nil {
			panic(err)
		}
		b, err := msg.Marshal()
		if err != nil {
			panic(err)
		}
		// send file to client
		_, err = conn.Write(b)
		if err != nil {
			log.Fatal(err)
		}
	}
}
