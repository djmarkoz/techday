package main

import "fmt"

type MyStruct struct {
	Name string
	Age  int
}

func zero() {
	var i int
	var s string
	var m MyStruct
	var p *MyStruct

	fmt.Printf("i: %d\n", i)
	fmt.Printf("s: %#v\n", s)
	fmt.Printf("m: %#v\n", m)
	fmt.Printf("p: %#v\n", p)
}

func main() {
	zero()
}
