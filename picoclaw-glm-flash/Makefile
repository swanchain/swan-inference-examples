# Build for current platform
build:
	go build -o swan-chat .

# Cross-compile for edge devices
build-arm:
	GOOS=linux GOARCH=arm64 go build -o swan-chat-arm64 .
	GOOS=linux GOARCH=arm GOARM=7 go build -o swan-chat-armv7 .

build-riscv:
	GOOS=linux GOARCH=riscv64 go build -o swan-chat-riscv64 .

# Build all targets
build-all: build build-arm build-riscv
	ls -lh swan-chat*

run:
	SWAN_API_KEY=$${SWAN_API_KEY} ./swan-chat "$(prompt)"

clean:
	rm -f swan-chat swan-chat-*
