package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://www.coursera.org/courses?query=free"

func main() {
	fmt.Println("Urls in golang")
	fmt.Println(URL)

	//parsing

	result, _ := url.Parse(URL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	// fmt.Println(result.Port())
	fmt.Println(result.Hostname())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)

	Qparams := result.Query()

	fmt.Printf("the type of quary params is %T\n",Qparams) // Key Value pair.
	fmt.Println(Qparams["query"])

	for _, val := range Qparams {
		fmt.Println(val) // wi;; print all the values in Qparams.
	}


	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl)
}