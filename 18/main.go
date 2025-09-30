// Реализовать структуру-счётчик,
// которая будет инкрементироваться в конкурентной среде
// (т.е. из нескольких горутин).
// По завершению программы структура должна выводить
// итоговое значение счётчика.
// Подсказка:
// вам понадобится механизм синхронизации,
// например, sync.Mutex или sync/Atomic для безопасного инкремента.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	count atomic.Int64
}

func (c *counter) add() {
	c.count.Add(1)
}

func (c *counter) load() int64 {
	return c.count.Load()
}

func main() {
	var (
		cnt = new(counter)
		wg  sync.WaitGroup
	)

	for range 1_000_000 {
		wg.Go(func() {
			cnt.add()
		})
	}
	wg.Wait()

	res := cnt.load()
	fmt.Println("Значение счетчика:", res)
}
