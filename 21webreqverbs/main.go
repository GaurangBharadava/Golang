package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web server in golang")
	// PerformRequestGet()
	// PerformPostJsonRequest()
	PerformPostFormRequest()

}

func PerformRequestGet() {
	const myUrl = "http://localhost:8000/get"

	responce, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer responce.Body.Close()

	fmt.Println("Status code:", responce.StatusCode)
	fmt.Println("Content length is", responce.ContentLength)

	var responceString strings.Builder
	content, _ := io.ReadAll(responce.Body)

	byteCount, _ := responceString.Write(content)

	fmt.Println("bytecount is:", byteCount)
	fmt.Println(responceString.String())
	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myUrl  = "http://localhost:8000/post"

	//fake json payload

	reqBody := strings.NewReader(`
		{
			"course" : "lets learn go",
			"prise" : 0,
			"plateform" : "youtube"
		}
	`)

	res, err := http.Post(myUrl,"application/json",reqBody)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl  = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname","Gaurang")
	data.Add("lastname","BRDV")
	data.Add("email","gaurang@yahoo.com")

	res, err := http.PostForm(myUrl,data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))
}