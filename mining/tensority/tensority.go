package tensority

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L./lib/ -l:cSimdTs.o -lstdc++
// #include "./lib/cSimdTs.h"
import "C"

import(
    "unsafe"
)

func Hash(blockHeader [32]uint8, seed [32]uint8) [32]uint8 {
    bhPtr := (*C.uchar)(unsafe.Pointer(&blockHeader))
    seedPtr := (*C.uchar)(unsafe.Pointer(&seed))

    resPtr := C.SimdTs(bhPtr, seedPtr)
    res := *(*[32]uint8)(unsafe.Pointer(resPtr))
    return res
}