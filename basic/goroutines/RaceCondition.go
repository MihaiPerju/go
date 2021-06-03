package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	msg := "Hello"

	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "Goodbye"

	wg.Wait()
}
