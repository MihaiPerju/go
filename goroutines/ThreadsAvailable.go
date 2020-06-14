package main

import (
	"fmt"
	"runtime"
)

func main() {
	// By default, Go will run the number of threads equal to the number of the CPU cores
	// available on the machine

	numOfThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Threads: ", numOfThreads)
}
