package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	//создаем канал
	ch := make(chan string)

	// создаём context с таймаутом (функция завершается по истечении таймаута)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//освобождаем ресурсы по завершению
	defer cancel()

	//логгер для наглядности времени выполнения
	log.Println("Start")

	//стартуем горутину
	go func(ctx context.Context, chData chan string) {
		i := 1
		for {
			select {
			//обработчик выполнения условий (таймаута)
			case <-ctx.Done():
				close(chData)
				return
			default:
			}
			//выводим данные в канал
			chData <- fmt.Sprintf("Data in channel: %d", i)
			i++
			//для наглядности отображения
			time.Sleep(1 * time.Second)
		}
	}(ctx, ch)

	for str := range ch {
		fmt.Println(str)
	}

	log.Println("Exit by timeout.")

	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	//второй вариант с передачей флага таймаута каналом из другой горутины
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func(chan1 chan string) {
		time.Sleep(10 * time.Second)
		chan1 <- "timer"
	}(chan1)

	//логгер для наглядности времени выполнения
	log.Println("Start2")
	//стартуем горутину
	go func(chan1, chan2 chan string) {
		i := 1
		for {

			select {
			//обработчик выполнения условий (таймаута)
			case <-chan1:
				close(chan2)
				return
			default:
			}
			//выводим данные в канал
			chan2 <- fmt.Sprintf("Data in channel: %d", i)
			i++
			//для наглядности отображения
			time.Sleep(1 * time.Second)
		}
	}(chan1, chan2)
	for str := range chan2 {
		fmt.Println(str)
	}

	log.Println("Exit2 by timeout.")
}
