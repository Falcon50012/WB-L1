// Дана переменная типа int64. Разработать программу,
// которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).

package main

import (
	"fmt"
	"math"
	"os"
)

func setBit(num int64, bit uint8, value uint8) int64 {
	var mask uint64 = 1 << (bit - 1)
	var res int64
	fmt.Printf("Маска =\n%064b\n", mask)

	if value == 0 {
		res = num &^ int64(mask)
	}

	if value == 1 {
		res = num | int64(mask)
	}

	return res
}

func main() {
	var (
		num   int64
		bit   uint8
		value uint8
	)

	fmt.Printf("Введите число в диапазоне от %d до %d: ", math.MinInt64, math.MaxInt64)
	_, err := fmt.Scanln(&num)
	if err != nil {
		fmt.Println("Выход за пределы допустимого диапазона значений", err)
		os.Exit(1)
	}

	fmt.Print("Введите бит: ")
	fmt.Scanln(&bit)
	if bit < 1 || bit > 64 {
		fmt.Println("Выход за пределы допустимого диапазона значений")
		os.Exit(1)
	}

	fmt.Print("Введите значение бита: ")
	fmt.Scanln(&value)
	if value > 1 {
		fmt.Println("Выход за пределы допустимого диапазона значений")
		os.Exit(1)
	}

	fmt.Printf("Введенное число %d в бинарном представлении =\n%064b\n", num, uint64(num))

	res := setBit(num, bit, value)

	fmt.Printf("Резулитирующее значение %d в бинарном представлении =\n%064b\n", res, uint64(res))

	fmt.Println("Программа завершена")
}
