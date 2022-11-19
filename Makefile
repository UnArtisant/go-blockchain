BINARY_NAME=bin
build:
	go build -o ./${BINARY_NAME}/blockchain
run:
	./${BINARY_NAME}/blockchain
test:
	go test -v ./...