package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

var counter = 0

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)

		// Locking for reading
		m.RLock()
		go printNum()

		// Locking for writing
		m.Lock()
		go increment()
	}

	wg.Wait()
}

func printNum() {
	fmt.Println("#", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
