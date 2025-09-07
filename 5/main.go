// Разработать программу, которая будет последовательно
// отправлять значения в канал, а с другой стороны канала –
// читать эти значения.
// По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер
// для ограничения времени работы.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	processingTime := flag.Int("time", 5, "processing time")
	flag.Parse()

	if *processingTime <= 0 {
		fmt.Println("Processing time must be > 0")
		os.Exit(1)
	}

	ch := make(chan int)

	go func() {
		for {
			data, ok := <-ch
			if !ok {
				fmt.Println("Channel closed, exiting...")
				return
			}
			fmt.Println("Reading:", data)
		}
	}()

	timeout := time.After(time.Duration(*processingTime) * time.Second)

	for i := 1; ; i++ {
		select {
		case <-timeout:
			fmt.Println("Closing channel...")
			close(ch)
			return
		default:
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}
}
