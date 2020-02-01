package main

import (
	"fmt"
)

var x  = 0
func increment() {
	x = x + 5
}

func main() {
	for i := 0; i < 100; i++ {
		go increment()
	}
	fmt.Println("final x", x)
}