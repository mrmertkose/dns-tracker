APP_NAME := $(notdir $(CURDIR))
GO := $(shell which go)
BUILD_DIR := bin

build:
	go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

build-windows:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME).exe .

run:
	sudo $(GO) run .

.PHONY: build build-linux build-windows run
