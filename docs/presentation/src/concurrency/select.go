package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "result"
	}()

	select {
	case res := <-channel:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
