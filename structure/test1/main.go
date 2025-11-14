package main

import "fmt"

type Student struct {
	Name  string
	Class int
	No    int
	Score float64
}

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var a int16 = 3456
	var b int8 = int8(a)
	fmt.Println(a, b)

	new_home := House{
		Address:  "Seoul",
		Size:     85,
		Price:    50000000,
		Category: "Apartment",
	}

	fmt.Printf("Address: %s, Size: %d, Price: %.2f, Category: %s\n",
		new_home.Address, new_home.Size, new_home.Price, new_home.Category)
}
