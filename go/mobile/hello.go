package library

/*
#include "../libhello.h"
#cgo amd64,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/x86_64 -lhello
#cgo arm64,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/arm64-v8a -lhello
#cgo arm,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/armeabi-v7a -lhello
#cgo 386,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/x86 -lhello
*/
import "C"

func HelloFromGo() {
	C.hello()
}
