package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")

	myChannel := make(chan int)

	// myChannel <- 5 //pushing the value in channel
	// fmt.Println(<-myChannel) //it will results in deadlock

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChannel)
		fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
