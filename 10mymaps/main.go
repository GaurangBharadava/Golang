package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on maps")

	languages := make(map[string]string) // map[key]value
	languages["js"] = "javascript"
	languages["sd"] = "solidity"
	languages["rt"] = "rust"

	fmt.Println("list of all programming languages is:",languages)
	fmt.Println("sd is hort for:",languages["sd"])

	//for delete the value from the map
	delete(languages,"rt")
	fmt.Println("list of all programming languages is:",languages)

	//loops are interesting in map

	for key,value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value)
	}
}