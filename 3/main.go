// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные
// из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров
// и при старте создавать указанное число горутин-воркеров.

package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type workerPool struct {
	workersNum int
	wg         sync.WaitGroup
	jobs       chan time.Time
}

func newWorkerPool(workersNum, bufferSize int) *workerPool {
	return &workerPool{
		workersNum: workersNum,
		jobs:       make(chan time.Time, bufferSize),
	}
}

func (wp *workerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		data, ok := <-wp.jobs
		if !ok {
			fmt.Printf("Worker %d: channel closed, shutting down...\n", id)
			return
		}
		fmt.Printf("Worker %d processing %v\n", id, data)
		time.Sleep(100 * time.Millisecond)
	}
}

func (wp *workerPool) start() {
	for i := 1; i <= wp.workersNum; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *workerPool) stop() {
	fmt.Println("Initiating drain...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for len(wp.jobs) > 0 {
		select {
		case <-ctx.Done():
			fmt.Println("Drain timeout, forcing close.")
			close(wp.jobs)
			return
		default:
			fmt.Printf("Draining channel: %d items left\n", len(wp.jobs))
			time.Sleep(50 * time.Millisecond)
		}
	}

	close(wp.jobs)
	fmt.Println("Channel closed. Waiting for workers to finish...")

	wp.wg.Wait()
	fmt.Println("All workers completed. Exiting.")
}

func main() {
	workersNum := flag.Int("workers", 10, "Enter workers number")
	chBuffer := flag.Int("buffer", 0, "Enter channel buffer size")
	flag.Parse()

	if *workersNum <= 0 {
		fmt.Println("Workers count must be > 0")
		os.Exit(1)
	}

	if *chBuffer < 0 {
		fmt.Println("Buffer size must be not negative")
		os.Exit(1)
	}

	pool := newWorkerPool(*workersNum, *chBuffer)
	pool.start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-sigCh:
			fmt.Println("\nReceived shutdown signal.")
			pool.stop()
			return
		default:
			pool.jobs <- time.Now()
			time.Sleep(time.Duration(rand.N(50)) * time.Millisecond)
		}
	}
}
