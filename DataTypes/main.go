package main

import "fmt"

func main() {
	// Section1 : Integers
	// Section1-A: singed integers
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	//assigning values to integer variables
	i = -128
	i8 = 127
	i16 = -32768
	i32 = 2147483647
	i64 = 1285042850

	fmt.Println("Signed integers", i, i8, i16, i32, i64)

	//Section1-B: unassigned integers
	//var u uint
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good morning"
	fmt.Println(msg, a)

}
