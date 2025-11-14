// Рассмотреть следующий код и ответить на вопросы:
// к каким негативным последствиям он может привести
// и как это исправить?

// Что происходит с переменной justString?

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// ОТВЕТ: Поскольку переменная justString является глобальной и данные на которые она указывает,
// хранятся в куче,
// при присваивании ей результата функии createHugeString в виде подстроки,
// произойдет утечка памяти, т.к. переменная justString будет указывать на тот же
// низлежащий массив, что и результат createHugeString.

//====================================================================================

// КОРРЕКТНЫЙ ВАРИАНТ:

package main

import (
	"strings"
)

var justString string

func createHugeString(size int) string {
	result := make([]byte, size)
	for i := range result {
		result[i] = 'a'
	}
	return string(result)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100]) // Копируем 100 байт
}

func main() {
	someFunc()
}
