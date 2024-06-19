package main

import "fmt"

func main() {
	fmt.Println("welcome to class on if else condition statement")

	loginCount := 14
	var result string 

	if loginCount<10 {
		result = "reguler user"
	} else {
		result = "hacker meetup"
	}
	fmt.Println(result)

	//for handeling web requests this syntax is used
	if num := 3; num<10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}