package main

import (
	"fmt"
	"strings"
)

func main() {
	//тестовые параметры
	fmt.Println(checkUniq("abcd"))      //true
	fmt.Println(checkUniq("abCdefGh"))  //true
	fmt.Println(checkUniq("aAbBcCdD"))  //false
	fmt.Println(checkUniq("abCdefAaf")) //false
	fmt.Println(checkUniq("aabcd"))     //false
}

//пустая структура
type EmptyStruct struct{}

//функция проверки символов в строке на уникальность
func checkUniq(input string) bool {
	//создаем мапу с ключем типа rune и значением типа пустая структура(не занимает места)
	tempMap := make(map[rune]struct{})
	//приводим строку (заранее приведенную в нижний регистр) к массиву рун и пробегаемся по ней циклом range
	for _, r := range []rune(strings.ToLower(input)) {
		//если данная руна уже есть в нашей мапе, то строка не уникальная, выпрыгиваем из цикла и возвращаем false
		if _, ok := tempMap[r]; ok {
			return false
		}
		//если такой руны нет, то добавляем ее в мапу и идем к следующей руне
		tempMap[r] = EmptyStruct{}
	}
	//если за всю строку были только уникальные руны, возвращаем true
	return true
}
