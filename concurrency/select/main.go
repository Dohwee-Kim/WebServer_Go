package main

import (
	"context"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for {
		select {
		case msg := <-strings:
			fmt.Println(msg)
		case num := <-nums:
			fmt.Println(num)
		case <-ctx.Done():
			fmt.Println("Context Cancelled:", ctx.Err())
			return
		}
	}
}
