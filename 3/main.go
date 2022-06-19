package main

import (
	"fmt"
	"os"
)

var (
	mass   = []int{2, 4, 6, 8, 10}
	result = 0
)

func main() {
	out := make(chan int)
	for _, v := range mass {
		go func(v int) {
			out <- v * v
		}(v)
	}
	for range mass {
		result += <-out
	}
	fmt.Fprintln(os.Stdout, result)
}
