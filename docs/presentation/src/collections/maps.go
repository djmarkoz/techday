package main

import "fmt"

func maps() {
	var favouriteFood map[string]string     // nil map (writes will panic)
	favouriteFood = make(map[string]string) // initialized map
	favouriteFood = map[string]string{}     // another way of initializing map
	favouriteFood = map[string]string{      // even with initial data
		"Tiger":  "Meat",
		"Monkey": "Banana",
		"Bird":   "Sunflower seed",
	}
	favouriteFood["Monkey"] = "Bananananana"
	delete(favouriteFood, "Bird")

	// order is not specified and is not guaranteed to be the same from one iteration to the next
	for key, value := range favouriteFood {
		fmt.Println(key, value)
	}

	fmt.Printf("tigersFood: '%s'\n", favouriteFood["Tiger"])
	fmt.Printf("zeroValue: '%s'\n", favouriteFood["Bird"])
}

func main() {
	maps()
}
