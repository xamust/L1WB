package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//генерируем два неупорядоченных множества
	mass1 := createMass(45)
	mass2 := createMass(45)
	fmt.Println(mass1)
	fmt.Println(mass2)
	//выводим результат
	fmt.Println(intersection(mass1, mass2))
}

//пустая структура для мапы
type Empty struct{}

//фунцкия реализации пересечения множеств
func intersection(mass ...[]int) map[int]struct{} {

	result := make(map[int]struct{})

	//пробегаемся по первому массиву
	for _, v := range mass[0] {
		//пробегаемся по 2му массиву
		for _, v2 := range mass[1] {
			//если значение из первого множества совпадает со значением из второго, то добавляем ввиде ключа в мапу
			if v == v2 {
				result[v] = Empty{}
				//выходим из цикла по второму слайсу
				break
			}
		}
	}
	//возвращаем результирующий слайс
	return result
}

//функция генерация массива со случайными элементами
func createMass(countElement int) []int {
	rand.Seed(time.Now().UnixNano())
	resultMass := make([]int, countElement)
	for i := 0; i < countElement; i++ {
		resultMass[i] = rand.Intn(100)
	}
	return resultMass
}
