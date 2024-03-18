SHELL := bash # the shell used internally by "make"

# used inside the included makefiles
BUILD_SYSTEM_DIR := vendor/nimbus-build-system

# we don't want an error here, so we can handle things later, in the ".DEFAULT" target
-include $(BUILD_SYSTEM_DIR)/makefiles/variables.mk

.PHONY: \
	all \
	deps \
	update \
	lib \
	main \
	build-mobile \
	build-desktop \
	copy \
	clean

ifeq ($(NIM_PARAMS),)
# "variables.mk" was not included, so we update the submodules.
GIT_SUBMODULE_UPDATE := git submodule update --init --recursive
.DEFAULT:
	+@ echo -e "Git submodules not found. Running '$(GIT_SUBMODULE_UPDATE)'.\n"; \
		$(GIT_SUBMODULE_UPDATE) && \
		echo
# Now that the included *.mk files appeared, and are newer than this file, Make will restart itself:
# https://www.gnu.org/software/make/manual/make.html#Remaking-Makefiles
#
# After restarting, it will execute its original goal, so we don't have to start a child Make here
# with "$(MAKE) $(MAKECMDGOALS)". Isn't hidden control flow great?

else # "variables.mk" was included. Business as usual until the end of this file.

# default target, because it's the first one that doesn't start with '.'
all: | main

# must be included after the default target
-include $(BUILD_SYSTEM_DIR)/makefiles/targets.mk

# add a default Nim compiler argument
NIM_PARAMS += -d:release

deps: | deps-common
	# Have custom deps? Add them above.


update: | update-common
	# Do you need to do something extra for this target?

SYSROOT = /home/richard/Android/Sdk/ndk/21.4.7075529/toolchains/llvm/prebuilt/linux-x86_64/sysroot/

# building Nim programs
build-mobile-android: | deps
	echo -e $(BUILD_MSG) "build/$@" && \
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/android/x86_64/libhello.so \
			--app:lib --passC:--sysroot="$(SYSROOT)" --passL:--sysroot="$(SYSROOT)" \
			--passL:-llog \
			--cincludes:"$(SYSROOT)/usr/include/x86_64-linux-android" \
			--cpu:amd64 --os:android  -d:androidNDK --noMain ./hello.nim && \
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/android/x86/libhello.so \
			--app:lib --passC:--sysroot="$(SYSROOT)" --passL:--sysroot="$(SYSROOT)" \
			--cincludes:"$(SYSROOT)/usr/include/i686-linux-android" \
			--passL:-llog \
			--cpu:i386 --os:android  -d:androidNDK --noMain ./hello.nim && \
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/android/arm64-v8a/libhello.so \
			--app:lib --passC:--sysroot="$(SYSROOT)" --passL:--sysroot="$(SYSROOT)" \
			--cincludes:"$(SYSROOT)/usr/include/aarch64-linux-android" \
			--passL:-llog \
			--cpu:arm64 --os:android  -d:androidNDK --noMain ./hello.nim && \
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/android/armeabi-v7a/libhello.so \
			--app:lib --passC:--sysroot="$(SYSROOT)" --passL:--sysroot="$(SYSROOT)" \
			--cincludes:"$(SYSROOT)/usr/include/armv7a-linux-androideabi" \
			--passL:-llog \
			--cpu:arm --os:android  -d:androidNDK --noMain ./hello.nim

IPHONE_MIN_VERSION=16.1
IPHONE_SIMULATOR_SYSROOT="/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator$(IPHONE_MIN_VERSION).sdk"
build-mobile-ios: | deps
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/ios/simulator/hello \
			--app:lib \
			--passC:-isysroot="$(SYSROOT)" --passL:-isysroot="$(SYSROOT)" \
			--passC:-mios-simulator-version-min=$(IPHONE_MIN_VERSION) --passL:-mios-simulator-version-min=$(IPHONE_MIN_VERSION) \
			--cpu:amd64 --os:macos -d:ios -d:simulator -d:iPhone --noMain ./hello.nim

build-desktop: | deps
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/linux/x86_64/libhello.so \
			--app:lib --cpu:amd64 --os:linux --noMain ./hello.nim

copy: build-mobile-android build-desktop
		mkdir -p ../mobile/AwesomeProject/android/app/src/main/jniLibs/ && \
		rm -Rf ../mobile/AwesomeProject/android/app/src/main/jniLibs/* && \
		cp -R ./build/android/* ../mobile/AwesomeProject/android/app/src/main/jniLibs/. && \
		cp ./build/linux/x86_64/*.so ../desktop/.

main: | copy

clean: | clean-common
	rm -rf build/*

endif # "variables.mk" was not included


