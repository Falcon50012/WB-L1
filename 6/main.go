// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления,
// через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Выход по условию
	wg.Go(func() {
		run := true
		if run {
			fmt.Println("G stopped by condition")
			return
		}
	})

	// Выход по закрытию канала
	ch := make(chan struct{})
	wg.Go(func() {
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("G stopped by closed channel")
				return
			}
		}
	})
	close(ch)

	// Выход с помощью тикера
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	wg.Go(func() {
		for i := 1; ; i++ {
			select {
			case <-ticker.C:
				fmt.Println("G stopped by ticker")
				return
			default:
				fmt.Println("WB-GO", i)
				time.Sleep(1 * time.Second)
			}
		}
	})

	// Выход с использованием контекста
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	wg.Go(func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("G stopped by context")
				return
			default:
				fmt.Println("Горутиновый Golang")
				time.Sleep(1 * time.Second)
			}
		}
	})

	// Остановка посредством Goexit
	wg.Go(func() {
		fmt.Println("G stopped by Goexit")
		runtime.Goexit()
	})

	wg.Wait()
}
