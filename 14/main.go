// Разработать программу,
// которая в runtime способна определить тип переменной,
// переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).
// Подсказка: оператор типа switch v.(type) поможет в решении.

package main

import (
	"fmt"
	"reflect"
)

func varIdentification(v any) {
	switch v.(type) {
	default:
		fmt.Println(reflect.TypeOf(v).Kind())
	}
}

func main() {
	var (
		a int
		b string
		c bool
		d chan any
	)

	varIdentification(a)
	varIdentification(b)
	varIdentification(c)
	varIdentification(d)
}
