PROTO_FILES:=hello
proto_gen: $(PROTO_FILES)
ROOT = ./

$(PROTO_FILES):
	GO111MODULE="on" PATH="$$PATH:$$(go env GOPATH)/bin" protoc proto/$@.proto \
	    --proto_path=./proto \
	    --go_out=./proto/$@ --go-grpc_out=./proto/$@ \
	    --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative

all: all-image 

all-image: client-image server-image

all-image-push: client-image-push server-image-push

client-image: Dockerfile clients/main.go
	DOCKER_BUILDKIT=1 docker build \
	-t vhiveease/transport-layer-bench-client:latest \
	--target client \
	-f Dockerfile \
	$(ROOT)

server-image: Dockerfile servers/main.go
	DOCKER_BUILDKIT=1 docker build \
	-t vhiveease/transport-layer-bench-server:latest \
	--target server \
	-f Dockerfile \
	$(ROOT)

client-image-push: client-image
	docker push vhiveease/transport-layer-bench-client:latest

server-image-push: server-image
	docker push vhiveease/transport-layer-bench-server:latest
