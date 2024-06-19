package main

import "fmt"

func main() {
	fmt.Println("defer in golang")

	//it will print in reverse order of defer
	defer fmt.Println("world")
	defer myDefer()
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
