// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации
// (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	m := make(map[string]int)

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[fmt.Sprintf("G#%d load", i)] = i * i
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(len(m))
	fmt.Println(m)
}
