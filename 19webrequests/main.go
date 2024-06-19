package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const URL = "https://google.com"

func main()  {
	fmt.Println("Web Requests")

	res , err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Responce is of type: %T\n", res)

	defer res.Body.Close() //caller's reaponsibility to close this connection.

	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)
}