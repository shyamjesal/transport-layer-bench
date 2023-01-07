package main

import (
	"capnproto.org/go/capnp/v3/rpc"
	"context"
	capnp_proto "github.com/shyamjesal/transfer-bench/capnp"
	log "github.com/sirupsen/logrus"
	"net"
)

func getWithCapnp(addr string) []byte {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	capconn := rpc.NewConn(rpc.NewStreamTransport(conn), nil)
	defer capconn.Close()

	// Now we resolve the bootstrap interface from the remote ArithServer.
	// Thanks to Cap'n Proto's promise pipelining, this function call does
	// NOT block.  We can start making RPC calls with 'a' immediately, and
	// these will transparently resolve when bootstrapping completes.
	//
	// The context can be used to time-out or otherwise abort the bootstrap
	// call.   It is safe to cancel the context after the first method call
	// on 'a' completes.
	ctx := context.Background()
	a := capnp_proto.Packet(capconn.Bootstrap(ctx))

	// Okay! Let's make an RPC call!  Remember:  RPC is performed simply by
	// calling a's methods.
	//
	// There are a couple of interesting things to note here:
	//  1. We pass a callback function to set parameters on the RPC call.  If the
	//     call takes no arguments, you MAY pass nil.
	//  2. We return a Future type, representing the in-flight RPC call.  As with
	//     the earlier call to Bootstrap, a's methods do not block.  They instead
	//     return a future that eventually resolves with the RPC results. We also
	//     return a release function, which MUST be called when you're done with
	//     the RPC call and its results.
	f, release := a.Get(ctx, func(ps capnp_proto.Packet_get_Params) error {
		ps.SetKey("bla")
		return nil
	})
	defer release()

	// You can do other things while the RPC call is in-flight.  Everything
	// is asynchronous. For simplicity, we're going to block until the call
	// completes.
	res, err := f.Struct()
	if err != nil {
		log.Fatal(err)
	}

	// Lastly, let's print the result.  Recall that 'product' is the name of
	// the return value that we defined in the schema file.
	payload, err := res.Payload()
	if err != nil {
		log.Fatal(err)
	}
	return payload
}
