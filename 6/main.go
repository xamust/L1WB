package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

//1.завершение горутины, путем закрытия канала
func CloseChannel() chan string {
	log.Println("Close chan, start")
	ch := make(chan string)
	go func() {

		for {
			select {
			case ch <- "1 stil works":

			case <-ch:
				log.Println("Goroutine1 stop, ctx")
				return
			}
		}
	}()
	wg.Done()
	return ch
}

//2.завершение горутины, при помощи контекста
func CloseCtx() {
	log.Println("Context, start")
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("Goroutine2 stop, ctx")
				return
			default:
				//some work...
				time.Sleep(time.Second)
				log.Println("2 stil works")
			}
		}
	}(ctx)
	time.Sleep(time.Second)
	cancel()
	wg.Done()
}

func main() {
	wg.Add(1)
	myNumber := CloseChannel()
	fmt.Println(<-myNumber)
	fmt.Println(<-myNumber)
	close(myNumber)

	wg.Add(1)
	CloseCtx()

	wg.Wait()
	log.Println("end")
}
