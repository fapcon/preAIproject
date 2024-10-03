TARGET=auth
PROTO_NAME=pb/sso/sso.proto

all: clean build

clean:
	rm -rf $(TARGET)

build:
	go build -o $(TARGET) main.go

proto:
	buf generate

gateway:
	protoc -I ./pb --grpc-gateway_out ./pb \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	pb/sso/sso.proto
