package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go")

	var fruitList [4] string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Watermelon"
	fmt.Println("the fruit list is",fruitList)
	fmt.Println("the fruit list is",len(fruitList))

	var vegList = [3]string{"potato","carrot","chilli"}
	fmt.Println("vegy list is:",vegList)

}