#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

OS = $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH = $(shell uname -m)
GIT_COMMIT = $(shell git rev-parse --short HEAD)
GIT_REPO_NAME = $(shell basename -s .git `git config --get remote.origin.url`)
GIT_OWNER = $(shell git config --get remote.origin.url | cut -d':' -f2 | cut -d'/' -f1)
GIT_SHA = $(shell git rev-parse HEAD)
SERVICE_NAME = solar-panel-api

.PHONY: clean build test

all: clean build

clean:
	rm -rf builds

build: clean
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) go build -a -installsuffix cgo -ldflags="-X main.Version=$(GIT_COMMIT)" -o builds/$(SERVICE_NAME)-$(OS)-$(ARCH)
