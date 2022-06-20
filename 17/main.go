package main

import "fmt"

//входные данные
var testMassInt = []int{100, 23, 67, 44, 78, 45, 34, 23, 31, 1, 0, 8, 4, 34, 35, 53, 46, 47, 90}

func main() {
	//для бинарного поиска нужен отсортированный массив
	sort(testMassInt)
	//для наглядности выводим отсортированный массив
	fmt.Println(testMassInt)
	//выполняем поиск и выводим результат
	fmt.Println(binarySearch(testMassInt, 23))
}

//сортируем массив сортировкой вставкой
func sort(sortedMass []int) {
	for i := 0; i < len(sortedMass)-1; i++ {
		for j := 0; j < len(sortedMass)-i-1; j++ {
			if sortedMass[j+1] < sortedMass[j] { //for revers sort >
				sortedMass[j+1], sortedMass[j] = sortedMass[j], sortedMass[j+1]
			}
		}
	}
}

func binarySearch(sortedMass []int, searchNum int) (result int) {
	result = -1
	min := 0
	max := len(sortedMass) - 1
	for max >= min {
		count := (max + min) / 2
		if sortedMass[count] == searchNum {
			i := 0
			for {
				if count > 0 && sortedMass[count-i] != sortedMass[count] {
					result = count
					break
				} else {
					i++
				}
			}
			break
		} else if sortedMass[count] > searchNum {
			max = count - 1
		} else {
			min = count + 1
		}
	}
	return
}
