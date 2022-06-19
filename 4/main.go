package main

import (
	"fmt"
	"os"
)

func main() {
	var workersNum int
	ch := make(chan string)
	fmt.Println("Set workers num:")
	fmt.Scan(&workersNum)
	for i := 0; i <= workersNum; i++ {
		go worker(i, ch)
	}
}

func worker(i int, input <-chan string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("Worker № %d send %v", i, input))
}
