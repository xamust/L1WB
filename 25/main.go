package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Start...")
	go func() {
		mySleep(20 * time.Second)
		log.Println("mySleep awake!")
	}()

	time.Sleep(20 * time.Second)
	log.Println("goSleep awake!")
}

func mySleep(i time.Duration) {
	//записываем время до старта цикла
	timeNow := time.Now()
	for {
		//в каждый такт записываем текущее время
		temp := time.Now()
		//проверяем, сколько времени прошло с момента старта цикла, если  то выкидываем из цикла
		if temp.Sub(timeNow).Nanoseconds() >= i.Nanoseconds() {
			break
		}
	}
}
