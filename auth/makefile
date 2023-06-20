proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/pb/*.proto && \
	protoc --go_out=../gateway/pkg/auth --go_opt=paths=import \
	--go-grpc_out=../gateway/pkg/auth --go-grpc_opt=paths=import pkg/pb/*.proto
server:
	cd cmd && go run main.go
