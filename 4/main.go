package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	var workersNum int
	in := make(chan int)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	fmt.Println("Set workers num:")
	fmt.Scan(&workersNum)
	log.Println("Workers pull start...")
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go worker(i, in, ctx)
	}

	for x := 0; ; x++ {
		select {
		case <-ctx.Done():
			close(in)
			wg.Wait()
			return
		default:
		}
		time.Sleep(time.Second)
		in <- x
	}
}

func worker(id int, in chan int, ctx context.Context) {
	defer wg.Done()
	for j := range in {
		select {
		case <-ctx.Done():
			log.Println("Pull is closed")
			return
		default:
			fmt.Printf("worker %d receive %d\n", id, j)
		}
	}
}
