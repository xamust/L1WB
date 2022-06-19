package main

import (
	"fmt"
	"os"
)

var (
	mass = []int{2, 4, 6, 8, 10}
)

func main() {
	in := make(chan int)
	for _, v := range mass {
		go func(v int) {
			in <- v * v
		}(v)
	}
	for range mass {
		fmt.Fprintln(os.Stdout, <-in)
	}
}
