// Написать программу, которая конкурентно
// рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10],
// и выведет результаты в stdout.
// Подсказка: запусти несколько горутин,
// каждая из которых возводит число в квадрат.

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	nums := []int{2, 4, 6, 8, 10}
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(n)
	}
	wg.Wait()
}
