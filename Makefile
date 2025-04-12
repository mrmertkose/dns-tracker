APP_NAME := $(notdir $(CURDIR))
BUILD_DIR := bin

build:
	go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME) .

build-linux:
	env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME)-linux .

build-windows:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o $(BUILD_DIR)/$(APP_NAME).exe .

run:
	go run .

.PHONY: build build-linux build-windows run
