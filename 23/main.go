package main

import "fmt"

func main() {
	//создаем тестовый слайс типа string
	testSlice := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	//выводим в stdout результат работы функции deleteFromSlice (новый слайс ссылающийся на исходный testSlice)
	fmt.Println(deleteFromSlice(testSlice, 7))
}

//функция удаления элемента по i (position)
func deleteFromSlice(slice []string, position int) []string {
	//копируем слайс с позиции position+1 до последнего элемента и вставляем на позицию position
	copy(slice[position:], slice[position+1:])
	//возвращаем новый слайс (ссылающийся на исходный testSlice) за исключением последнего элемента
	return slice[:len(slice)-1]
}
