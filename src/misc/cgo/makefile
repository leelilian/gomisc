GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=cgo
BINARY_UNIX=$(BINARY_NAME)_unix

all : build test
build :
	${GOBUILD}

test :
	./${BINARY_NAME}

clean :
	rm -rf ${BINARY_NAME} *.o *.a
