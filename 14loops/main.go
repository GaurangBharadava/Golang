package main

import "fmt"

func main() {
	fmt.Println("welcome to class on loop in golang")

	days := []string {"sunday", "monday", "tuseday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	for d:=0 ; d<len(days) ; d++ {
		fmt.Println(days[d])
	}

	//another variation is for unknown range of traversel

	for i := range days {
		fmt.Println(days[i])
	}

	//for each loop

	for index,day := range days {
		fmt.Printf("%v:%v ",index,day)
	}

	//while loop

	roughValue := 1
	for roughValue<10 {
		fmt.Println("Value is: ",roughValue)
		roughValue++

		if roughValue==6 {
			break // we can use continue also.
		}
	}
}