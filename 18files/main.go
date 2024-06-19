package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")

	content := "this needs to go in files - By gaurangbrdv@14"

	file, err := os.Create("./demo.txt")

	if err != nil {
		panic(err)
	}

	length,err := io.WriteString(file,content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:",length)
	defer file.Close()
	readFile("./demo.txt")
}

func readFile(filename string)  {
	databyte , err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("text data in file is \n",string(databyte))
}