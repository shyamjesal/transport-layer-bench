module transfer-bench

go 1.19

require (
	capnproto.org/go/capnp/v3 v3.0.0-alpha.17
	github.com/JoeReid/fastTCP v0.0.0-20170128213645-970c676a0702
	github.com/ishidawataru/sctp v0.0.0-20210707070123-9a39160e9062
	github.com/lucas-clemente/quic-go v0.31.1
	github.com/shyamjesal/transfer-bench/capnp v0.0.0-00010101000000-000000000000
	github.com/shyamjesal/transfer-bench/flatbuf v0.0.0-00010101000000-000000000000
	github.com/shyamjesal/transfer-bench/proto v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	github.com/valyala/fasthttp v1.43.0
	github.com/xtaci/kcp-go/v5 v5.6.1
	google.golang.org/grpc v1.51.0
)

replace (
	github.com/shyamjesal/transfer-bench/capnp => ./capnp
	github.com/shyamjesal/transfer-bench/flatbuf => ./flatbuf/flatMsg
	github.com/shyamjesal/transfer-bench/proto => ./proto/hello
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-task/slim-sprig v0.0.0-20210107165309-348f09dbbbc0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/flatbuffers v22.11.23+incompatible // indirect
	github.com/google/pprof v0.0.0-20210407192527-94a9f03dee38 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid v1.3.1 // indirect
	github.com/klauspost/reedsolomon v1.9.9 // indirect
	github.com/marten-seemann/qpack v0.3.0 // indirect
	github.com/marten-seemann/qtls-go1-18 v0.1.3 // indirect
	github.com/marten-seemann/qtls-go1-19 v0.1.1 // indirect
	github.com/mmcloughlin/avo v0.0.0-20200803215136-443f81d77104 // indirect
	github.com/onsi/ginkgo/v2 v2.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/templexxx/cpu v0.0.7 // indirect
	github.com/templexxx/xorsimd v0.4.1 // indirect
	github.com/tjfoc/gmsm v1.3.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292 // indirect
	golang.org/x/exp v0.0.0-20220722155223-a9213eeb770e // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/net v0.0.0-20220906165146-f3363e06e74c // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.1.1-0.20221102194838-fc697a31fa06 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
