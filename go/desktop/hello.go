package main

/*
#include "../../nim/libhello.h"
#cgo amd64,linux,!android LDFLAGS:-L${SRCDIR}/../../nim/build/linux/x86_64 -lhello
*/
import "C"

//export HelloFromGo
func HelloFromGo() {
	C.hello()
}

func main() {
	// Do nothing
}
