// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».
// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы,
// а выполнять операцию «на месте».

package main

import (
	"fmt"
)

func reverseWords(str string) string {
	runes := []rune(str)
	{
		l, r := 0, len(runes)-1
		for l < r {
			runes[l], runes[r] = runes[r], runes[l]
			l++
			r--
		}
	}
	{
		l, r := 0, len(runes)-1
		b := 0
		for i, c := range runes {
			l = b
			if c == ' ' || i == len(runes)-1 {
				if c == ' ' {
					r = i - 1
				} else {
					r = i
				}
				for l < r {
					runes[l], runes[r] = runes[r], runes[l]
					l++
					r--
				}
				b = i + 1
			}
		}
	}
	return string(runes)
}

func main() {
	str := "snow dog sun"
	fmt.Printf("%q -> %q\n", str, reverseWords(str))
}
