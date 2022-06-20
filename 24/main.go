package main

import (
	"fmt"
	"math"
)

//создаем структуру Point с 2мя параметрами типа float
type Point struct {
	x, y float64
}

//инициализируем структуру Point
func NewPoint(x, y float64) Point {
	return Point{
		//сразу убираем минус
		x: math.Abs(x),
		y: math.Abs(y),
	}
}

//при помощи пакета math вычисляем расстояние между точками
func (p *Point) GetDistance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p.x+p2.x, 2) + math.Pow(p.y+p2.y, 2))
}

func main() {
	//инициализируем две структуры
	p := NewPoint(0, 2)
	p2 := NewPoint(0, -2)

	//вычисляем расстояние
	result := p.GetDistance(p2)
	//выводим результат
	fmt.Printf("%.2f", result)
}
