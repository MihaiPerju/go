package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	ch := make(chan int)

	wg.Add(1)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	// Spinning 5 goroutines
	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func() {
			ch <- i
			wg.Done()
		}()
	}

	wg.Wait()
}
