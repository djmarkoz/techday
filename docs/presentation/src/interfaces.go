package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}

// START DEF OMIT

type Animal struct {
	Name  string
	Sound string
}

type Human struct {
	Name string
	Age  int
}

// END DEF OMIT

// START IMPL OMIT

func (a Animal) Greet() {
	fmt.Printf("%s, said the %s.\n", a.Sound, a.Name)
}

func (h Human) Greet() {
	fmt.Printf("%s (%d) says hi.\n", h.Name, h.Age)
}

// END IMPL OMIT

func main() {
	greeters := []Greeter{
		Animal{
			Name:  "Dog",
			Sound: "Whoof",
		},
		Human{
			Name: "Mark",
			Age:  31,
		},
	}

	for _, g := range greeters {
		g.Greet()
	}
}

// START EMPTY OMIT

func printAnything(a ...interface{}) {
	fmt.Println(a)
}

func empty() {
	mark := Human{
		Name: "Mark",
		Age:  31,
	}
	printAnything(mark, 1)
}

// END EMPTY OMIT
