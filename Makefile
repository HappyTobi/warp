GOCMD=go
VERSION?=0.0.1
BINARY_NAME=warp

.PHONY: build

build:
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME) main.go

release:
	mkdir -p build

 	GOARCH=amd64 GOOS=darwin 	$(GOCMD) build -o ./build/${BINARY_NAME}-darwin 			main.go
 	GOARCH=arm64 GOOS=darwin 	$(GOCMD) build -o ./build/${BINARY_NAME}-darwin-silicon 	main.go
 	GOARCH=amd64 GOOS=linux 	$(GOCMD) build -o ./build/${BINARY_NAME}-linux 				main.go
 	GOARCH=arm64 GOOS=linux 	$(GOCMD) build -o ./build/${BINARY_NAME}-linux-arm 			main.go
 	GOARCH=amd64 GOOS=windows 	$(GOCMD) build -o ./build/${BINARY_NAME}-windows 			main.go

clean:
	rm -rf ./build

vendor:
	$(GOCMD) mod vendor

test:
	$(GOCMD) test ./...