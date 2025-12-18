// Разработать программу нахождения расстояния между двумя точками
// на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными
// (приватными) полями x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка: используйте функцию-конструктор NewPoint(x, y),
// Point и метод Distance(other Point) float64.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	a := NewPoint(6.3, 9.6)
	b := NewPoint(8.13, 22.0)

	distance := a.Distance(b)

	fmt.Printf("Расстояние между точками: %.3f\n", distance)
}
