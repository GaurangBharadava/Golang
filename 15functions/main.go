package main

import "fmt"

func main() {
	fmt.Println("functions in golang")
	greeter()

	result := adder(3,4)
	fmt.Println("the result is:", result)

	proResult := proAdder(2,3,4,5,6,7,8,10)
	fmt.Println(proResult)
}

func adder(valOne int, valTwo int) int {
	return valOne+valTwo
}

func greeter() {
	fmt.Println("namastey")
}

//function for no fix argument

func proAdder(values ... int) int {
	total := 0

	for _,value := range values{
		total += value
	}
	return total
}