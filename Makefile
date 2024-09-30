gen:
	protoc --go_out=. --go-grpc_out=. calculator.proto

build:
	go build -o bin/calculator ./main.go

run: build
	./bin/calculator
