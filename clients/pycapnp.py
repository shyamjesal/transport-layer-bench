#!/usr/bin/env python3

import argparse
import asyncio
import logging
import sys
import time

import capnp
capnp.remove_import_hook()
transfer_capnp = capnp.load('./transfer.capnp')

logging.basicConfig(stream=sys.stdout, level=logging.INFO)


def parse_args():
    parser = argparse.ArgumentParser(
        usage="Connects to the Calculator server \
at the given address and does some RPCs"
    )
    parser.add_argument("host", help="HOST:PORT")

    return parser.parse_args()


def main(host):
    logging.info("connecting to host %s", host)
    start = time.time()
    promises = []
    for i in range(10):
        client = capnp.TwoPartyClient(host)

        # Bootstrap the server capability and cast it to the Calculator interface
        packet = client.bootstrap().cast_as(transfer_capnp.Packet)

        """Make a request that just evaluates the literal value 123.
    
        What's interesting here is that evaluate() returns a "Value", which is
        another interface and therefore points back to an object living on the
        server.  We then have to call read() on that object to read it.
        However, even though we are making two RPC's, this block executes in
        *one* network round trip because of promise pipelining:  we do not wait
        for the first call to complete before we send the second call to the
        server."""

        logging.debug("Getting from server... ")

        # Make the request. Note we are using the shorter function form (instead
        # of evaluate_request), and we are passing a dictionary that represents a
        # struct and its member to evaluate
        # get_promise = packet.get({"key": "bla"})

        # This is equivalent to:

        request = packet.get_request()
        request.key = "bla"

        # Send it, which returns a promise for the result (without blocking).
        get_promise = request.send()
        print(type(get_promise))
        promises.append(get_promise)

        # Using the promise, create a pipelined request to call read() on the
        # returned object. Note that here we are using the shortened method call
        # syntax read(), which is mostly just sugar for read_request().send()
        # read_promise = get_promise.key.read()

        # Now that we've sent all the requests, wait for the response.  Until this
        # point, we haven't waited at all!
        response = get_promise.wait()
        logging.info("received from server %d", len(response.payload))
    # await asyncio.gather(promises)
    print("This took ", time.time()-start)


if __name__ == "__main__":
    main(parse_args().host)
