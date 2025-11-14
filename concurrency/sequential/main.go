package main

import (
	"fmt"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(order Order) {
	//Simulate order preparation time
	fmt.Printf("Preparing order for table %d\n", order.TableNumber)

	time.Sleep(order.PrepTime)

	fmt.Printf("Order for table %d is ready!\n", order.TableNumber)
}

func main() {
	orders := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 1 * time.Second},
		{TableNumber: 3, PrepTime: 3 * time.Second},
		{TableNumber: 4, PrepTime: 2 * time.Second},
		{TableNumber: 5, PrepTime: 4 * time.Second},
	}

	for _, order := range orders {
		// processOrder(order)
		go processOrder(order) // 아무것도 출력되지 않는다 -> goroutine,
	}

	// how can we synchronize here?
	// Sync problem ?
	//time.Sleep(10 * time.Second)
	fmt.Scanln()
}
