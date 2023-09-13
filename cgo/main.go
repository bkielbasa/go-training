package main

/*
#cgo CFLAGS: -I${SRCDIR}/ctestlib
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/ctestlib
#cgo LDFLAGS: -L${SRCDIR}/ctestlib
#cgo LDFLAGS: -ltest

#include <test.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Println("Get string")
	getString := C.GoString(C.get_string())
	fmt.Println(getString)
}
