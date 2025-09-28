// Поменять местами два числа без использования временной переменной.
// Подсказка: примените сложение/вычитание или XOR-обмен.

package main

import "fmt"

func main() {
	a, b := 3, 6
	fmt.Println(a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
