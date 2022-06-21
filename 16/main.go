package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//генерация массива
	testMassInt := createMass(21)
	fmt.Println(testMassInt)
	//сортируем
	quickSort(testMassInt)
	//для наглядности выводим отсортированный массив
	fmt.Println(testMassInt)
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

//функция генерация массива со случайными элементами
func createMass(countElement int) []int {
	rand.Seed(time.Now().UnixNano())
	resultMass := make([]int, countElement)
	for i := 0; i < countElement; i++ {
		resultMass[i] = rand.Intn(100)
	}
	return resultMass
}
