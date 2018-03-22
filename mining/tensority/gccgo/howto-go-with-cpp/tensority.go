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
	// a := (*[32]C.uchar)(unsafe.Pointer(uintptr(int(reflect.ValueOf(blockHeader).Pointer()) - reflect.TypeOf(blockHeader).Align()*1)))
	a := (*[32]C.uchar)(unsafe.Pointer(&blockHeader[0]))
	b := (*[32]C.uchar)(unsafe.Pointer(&seed[0]))
	C.get(&a[0])
	C.get(&b[0])

	return hash
}

func main() {
	foo := New()
	foo.Bar()
	foo.Free()
}