package main

import (
	"fmt"
	"time"
)

func main() {

	strings := make(chan string)
	nums := make(chan int)

	go func() {
		for {
			strings <- "hello"
			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		for {
			nums <- 15
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		select {
		case msg := <-strings:
			fmt.Println(msg)
		case num := <-nums:
			fmt.Println(num)

		}
	}
}
