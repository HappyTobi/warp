GOCMD=go
VERSION?=0.0.1
BINARY_NAME=warp

.PHONY: build

build:
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME) main.go

release:
	mkdir -p build/darwin
	mkdir -p build/darwin-arm
	mkdir -p build/linux
	mkdir -p build/linux-arm
	mkdir -p build/windows

	GOARCH=amd64 GOOS=darwin $(GOCMD) build -o ./build/darwin/${BINARY_NAME} main.go
	GOARCH=arm64 GOOS=darwin $(GOCMD) build -o ./build/darwin-arm/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=linux $(GOCMD) build -o ./build/linux/${BINARY_NAME} main.go
	GOARCH=arm64 GOOS=linux $(GOCMD) build -o ./build/linux-arm/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=windows $(GOCMD) build -o ./build/windows/${BINARY_NAME} main.go

clean:
	rm -rf ./build

vendor:
	$(GOCMD) mod vendor

test:
	$(GOCMD) test ./...