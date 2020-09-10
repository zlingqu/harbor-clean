BINARY_NAME=harbor-clean
BINARY_UNIX=$(BINARY_NAME)_unix

# Go parameters
GOCMD := CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go
GOBUILD=$(GOCMD) build
GOBUILDWIN=$(GOCMDWIN) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
DIR=targets

all: run

build:
	$(GOBUILD) -o ./$(DIR)/$(BINARY_NAME) -v ./

test:
	$(GOTEST) -v ./

clean:
	rm -rf ./$(DIR)/$(BINARY_NAME)*

compile:
	$(GOBUILD) -o ./$(DIR)/$(BINARY_NAME) -v ./

run:
	cd ./$(DIR) && ./$(BINARY_NAME)
