# Building a nim library to be consumed from a go library from mobile or desktop environments

## Requirements
- go
- gomobile
- nodejs
- yarn
- AndroidSDK
- AndroidNDK 
- Build essentials (gcc/clang)

## Clone the repo
```
git clone https://github.com/richard-ramos/gonimffi.git
cd gonimffi
git submodule init
```

## Android

### Nim
1. Edit in `nim/nim.cfg` the `clang.path` variables to point to the folder in which the AndroidNDK clang resides
2. Edit in `nim/Makefile` the variable `ANDROID_SYSROOT` to point to the `sysroot` folder of AndroidNDK
3. Execute 
```
cd nim
make build-mobile-android
```

This will generate `libhello.so` libraries for the four architectures supported in Android. It will proceed to copy these
files in `mobile/AwesomeProject/android/app/src/main/jniLibs`

To add more functions, edit `nim/hello.nim`, and `nim/libhello.h`. The function `NimMain()` must be called in an initialization
function somewhere in order to initialize Nim's garbage collector.

You might need to link against Android libraries depending on the functions you use. i.e., if you use `echo`, then you need to
link against `liblog`, which is done with `--passL:-llog` (already done on the `Makefile`)

### Go
1. Execute 
```
cd go
make build-mobile-android
```

This will use gomobile to generate the file `hello.aar`. This will be copied in `mobile/AwesomeProject/android/app/libs/`.

To add any additional function, edit `go/mobile/hello.go`. `gomobile` will export functions that start with an Uppercase. Not all datatypes are supported: `bool`, `string`, `int`, `uint`, `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `uintptr`, `byte`, `rune`, `float32`, `float64`, `complex64`, `complex128`.

### React Native
1. Install dependencies
```
cd mobile/AwesomeProject
yarn
```

To load the nim and go libraries, `mobile/AwesomeProject/android/app/build.gradle` was edited with:
```
android {
    ...
    sourceSets {
        main {
            jniLibs {
                srcDirs 'src/main/jniLibs'
            }
        }
    }
    ...
}

dependencies {
    ...
    implementation fileTree(dir: "libs", include: ["*.aar"])
    ... 
}
```

The following files were created:
- mobile/AwesomeProject/android/app/src/main/java/com/awesomeproject/ReactNativeModule.kt
- mobile/AwesomeProject/android/app/src/main/java/com/awesomeproject/ReactNativePackage.kt

Additional functions must be added to `ReactNativeModule.kt`.

In `mobile/AwesomeProject/android/app/src/main/java/com/awesomeproject/MainActivity.kt` `libhello.so` (the nim library) is loaded.
This is necessary so the go library (`hello.aar`) is able to work (otherwise you'll get an `Undefined Reference` error)
```kotlin
    companion object {
        init {
            System.loadLibrary("hello")
        }
    }
```
Finally, an invocation to the function we exposed in Go, can be found in `mobile/AwesomeProject/App.tsx`. This should
print a `hello world` in the android logs.

### Testing

Using the emulator, assuming you have a Pixel6 API 33 image (this assumes the `platform-tools` folder from the AndroidNDK is available on the `PATH`)
```
./emulator -avd Pixel_6_API_33
```
Docs on how to create an image can be found in https://developer.android.com/studio/run/managing-avds

Once the emulator is running, in a different terminal:
```
cd mobile/AwesomeProject
npm run start
```
And press `a` key once a menu appears to run the app on Android.

If using a mobile device, refer to this documentation: https://reactnative.dev/docs/running-on-device

To see the `hello world` message, you need to use `adb logcat` (available on the AndroidNDK `platform-tools` folder)


## IOS
TODO

## Desktop (Linux)
This assumes you're building in a AMD64 machine with linux
### Nim
```
cd nim
make build-desktop-linux
```
This will generate a shared library `libhello.so` and copy it into the `desktop` folder.

###  Go
```
cd go
make build-desktop-linux
```
This will generate a shared library `libgohello.so` and copy it into the `desktop` folder.

### Desktop app
For the purposes of this example, we use a simple C program:
```
cd desktop
make
make run
```

## Desktop (MacOSX)
TODO

## Desktop (Windows)
TODO
