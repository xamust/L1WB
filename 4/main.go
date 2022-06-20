package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var workersNum int
	ch := make(chan int)
	out := make(chan int)
	fmt.Println("Set workers num:")
	fmt.Scan(&workersNum)
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go worker(i, ch, out)
	}
	for j := 1; j <= 5; j++ {
		ch <- j
	}
	close(ch)
	wg.Wait()
	for a := 1; a <= 5; a++ {
		<-out
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
	wg.Done()
}
