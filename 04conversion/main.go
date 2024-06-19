package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pexxa between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for reading,", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // when we read the input and hit enter key then it will also read enter so that we use strings package
	if err != nil { //64 is for float 64 bit
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your reading:", numRating+1)

	}

}
