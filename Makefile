gen:
	protoc --go_out=. --go-grpc_out=. calculator.proto

build:
	go build -o bin/calculator ./main.go

run: build
	./bin/calculator

ui:
	make run &
	sleep 0.5
	grpcui --plaintext 127.0.0.1:50051
