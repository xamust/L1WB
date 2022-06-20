package main

import (
	"log"
	"time"
)

func main() {
	//для наглядности выдаем лог с указание времени старта
	log.Println("Start...")

	//в отдельной горутине запускаем собственный Sleep на 20 секунд
	//по выполнении которого будет выдано сообщение в лог с указанием времени завершения
	go func() {
		mySleep(10 * time.Second)
		log.Println("mySleep awake!")
	}()

	//в основной горутине запускаем time.Sleep на 20 секунд
	//по выполнении которого будет выдано сообщение в лог с указанием времени завершения
	time.Sleep(10 * time.Second)
	log.Println("goSleep awake!")
}

//собственная функция Sleep
func mySleep(i time.Duration) {
	//записываем время до старта цикла
	timeNow := time.Now()
	for {
		//в каждый такт записываем текущее время
		temp := time.Now()
		//проверяем, сколько времени прошло с момента старта цикла, если больше, либо равно, то выкидываем из цикла
		if temp.Sub(timeNow).Nanoseconds() >= i.Nanoseconds() {
			break
		}
	}
}
