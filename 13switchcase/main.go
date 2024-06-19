package main

import (
	"math/rand"
	"fmt"
	
)

func main() {
	fmt.Println("switch case in golang")

	//dice game
	diceNumber := rand.Intn(6) + 1
	fmt.Println("the value of dice is:",diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the dice value is one")
	case 2:
		fmt.Println("the dice value is two")
	case 3:
		fmt.Println("the dice value is three")
		// fallthrough // the falltrough is use to print next case also.
	case 4:
		fmt.Println("the dice value is four")
	case 5:
		fmt.Println("the dice value is five")
	case 6:
		fmt.Println("the dice value is six")
	default:
		fmt.Println("invalid Error")
	}

}