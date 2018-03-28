package main

import "fmt"

func strings() {
	first := "Mark"
	var last string = "Freriks"
	name := first + " " + last
	fmt.Println("My name is", name)
}

func main() {
	strings()
}
