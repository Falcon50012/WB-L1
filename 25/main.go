// Реализовать собственную функцию sleep(duration)
// аналогично встроенной функции time.Sleep,
// которая приостанавливает выполнение текущей горутины.

// Важно: в отличие от настоящей time.Sleep,
// ваша функция должна именно блокировать выполнение
// (например, через таймер или цикл),
// а не просто вызывать time.Sleep :) — это упражнение.

package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	fmt.Println("Блокируем основную горутину кастомной sleep")
	sleep(3 * time.Second)
	fmt.Println("Разблокируем основную горутину кастомной sleep")
	fmt.Println("Выход из main")
}
