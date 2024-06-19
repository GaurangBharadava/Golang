package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	fmt.Println("maths in golang")

	var noOne int = 2
	var noTwo float64 = 4.5

	fmt.Println("the sum of the numbers are:",noOne + int(noTwo))

	//random from crypto
	number, _ := rand.Int(rand.Reader,big.NewInt(5))
	fmt.Println("the random number is:",number)

}