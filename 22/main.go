// Разработать программу,
// которая перемножает, делит, складывает, вычитает две числовых переменных a, b,
// значения которых > 2^20 (больше 1 миллион).

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1 << 40)
	fmt.Printf("a = %v\n", a)

	b := big.NewInt(1 << 30)
	fmt.Printf("b = %v\n", b)

	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сумма %v и %v = %v\n", a, b, sum)

	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Разность %v и %v = %v\n", a, b, sub)

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение %v и %v = %v\n", a, b, mul)

	div := new(big.Int).Div(a, b)
	fmt.Printf("Частное %v и %v = %v\n", a, b, div)
}
