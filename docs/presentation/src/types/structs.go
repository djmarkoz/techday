package main

import "fmt"

// START OMIT

type Person struct {
	Name      string
	Age       int
	foodEaten int
}

func structs() {
	mark := Person{
		Name: "Mark",
		Age:  31,
	}
	fmt.Printf("%#v\n", mark)
}

// END OMIT

func main() {
	structs()
}
