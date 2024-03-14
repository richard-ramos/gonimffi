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
main: | build deps
	echo -e $(BUILD_MSG) "build/$@" && \
		$(ENV_SCRIPT) "$(NIMC)" c --verbosity:3 --parallelbuild:1 -f -o:./build/android/x86_64/libhello.so \
			--app:lib --passC:--sysroot="$(SYSROOT)" --passL:--sysroot="$(SYSROOT)" \
			--cincludes:"$(SYSROOT)/usr/include/x86_64-linux-android" \
			--cpu:amd64 --os:linux  -d:androidNDK --noMain ./hello.nim
clean: | clean-common
	rm -rf build/*

endif # "variables.mk" was not included


