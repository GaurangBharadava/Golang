package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println("the value of ptr id",ptr) the initial value of a pointer is nil.

	myNumber := 14
	var ptr = &myNumber

	fmt.Println("the value of pointer is",ptr)
	fmt.Println("the value of pointer is",*ptr)

	*ptr = *ptr + 2
	fmt.Println("the new value of pointer is",myNumber)

}