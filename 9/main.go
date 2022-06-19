package main

import (
	"fmt"
	"os"
)

var mass = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

func main() {
	num := make(chan int)
	incr := make(chan int)

	go func(mass []int) {
		for _, v := range mass {
			num <- v
		}
		close(num)
	}(mass)

	go func() {
		for {
			result, ok := <-num
			if !ok {
				break
			}
			incr <- result * 2
		}
		close(incr)
	}()

	for v := range incr {
		fmt.Fprintln(os.Stdout, v)
	}
}
