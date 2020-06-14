package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 50)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		for i := 0; i < 10; i++ {
			ch <- i
		}

		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
