OS := $(shell uname)

ifeq ($(OS), Windows_NT)
	BIN_PATH = bin\api.exe
else
	BIN_PATH = ./bin/api
endif

run: build
	@$(BIN_PATH)

build:
	@go build -o $(BIN_PATH)

test:
	@go test -v ./...
