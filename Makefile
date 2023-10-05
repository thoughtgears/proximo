#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH = $(shell uname -m)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
SERVICE_NAME = proximo

.PHONY: clean build test

all: clean build

clean:
	rm -rf builds

build:
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -a -installsuffix cgo -ldflags="-X main.Version=$(GIT_COMMIT)" -o builds/$(SERVICE_NAME)-$(OS)-$(ARCH)
