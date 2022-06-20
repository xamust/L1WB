package main

import (
	"sync"
)

var (
	//для конкурентной записи в мапу, необходимо воспользоваться мьютексами
	mu sync.Mutex

	//для наглядной демонстрации
	//wg sync.WaitGroup
)

//пустая структура для заполнения значений мапы (пустая структура = 0)
type Empty struct{}

func main() {
	//реализуем нашу мапу
	myMap := make(map[int]struct{})
	//циклом генерируем 100 рутин
	for i := 0; i < 100; i++ {
		//для наглядной демонстрации
		//wg.Add(1)

		//вызываем нашу функцию в каждой из сгенерированных горутин
		go write(i, myMap)
	}

	//для наглядной демонстрации
	//wg.Wait()
	//for mu.TryLock() {
	//fmt.Printf("len:%d\n %#v\n", len(myMap), myMap)
	//}

}

//функция записи в мапу, в качестве аргумента значение для записи и целевая мапа
func write(i int, input map[int]struct{}) {
	//блокируем горутиной, вызвавшей функцию
	mu.Lock()
	//разблокируем
	defer mu.Unlock()
	//записываем в мапу значение пустой структуры по ключу i, полученному в качестве аргумента
	input[i] = Empty{}

	//для наглядной демонстрации
	//wg.Done()
}