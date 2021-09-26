package gowapper

//#cgo CFLAGS:-I${SRCDIR}/lib
//#cgo LDFLAGS:-L${SRCDIR}/lib -ltest_calc
//#include "lib/test_calc.h"
import "C"
import "fmt"

func CallStaticLib() {
	sum := C.TestAdd(1,2)
	fmt.Printf("call cgo TestAdd function, sum is %d\n", sum)
}