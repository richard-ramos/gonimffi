package library

/*
#include "../../nim/libhello.h"

#cgo amd64,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/x86_64 -lhello
#cgo arm64,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/arm64-v8a -lhello
#cgo arm,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/armeabi-v7a -lhello
#cgo 386,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/x86 -lhello

#cgo darwin,arm64,ios LDFLAGS:-L${SRCDIR}/../../nim/build/ios/iphone/hello
#cgo darwin,amd64,ios LDFLAGS:-L${SRCDIR}/../../nim/build/ios/simulator/hello
*/
import "C"

func HelloFromGo() {
	C.hello()
}
