package main

import "fmt"

func arrays() {
	array := [2]string{"Mark", "Freriks"}
	fmt.Println(array)

	var a [4]int
	a[1] = 1337
	i := a[1]
	fmt.Println("i:", i)
}

func main() {
	arrays()
}
