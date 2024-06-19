package main

import "fmt"

func main() {
	fmt.Println("welcome to class on struct")

	//no inheritance in golang, no super nd nno parent
	user1 := User{"Gaurang","gaurang@gamil.com",true,20}
	fmt.Println(user1)
	fmt.Printf("user1 details are %+v\n",user1)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
