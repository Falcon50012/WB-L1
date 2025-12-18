// Разработать программу, которая проверяет,
// что все символы в строке встречаются один раз
// (т.е. строка состоит из уникальных символов).

// Вывод: true, если все символы уникальны,
// false, если есть повторения.
// Проверка должна быть регистронезависимой.

package main

import (
	"fmt"
	"unicode"
)

func uniqueChars(s string) bool {
	uniqueCharsMap := make(map[rune]bool)

	for _, c := range s {
		lowChars := unicode.ToLower(c)

		if uniqueCharsMap[lowChars] {
			return false
		}

		uniqueCharsMap[lowChars] = true
	}

	return true
}

func main() {
	testStrings := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
	}

	for _, ts := range testStrings {
		fmt.Printf("%q -> %v\n", ts, uniqueChars(ts))
	}
}
