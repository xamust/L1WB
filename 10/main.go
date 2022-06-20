package main

import "fmt"

//входная строка
var tempMes = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {
	//инициализируем мапу для хранения (ключ - значение int, значение - слайс float64, подмножество)
	mapRes := make(map[int][]float64)
	//пробегаемся по входному массиву
	for _, v := range tempMes {
		//данной конструкцией получаем группы десятков для группировки подмножеств
		temp := (int(v) / 10) * 10
		//если в нашей мапе нет такой строки, то добавляем ее
		if _, ok := mapRes[temp]; !ok {
			//добавляем новое подмножество в мапу
			mapRes[temp] = append(mapRes[temp], v)
		} else {
			//если в нашей мапе есть такое подмножество, то добавляем значение в существующее подмножество
			mapRes[temp] = append(mapRes[temp], v)
		}
	}
	//выводим мапу с множествами
	fmt.Printf("%#v\n", mapRes)
}