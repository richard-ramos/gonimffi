package library

/*
#include "../libhello.h"
#cgo amd64,android LDFLAGS:-L${SRCDIR}/../../nim/build/android/x86_64 -lhello
*/
import "C"

func HelloFromGo() {
	C.hello()
}
