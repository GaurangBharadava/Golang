package main

import "fmt"

func main() {
	fmt.Println("methods in golang")

	user1 := User{"Gaurang","gaurang@gamil.com",true,20}
	fmt.Println(user1)
	fmt.Printf("user1 details are %+v\n",user1)

	user1.GetStatus()
	user1.newEmail()
}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

//Methods for struct

func (u User) GetStatus() {
	fmt.Println("is user active?", u.Status)
}

func (u User) newEmail() {
	//it won't change the actual value.
	u.Email = "ksdbabcvja@sadj.com"
	fmt.Println("the new Email is:",u.Email)
}