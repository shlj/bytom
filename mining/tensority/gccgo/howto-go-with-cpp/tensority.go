package main

// #cgo LDFLAGS: -L . -lstdc++
// #cgo CFLAGS: -I ./
// #include "foo.h"
import "C"

import(
	// "fmt"
	"unsafe"
	// "reflect"
)

type GoFoo struct {
	foo C.Foo
}

func New() GoFoo {
	var ret GoFoo
	ret.foo = C.FooInit()
	return ret
}

func (f GoFoo) Free() {
	C.FooFree(f.foo)
}

func (f GoFoo) Bar() {
	C.FooBar(f.foo)
}

func testHelper(blockHeader [32]uint8, seed [32]uint8, hash [32]uint8) [32]uint8 {
	bhPtr := (*C.uchar)(unsafe.Pointer(&blockHeader))
	seedPtr := (*C.uchar)(unsafe.Pointer(&seed))
	hashPtr := (*C.uchar)(unsafe.Pointer(&hash))

	resPtr := C.get(bhPtr, seedPtr, hashPtr)
	res := *(*[32]uint8)(unsafe.Pointer(resPtr))
	return res
}

func main() {
	foo := New()
	foo.Bar()
	foo.Free()
}