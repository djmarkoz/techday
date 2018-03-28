package main

import (
	"fmt"
	"time"
)

// START OMIT

func producer(channel chan<- int, count int) {
	for i := 0; i < count; i++ {
		channel <- i
	}
	close(channel)
}

func consumer(channel <-chan int) {
	for i := range channel {
		time.Sleep(1 * time.Second)
		fmt.Println("done", i)
	}
}

func main() {
	channel := make(chan int)
	go producer(channel, 5)

	consumer(channel)
}

// END OMIT
