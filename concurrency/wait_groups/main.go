package main

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func processOrder(waiterID int, order Order) {
	//Simulate order preparation time
	fmt.Printf("Waiter %d - Preparing order for table %d\n", waiterID, order.TableNumber)

	time.Sleep(order.PrepTime)

	fmt.Printf("Waiter %d - Order for table %d is ready!\n", waiterID, order.TableNumber)
}

func main() {
	orders := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 1 * time.Second},
		{TableNumber: 3, PrepTime: 3 * time.Second},
		{TableNumber: 4, PrepTime: 2 * time.Second},
		{TableNumber: 5, PrepTime: 4 * time.Second},
	}

	wg := sync.WaitGroup{}
	//wg.Add(len(orders)) // Set the number of goroutines to wait for 이렇게 하는 것보다

	for waiterID, order := range orders {
		wg.Add(1) // Increment the WaitGroup counter for each goroutine
		go func(wID int, ord Order) {
			defer wg.Done() // Decrement the counter when the goroutine completes
			processOrder(wID, ord)
		}(waiterID, order)
	}

	wg.Wait() // Wait for all goroutines to finish

}
