package main

import (
	"fmt"
	"unsafe"
)

// https://en.wikipedia.org/wiki/Endianness

// INTSIZE .
const INTSIZE int = int(unsafe.Sizeof(0))

// SystemEndian .
func SystemEndian() {
	var i int = 0x1
	byteSlice := (*[INTSIZE]byte)(unsafe.Pointer(&i))
	if byteSlice[0] == 0 {
		fmt.Println("Little endian")
	} else {
		fmt.Println("Big endian")
	}
	fmt.Println()
}

func main() {
	SystemEndian()
}
