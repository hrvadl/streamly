proto:
	protoc --go_out=../gateway/pkg/user --go_opt=MProtos/user.proto=./pb \
	--go-grpc_out=../gateway/pkg/user --go-grpc_opt=MProtos/user.proto=./pb Protos/*.proto

server:
	go run cmd/main.go
