package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the reading")

	// comma ok || error ok syntax.
	input, _ := reader.ReadString('\n') // like try catch method.
	fmt.Println("Thanks for reading,",input)
	fmt.Printf("Type of reding is %T",input)
} 