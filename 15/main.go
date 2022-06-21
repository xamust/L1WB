package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}
*/
//первый - это нереализована сама функция createHugeString(int) :)
//1. глобальная переменная
//2. избыточный результат работы функции

func main() {
	fmt.Println(createHugeString(1 << 7)[:100])
}

//реализация createHugeString
func createHugeString(lenStr int) string {
	var result []rune
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	for i := 0; i < lenStr-1; i++ {
		result = append(result, letterRunes[rand.Intn(len(letterRunes))])
	}
	return string(result)
}
