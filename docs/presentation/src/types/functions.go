package main

import (
	"fmt"
)

// START OMIT
type Food struct {
	Name string
}

type Animal struct {
	Name      string
	foodEaten int
}

func (a *Animal) Eat(food Food) {
	fmt.Printf("%s eats %s\n", a.Name, food.Name)
	a.foodEaten++
}

func main() {
	monkey := Animal{Name: "Monkey"}
	banana := Food{"Banana"}
	for i := 0; i < 5; i++ {
		monkey.Eat(banana)
	}
	fmt.Printf("%#v\n", monkey)
}

// END OMIT
