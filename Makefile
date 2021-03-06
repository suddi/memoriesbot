BUILD_PATH := build

GOPATH := $(shell go env GOPATH)
BIN_NAME := memoriesbot

.PHONY: build compile zip test coverage builddir install clean

build: .
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_PATH)/main ./cmd

compile:
	go build -o $(BUILD_PATH)/main ./cmd

zip:
	zip -j $(BUILD_PATH)/$(BIN_NAME).zip $(BUILD_PATH)/main

test:
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out

builddir:
	mkdir -p $(BUILD_PATH)

install:
	go mod download

clean:
	rm -rf $(BUILD_PATH)
