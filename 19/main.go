// Разработать программу,
// которая переворачивает подаваемую на вход строку.
// Например:
// при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode
// (русские буквы, emoji и пр.), то есть просто iterating по байтам
// может не подойти — нужен срез рун ([]rune).

package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	runes := []rune(str)
	var res strings.Builder
	res.Grow(len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		res.WriteRune(runes[i])
	}
	return res.String()
}

func reverseStringTwoPointers(str string) string {
	runes := []rune(str)
	l, r := 0, len(runes)-1
	for l < r {
		runes[l], runes[r] = runes[r], runes[l]
		l++
		r--
	}
	return string(runes)
}

func main() {
	str := "😊BW в ved oG ьтатс юалеЖ✨"
	fmt.Printf("input: %q\n", str)
	fmt.Printf("reverseString: %q\n", reverseString(str))
	fmt.Printf("reverseStringTwoPointers: %q\n", reverseStringTwoPointers(str))
}
