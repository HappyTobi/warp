GOCMD=go
VERSION?=0.0.1
VERSION_TRIMMED := $(VERSION:v%=%)
TAR ?= tar
BINARY_NAME=warp

linker_flags = '-s -X github.com/HappyTobi/warp/pkg/cmd/version.version=${VERSION}'

.PHONY: build

build:
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME) main.go

clean:
	rm -rf build/output

release:
	mkdir -p build/artifacts

	make clean
	GOARCH=amd64 GOOS=darwin $(GOCMD) build -ldflags=${linker_flags} -o ./build/output/${BINARY_NAME} main.go
	$(TAR) -C ./build/output -czvf ./build/artifacts/warp-$(VERSION_TRIMMED)-Darwin-x86_64.tar.gz .
	make clean
	GOARCH=arm64 GOOS=darwin $(GOCMD) build -ldflags=${linker_flags} -o ./build/output/${BINARY_NAME} main.go
	$(TAR) -C ./build/output -czvf ./build/artifacts/warp-$(VERSION_TRIMMED)-Darwin-arm64.tar.gz .
	make clean
	GOARCH=amd64 GOOS=linux $(GOCMD) build -ldflags=${linker_flags} -o ./build/output/${BINARY_NAME} main.go
	$(TAR) -C ./build/output -czvf ./build/artifacts/warp-$(VERSION_TRIMMED)-Linux-x86_64.tar.gz .
	make clean
	GOARCH=arm64 GOOS=linux $(GOCMD) build -ldflags=${linker_flags} -o ./build/output/${BINARY_NAME} main.go
	$(TAR) -C ./build/output -czvf ./build/artifacts/warp-$(VERSION_TRIMMED)-Linux-aarch64.tar.gz .
	make clean
	GOARCH=amd64 GOOS=windows $(GOCMD) build -ldflags=${linker_flags} -o ./build/output/${BINARY_NAME}.exe main.go
	$(TAR) -C ./build/output -czvf ./build/artifacts/warp-$(VERSION_TRIMMED)-Windows-x86_64.tar.gz .
	make clean
lint:
	golangci-lint run ./...

vendor:
	$(GOCMD) mod vendor

test:
	$(GOCMD) test ./...