
all: linux compress

linux:
	MSYS_NO_PATHCONV=1 docker run --rm -e GOPROXY="https://goproxy.io,direct" -v $(dir $(abspath $(lastword $(MAKEFILE_LIST)))):/usr/src/myapp -w /usr/src/myapp golang:latest go build -ldflags '-s -w -extldflags "-static"' -o target/xxqg-tk

compress:
	upx --lzma target/*