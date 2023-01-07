package main

import (
	"capnproto.org/go/capnp/v3"
	capnp_proto "github.com/shyamjesal/transfer-bench/capnp"
	log "github.com/sirupsen/logrus"
	"net"
)

func getWithTcp(addr string) []byte {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 10240)
	conn.Write([]byte("bla\n"))
	//_, err = io.ReadFull(conn, buf)
	_, err = conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return buf
}
func getWithTcpMarshal(addr string) []byte {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write([]byte("bla\n"))
	buf := make([]byte, 20240)
	conn.Read(buf)
	msg, err := capnp.Unmarshal(buf)
	if err != nil {
		panic(err)
	}

	// Again, don't worry about the meaning of "root" for now.
	// When in doubt, use the "root" version of functions.
	book, err := capnp_proto.ReadRootBook(msg)
	if err != nil {
		panic(err)
	}

	title, _ := book.Title()
	return title
}
