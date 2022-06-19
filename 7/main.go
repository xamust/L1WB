package main

import (
	"sync"
)

var (
	mu sync.Mutex
	//wg sync.WaitGroup
)

type Empty struct{}

func main() {
	myMap := make(map[int]struct{})
	for i := 0; i < 100; i++ {
		//wg.Add(1)
		go write(i, myMap)
	}
	//wg.Wait()
	//for mu.TryLock() {
	//fmt.Printf("len:%d\n %#v\n", len(myMap), myMap)
	//}

}

func write(i int, input map[int]struct{}) {
	mu.Lock()
	input[i] = Empty{}
	mu.Unlock()
	//wg.Done()
}
