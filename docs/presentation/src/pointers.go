package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func pointers() (Person, *Person, Person, *Person) {
	value := Person{Name: "Mark", Age: 30}

	pointer := &value
	pointer.Age = 31

	copy := *pointer
	copy.Name = "Copied " + copy.Name

	return value, pointer, copy, nil
}

func main() {
	value, pointer, copy, nilPerson := pointers()

	fmt.Printf("value: %[1]v (%[1]T)\n", value)
	fmt.Printf("pointer: %[1]v %[1]p (%[1]T)\n", pointer)
	fmt.Printf("copy: %[1]v (%[1]T)\n", copy)
	fmt.Printf("nilPerson: %[1]v (%[1]T)\n", nilPerson)
}
