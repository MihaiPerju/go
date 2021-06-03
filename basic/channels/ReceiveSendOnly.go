package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)

	wg.Add(2)

	// Channel that is going to receive data only
	go func(ch <-chan int) {

		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	// Channel that is going to send data only
	go func(ch chan<- int) {
		ch <- 43
		wg.Done()
	}(ch)

	wg.Wait()
}
