PROJECT=cluster-client

BUILD_PATH := $(shell pwd)/.gobuild
VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)

.PHONY=all get-deps build clean

PROJECT_PATH := "$(BUILD_PATH)/src/github.com/giantswarm"

GOPATH := $(BUILD_PATH)

SOURCE=$(shell find . -name '*.go')

BIN := $(PROJECT)

ifndef GOOS
	GOOS := $(shell go env GOOS)
endif
ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
endif

all: .gobuild get-deps $(BIN)

get-deps: .gobuild
	GOPATH=$(GOPATH) go get -d -v github.com/giantswarm/$(PROJECT)

.gobuild:
	mkdir -p $(PROJECT_PATH)
	cd "$(PROJECT_PATH)" && ln -s ../../../.. $(PROJECT)

$(BIN): VERSION $(SOURCE)
	echo Building for $(GOOS)/$(GOARCH)
	docker run \
	    --rm \
	    -v $(shell pwd):/usr/code \
	    -e GOPATH=/usr/code/.gobuild \
	    -e GOOS=$(GOOS) \
	    -e GOARCH=$(GOARCH) \
	    -w /usr/code \
	    golang:1.6 \
	    go build -a -ldflags "-X main.projectVersion=$(VERSION) -X main.projectBuild=$(COMMIT)" -o $(BIN)

clean:
	rm -rf $(BUILD_PATH) $(BIN)

fmt:
	gofmt -w -l .
