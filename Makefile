export PATH := $(GOPATH)/bin
BINARY_NAME=sfs

build:
	go build -o ${BINARY_NAME} main.go

install:
	go build -o ${PATH}/${BINARY_NAME} main.go