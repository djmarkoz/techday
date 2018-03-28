package main

import "fmt"

// START OMIT
func say(s string) {
	fmt.Println(s)
}

func main() {
	say("Hello?!")
	defer say("Hi there!")

	say("Is there any one here?!")
}
