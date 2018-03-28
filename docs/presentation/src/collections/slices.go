package main

import "fmt"

func slices() {
	numbers := []int{1, 3, 3, 7}
	numbers[0] = 2
	numbers = append(numbers, 4) // HL
	fmt.Println("numbers:", numbers, "len:", len(numbers), "cap: ", cap(numbers))

	for _, v := range numbers {
		fmt.Println(v)
	}

	zero := make([]int, 4, 13)
	fmt.Println("zero:", zero, "len:", len(zero), "cap: ", cap(zero))
}

func main() {
	slices()
}
