// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные
// из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров
// и при старте создавать указанное число горутин-воркеров.

package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

type workerPool struct {
	workersNum int
	jobs       chan time.Time
}

func newWorkerPool(workersNum int) *workerPool {
	return &workerPool{
		workersNum: workersNum,
		jobs:       make(chan time.Time),
	}
}

func (wp *workerPool) worker(id int) {
	for data := range wp.jobs {
		fmt.Printf("Worker %d processing %v\n", id, data)
		time.Sleep(100 * time.Millisecond)
	}
}

func (wp *workerPool) start() {
	for i := 1; i <= wp.workersNum; i++ {
		go wp.worker(i)
	}
}

func main() {
	workersNum := flag.Int("workers", 10, "Enter workers number")
	flag.Parse()

	if *workersNum <= 0 {
		fmt.Println("Workers count must be > 0")
		os.Exit(1)
	}

	pool := newWorkerPool(*workersNum)
	pool.start()

	for {
		pool.jobs <- time.Now()
		time.Sleep(100 * time.Millisecond)
	}
}
