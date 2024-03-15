.PHONY: build

ANDROID_TARGET ?= 30
GOBIN ?= $(shell which go)

build:
	@echo "Android target: ${ANDROID_TARGET} (override with ANDROID_TARGET var)"
	gomobile init && \
	${GOBIN} get -d golang.org/x/mobile/cmd/gomobile && \
	CGO=1 gomobile bind -v -target=android -androidapi=${ANDROID_TARGET} -ldflags="-s -w" -o ./build/hello.aar ./library && \
	mkdir -p ../mobile/AwesomeProject/android/app/libs/ && \
	rm -Rf ../mobile/AwesomeProject/android/app/libs/* && \
	cp build/*.aar ../mobile/AwesomeProject/android/app/libs/.