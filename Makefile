GOCMD=go
VERSION?=0.0.1
BINARY_NAME=warp

linker_flags = '-s -X github.com/HappyTobi/warp/pkg/cmd/version.version=${VERSION}'

.PHONY: build

build:
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME) main.go

release:
	mkdir -p build

	GOARCH=amd64 GOOS=darwin $(GOCMD) build -ldflags=${linker_flags} -o ./build/darwin/${BINARY_NAME} main.go
	GOARCH=arm64 GOOS=darwin $(GOCMD) build -ldflags=${linker_flags} -o ./build/darwin/arm/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=linux $(GOCMD) build -ldflags=${linker_flags} -o ./build/linux/${BINARY_NAME} main.go
	GOARCH=arm64 GOOS=linux $(GOCMD) build -ldflags=${linker_flags} -o ./build/linux/arm/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=windows $(GOCMD) build -ldflags=${linker_flags} -o ./build/windows/${BINARY_NAME}.exe main.go

lint:
	golangci-lint run ./...

clean:
	rm -rf ./build

vendor:
	$(GOCMD) mod vendor

test:
	$(GOCMD) test ./...