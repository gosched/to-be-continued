package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func showPi() {
	// fmt.Println(math.Pi)

	// var f float64 = math.Pi
	// b1 := make([]byte, 8)
	// binary.LittleEndian.PutUint64(b1, math.Float64bits(f))
	// fmt.Printf("%#v\n", b1)
	// fmt.Printf("%v\n", b1)

	// bits := binary.LittleEndian.Uint64(b1)
	// fmt.Println(math.Float64frombits(bits))

	var err error

	var pi float64 = math.Pi
	var piPi float64

	buffer := new(bytes.Buffer)
	err = binary.Write(buffer, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buffer.Bytes()) // 18 2d 44 54 fb 21 09 40

	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.LittleEndian, &piPi)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(piPi)
}

func main() {
	showPi()
	// bytes1 := []byte{byte(0), byte(16), byte(64), byte(255)}
	// bytes2 := []byte{0x0, 0x10, 0x40, 0xff}
	// fmt.Printf("%v %#v\n", bytes1, bytes1)
	// fmt.Printf("%v %#v\n", bytes2, bytes2)
}
