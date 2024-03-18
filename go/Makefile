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

all: | build-mobile-android build-desktop

# =============================================================================
build-mobile-android:
	gomobile init && \
	${GOBIN} get -d golang.org/x/mobile/cmd/gomobile && \
	CGO=1 gomobile bind -v -target=android -androidapi=${ANDROID_TARGET} -ldflags="-s -w" -o ./build/hello.aar ./mobile

# Copying output into react native androiod project
	mkdir -p ../mobile/AwesomeProject/android/app/libs/ && \
	rm -Rf ../mobile/AwesomeProject/android/app/libs/* && \
	cp build/*.aar ../mobile/AwesomeProject/android/app/libs/.

# =============================================================================
build-mobile-ios:
	gomobile init && \
	${GOBIN} get -d golang.org/x/mobile/cmd/gomobile && \
	CGO=1 gomobile bind -v -target=ios -ldflags="-s -w" -o ./build/hello.xcframework ./mobile

# =============================================================================
build-desktop-linux:
	$(GOBIN_SHARED_LIB_CFLAGS) $(GOBIN_SHARED_LIB_CGO_LDFLAGS) ${GOBIN} build \
		-buildmode=c-shared \
		-o ./build/libgohello.$(GOBIN_SHARED_LIB_EXT) ./desktop/

# Removing include automatically added by cgo
ifeq ($(detected_OS),Darwin)
	sed -i '' -e "s/#include \"..\/..\/nim\/libhello.h\"//gi" ./build/libgohello.h
else
	sed -i 's/#include "..\/..\/nim\/libhello.h"//gi' ./build/libgohello.h
endif

# Copying output into desktop project
	cp build/libgohello.$(GOBIN_SHARED_LIB_EXT) ../desktop/.
	cp build/libgohello.h ../desktop/.
