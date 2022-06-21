package main

import (
	"fmt"
	"math/rand"
	"time"
)

//структура для вывода корректной информации
type Result struct {
	result int
	flag   bool
}

func main() {
	//генерим массив
	testMassInt := createMass(21)

	//для бинарного поиска нужен отсортированный массив
	//sort(testMassInt)
	quickSort(testMassInt)
	//для наглядности выводим отсортированный массив
	fmt.Println(testMassInt)
	//выполняем поиск и выводим результат
	fmt.Println(binarySearch(testMassInt, 23))

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

//сортировка вставкой
func sort(sortedMass []int) {
	for i := 0; i < len(sortedMass)-1; i++ {
		for j := 0; j < len(sortedMass)-i-1; j++ {
			if sortedMass[j+1] < sortedMass[j] { //for revers sort >
				sortedMass[j+1], sortedMass[j] = sortedMass[j], sortedMass[j+1]
			}
		}
	}
}

//быстрая сортировка
func quickSort(sortedMass []int) {
	for i := 0; i < len(sortedMass)-1; i++ {
		if sortedMass[i] > sortedMass[i+1] {
			sortedMass[i], sortedMass[i+1] = sortedMass[i+1], sortedMass[i]
			quickSort(sortedMass)
		}
	}
}

func binarySearch(sortedMass []int, searchNum int) (result Result) {
	//в переменных min и max хранятся границы списка, в рамках которых проходит поиск
	min := 0
	max := len(sortedMass) - 1
	//циклом бегаем до тех пор, пока список не сократится до 1го элемента
	for max >= min {
		//делим список на 2 части
		count := (max + min) / 2
		//проверяем значение в середине, если равно искомому, то возвращаем структуру со значением и флагом true
		if sortedMass[count] == searchNum {
			return Result{
				result: count,
				flag:   true,
			}
			//если значение в середине списка больше искомого, то смещаем наш поиск на первую половину списка
		} else if sortedMass[count] > searchNum {
			max = count - 1
			//если значение в середине меньше искомого, то смещаем наш поиск на вторую половину списка
		} else {
			min = count + 1
		}
	}
	//если ничего не нашли, то возвращаем структуру c флагом false
	return Result{flag: false}
}
