package main

import "fmt"

// jwtToken := 300000 //it is not allowed
// var jwtToken int = 300000 //it is allowed

const LoginToken string = "skajdnkjadni" //capital L is denoting that it is public variable

func main() {
	var username string = "Gaurang"
	fmt.Println(username)
	fmt.Printf("variablesis type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variablesis type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //we can use up to 255 in uint8
	fmt.Println(smallVal)
	fmt.Printf("variablesis type: %T \n", smallVal)

	var smallFloat float64 = 255.455254555555555
	fmt.Println(smallFloat)
	fmt.Printf("variablesis type: %T \n", smallFloat)

	//default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variablesis type: %T \n", anothervariable)

	//implicit type
	var website = "Gaurang.BRDV" //auto assign type of variable
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variablesis type: %T \n", LoginToken)
}