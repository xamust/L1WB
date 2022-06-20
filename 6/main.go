package main

import "fmt"

func CloseChannel() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func BoolToStop(bool) {
	ch := make(chan bool)
	go func(chan bool) {
		for {
			select {
			case <-ch:
				return
			default:
				fmt.Println("Goroutine 2 is stil work...")
			}
		}
		fmt.Println("Goroutine 2 is finished")
	}(ch)
}

func main() {
	//?????
	//boolCh := make(chan bool)
	//BoolToStop(true)
	//boolCh <- true

	myNumber := CloseChannel()
	fmt.Println(<-myNumber)
	fmt.Println(<-myNumber)
	close(myNumber)
}
