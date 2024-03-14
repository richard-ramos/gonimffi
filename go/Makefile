ANDROID_TARGET ?= 30
GOBIN ?= $(shell which go)

# CC=/home/richard/Android/Sdk/ndk/21.4.7075529/toolchains/llvm/prebuilt/linux-x86_64/bin/aarch64-linux-android30-clang CGO_ENABLED=1 GOOS=android GOARCH=arm64 $(GOBIN) build -v -buildmode=c-shared -o ./libgohello.so . 

build:
	@echo "Android target: ${ANDROID_TARGET} (override with ANDROID_TARGET var)"
	gomobile init && \
	${GOBIN} get -d golang.org/x/mobile/cmd/gomobile && \
	CGO=1 gomobile bind -v -target=android -androidapi=${ANDROID_TARGET} -ldflags="-s -w" -o ./hello.aar ./library
