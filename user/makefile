proto:
	protoc --go_out=../gateway/pkg/user --go_opt=MProtos/user.proto=./pb \
	--go-grpc_out=../gateway/pkg/user --go-grpc_opt=MProtos/user.proto=./pb Protos/*.proto
	protoc --go_out=../auth/pkg/ --go_opt=MProtos/user.proto=./pb \
	--go-grpc_out=../auth/pkg/ --go-grpc_opt=MProtos/user.proto=./pb Protos/*.proto
