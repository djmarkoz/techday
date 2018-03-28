package main

import (
	"fmt"
	"time"
)

// START OMIT

func slowFunction(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("done", i)
}

func main() {
	for i := 0; i < 5; i++ {
		go slowFunction(i) // HL
	}
}

// END OMIT
