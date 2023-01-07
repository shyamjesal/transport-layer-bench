#!/usr/bin/env python3

import argparse
import asyncio
import logging
import os
import socket
import sys

import capnp
capnp.remove_import_hook()
transfer_capnp = capnp.load('./transfer.capnp')
logging.basicConfig(stream=sys.stdout, level=logging.INFO)


class Server:
    async def myreader(self):
        while self.retry:
            try:
                # Must be a wait_for so we don't block on read()
                data = await asyncio.wait_for(self.reader.read(4096), timeout=0.1)
            except asyncio.TimeoutError:
                logging.debug("myreader timeout.")
                continue
            except Exception as err:
                logging.error("Unknown myreader err: %s", err)
                return False
            await self.server.write(data)
        logging.debug("myreader done.")
        return True

    async def mywriter(self):
        while self.retry:
            try:
                # Must be a wait_for so we don't block on read()
                data = await asyncio.wait_for(self.server.read(4096), timeout=0.1)
                self.writer.write(data.tobytes())
            except asyncio.TimeoutError:
                logging.debug("mywriter timeout.")
                continue
            except Exception as err:
                logging.error("Unknown mywriter err: %s", err)
                return False
        logging.debug("mywriter done.")
        return True

    async def myserver(self, reader, writer):
        # Start TwoPartyServer using TwoWayPipe (only requires bootstrap)
        self.server = capnp.TwoPartyServer(bootstrap=PacketImpl())
        self.reader = reader
        self.writer = writer
        self.retry = True

        # Assemble reader and writer tasks, run in the background
        coroutines = [self.myreader(), self.mywriter()]
        tasks = asyncio.gather(*coroutines, return_exceptions=True)

        while True:
            self.server.poll_once()
            # Check to see if reader has been sent an eof (disconnect)
            if self.reader.at_eof():
                self.retry = False
                break
            await asyncio.sleep(0.01)

        # Make wait for reader/writer to finish (prevent possible resource leaks)
        await tasks


class PacketImpl(transfer_capnp.Packet.Server):

    """Implementation of the Calculator.Function Cap'n Proto interface, where the
    function is defined by a Calculator.Expression."""

    def __init__(self):
        self.payloadData = bytes(os.urandom(1024 * 1024 * 10))

    def get(self, key, _context, **kwargs):
        """Note that we're returning a Promise object here, and bypassing the
        helper functionality that normally sets the results struct from the
        returned object. Instead, we set _context.results directly inside of
        another promise"""
        logging.info("server received the key %s", key)
        return self.payloadData


def parse_args():
    parser = argparse.ArgumentParser(
        usage="""Runs the server bound to the\
given address/port ADDRESS. """
    )

    parser.add_argument("address", help="ADDRESS:PORT")

    return parser.parse_args()


async def new_connection(reader, writer):
    server = Server()
    await server.myserver(reader, writer)


async def main():
    address = parse_args().address
    host = address.split(":")
    addr = host[0]
    port = host[1]

    # Handle both IPv4 and IPv6 cases
    try:
        logging.info("Try IPv4 on %s %s", addr, port)
        server = await asyncio.start_server(
            new_connection, addr, port, family=socket.AF_INET
        )
    except Exception:
        logging.info("Try IPv6 on %s %s", addr, port)
        server = await asyncio.start_server(
            new_connection, addr, port, family=socket.AF_INET6
        )

    async with server:
        await server.serve_forever()


if __name__ == "__main__":
    asyncio.run(main())
