package main

import (
	"fmt"
	"math"
)

func numbers() {
	// int int8 int16 int32 int64
	var i int = 55
	/*
	 * Usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
	 * When you need an integer value you should use int unless you have a specific reason
	 * to use a sized or unsigned integer type.
	 */
	var i16 int16 = 32767
	fmt.Println("int:", i, ", i16:", i16)

	// uint uint8 uint16 uint32 uint64
	var unsignedInt uint = 1234567890
	var unsignedInt8 uint8 = 255
	fmt.Println("unsignedInt:", unsignedInt, ", unsignedInt8:", unsignedInt8)

	// float32 float64
	var pi = 3.1415926536
	var pi32 float32 = math.Pi // 3.1415926536
	var pi64 float64 = math.Pi
	fmt.Printf("pi: %v (%T)\n", pi, pi)
	fmt.Printf("pi32: %v (%T)\n", pi32, pi32)
	fmt.Printf("pi64: %v (%T)\n", pi64, pi64)
}

func main() {
	numbers()
}
