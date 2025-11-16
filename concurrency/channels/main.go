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

func sendMessage(ch chan<- string, num int, wg *sync.WaitGroup) {
	fmt.Printf("Sending message %d\n", num)
	time.Sleep(time.Second * time.Duration(num)) // simulate variable processing time
	msg := fmt.Sprintf("Message %d sent\n", num)
	ch <- msg // Send message to channel
	wg.Done()
}

func receiveMessage(msgs <-chan string) { // Receive messages from channel
	fmt.Println("Waiting for message")

	for msg := range msgs {
		fmt.Println("Received:", msg)
	}
}

func main() {

	msgs := make(chan string) // Create a channel to send messages

	wg := sync.WaitGroup{} // WaitGroup to wait for all senders to finish
	wg.Add(2)
	go sendMessage(msgs, 1, &wg)
	go sendMessage(msgs, 2, &wg)

	go func() {
		receiveMessage(msgs)
	}()

	wg.Wait()
	close(msgs) // Close the channel when done

	fmt.Println("Work done")
	fmt.Scanln()
}
