package library

/*
#include "../libhello.h"
*/
import "C"

//export HelloFromGo
func HelloFromGo() {
	C.hello()
}
