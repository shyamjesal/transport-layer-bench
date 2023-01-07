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

python-server-run: python-server-build
	docker run shyamjesal/transfer-bench-python-server:latest python3 capnp_server.py localhost:8090

python-server-build:
	DOCKER_BUILDKIT=1 docker build \
    	-t shyamjesal/transfer-bench-python-server:latest \
    	--target pyserver \
    	-f Dockerfile \
    	$(ROOT)

python-client-run: python-client-build
	docker run shyamjesal/transfer-bench-python-client:latest python3 capnp_client.py localhost:8090

python-client-build:
	DOCKER_BUILDKIT=1 docker build \
    	-t shyamjesal/transfer-bench-python-client:latest \
    	--target pyclient \
    	-f Dockerfile \
    	$(ROOT)

python-build: python-client-build python-server-build

python-test: python-build
	docker compose -f docker-compose.yml up