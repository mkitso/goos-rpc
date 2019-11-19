.PHONY: proto server client

proto:
	@protoc -I os/ os/os.proto --go_out=plugins=grpc:os

server:
	@go run server/os.go

client:
	@go run client/os.go
