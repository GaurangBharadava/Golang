package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to class on slices")

	var fruitList = []string{}
	fmt.Printf("the type of fruit list is %T\n",fruitList)

	fruitList = append(fruitList, "Apple","Mango")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0]=234
	highScores[1]=24
	highScores[2]=34
	highScores[3]=2
	// highScores[4]=2 it is not possible because we have allocate only 4 space to it

	highScores = append(highScores, 435, 556, 776) // it is possible, it will reallocate the value of memory.

	fmt.Println(highScores)

	sort.Ints(highScores) //it will sort integers to increasing order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //will return true or false value

	//remove the element from theslice
	var courses = []string{"react","hacking","js","swift","blockchain","node","express"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)
	fmt.Println(courses)
}
