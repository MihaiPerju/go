package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#74bf74",
	}

	// var colors map[string]string

	// colors := make(map[string]string)
	colors["white"] = "ff0000"

	for color, hex := range colors {
		fmt.Println(color, hex)
	}

	fmt.Println(colors)
}
