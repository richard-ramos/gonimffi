.PHONY: build

ANDROID_TARGET ?= 30
GOBIN ?= $(shell which go)


ifeq ($(OS),Windows_NT)     # is Windows_NT on XP, 2000, 7, Vista, 10...
 detected_OS := Windows
else
 detected_OS := $(strip $(shell uname))
endif

ifeq ($(detected_OS),Darwin)
 GOBIN_SHARED_LIB_EXT := dylib
else ifeq ($(detected_OS),Windows)
 # on Windows need `--export-all-symbols` flag else expected symbols will not be found in the .dll
 GOBIN_SHARED_LIB_CGO_LDFLAGS := CGO_LDFLAGS="-Wl,--export-all-symbols"
 GOBIN_SHARED_LIB_EXT := dll
else
 GOBIN_SHARED_LIB_EXT := so
endif

all: | build-mobile build-desktop

build-mobile:
	@echo "Android target: ${ANDROID_TARGET} (override with ANDROID_TARGET var)"
	gomobile init && \
	${GOBIN} get -d golang.org/x/mobile/cmd/gomobile && \
	CGO=1 gomobile bind -v -target=android -androidapi=${ANDROID_TARGET} -ldflags="-s -w" -o ./build/hello.aar ./mobile && \
	mkdir -p ../mobile/AwesomeProject/android/app/libs/ && \
	rm -Rf ../mobile/AwesomeProject/android/app/libs/* && \
	cp build/*.aar ../mobile/AwesomeProject/android/app/libs/.

build-desktop:
	$(GOBIN_SHARED_LIB_CFLAGS) $(GOBIN_SHARED_LIB_CGO_LDFLAGS) ${GOBIN} build \
		-buildmode=c-shared \
		-o ./build/libgohello.$(GOBIN_SHARED_LIB_EXT) ./desktop/
ifeq ($(detected_OS),Darwin)
	sed -i '' -e "s/#include \"../libhello.h\"//gi" ./build/libgohello.h
else
	sed -i 's/#include "..\/libhello.h"//gi' ./build/libgohello.h
endif
	cp build/libgohello.$(GOBIN_SHARED_LIB_EXT) ../desktop/.
	cp build/libgohello.h ../desktop/.
