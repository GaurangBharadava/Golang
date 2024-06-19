package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string {"test"}

var wg sync.WaitGroup //usually this groups are poiners

var mutex sync.Mutex //pointer

func main() {
	fmt.Println("go routines")
	//go routine is simply created by adding a keyword go.
	// go greeter("namaste")
	// greeter("hwllo")

	websiteList := []string{
		"https://github.com",
		"https://go.dev",
		"https://google.com",
	}

	for _ , web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait() // always go to the end of the main method
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(2* time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {

	defer wg.Done()
	result ,err := http.Get(endpoint)

	if err!=nil {
		fmt.Println("oops in endpoint")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d status code for %s\n",result.StatusCode,endpoint)
	}

	
}