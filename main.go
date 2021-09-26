package main

//#cgo CFLAGS:-I${SRCDIR}/lib
//#cgo LDFLAGS:-L${SRCDIR}/lib -ltest_calc
//#include "test_calc.h"
import "C"
import "fmt"

func main() {
	sum := C.test_add(1,2)
	fmt.Printf("call cgo test_add function, sum is %d\n", sum)
}