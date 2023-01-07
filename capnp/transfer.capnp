using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("capnp");
$Go.import("github.com/shyamjesal/transfer-bench/capnp");

# Declare the Arith capability, which provides multiplication and division.
interface Packet {
	get @0 (key :Text) -> (payload :Data);
}

struct Book {
    title @0 :Data;
    # Title of the book.

    pageCount @1 :Int32;
    # Number of pages in the book.
}