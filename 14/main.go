package main

import (
	"fmt"
	"math/big"
	"reflect"
)

type Empty struct{}

var (
	ch   = make(chan int)
	bInt = big.NewInt(100)
	//входной массив интерфейсов с различными типами
	massInterface = []interface{}{1, 0.87, "string", true, ch, Empty{}, bInt}
)

func main() {
	//цикл перебора и вывода типов переменных в массиве
	for _, v := range massInterface {
		//при помощи пакета reflect выведем типы переменных
		fmt.Println(reflect.TypeOf(v).String())
	}
}
