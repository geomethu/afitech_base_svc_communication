generate:
	protoc --proto_path=api/proto --go_out=:internal/delivery/grpc/v1 --go-grpc_out=:internal/delivery/grpc/v1 api/proto/*.proto

gen:
	protoc --proto_path=api/proto --go_out=:. --go-grpc_out=:. api/proto/*.proto
	
run:
	go run ./cmd/main.go