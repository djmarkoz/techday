package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	number, err := GetNumber()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number was: %d", number)
}

func GetNumber() (int, error) {
	return 0, errors.New("failed to get number") // HLerr
}
