package workerpool

import (
	"log"
	"time"
)

type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	MaxWorker  int
	QueuedTask chan func()
}

func (w *workerPool) Run() {

	for i := 0; i < w.MaxWorker; i++ {
		go func(workerId int) {
			for task := range w.QueuedTask {
				task()
			}
		}(i + 1)
	}
}

func NewWorkerPool(maxWorker int) WorkerPool {
	return &workerPool{
		MaxWorker:  maxWorker,
		QueuedTask: make(chan func()),
	}
}

func (w *workerPool) AddTask(task func()) {
	w.QueuedTask <- task
}

func WorkerPoolFunc() {

	totalWorker := 3
	wp := NewWorkerPool(totalWorker)

	wp.Run()

	type Result struct {
		id    int
		value int
	}

	totalTask := 5
	resultC := make(chan Result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.AddTask(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(5 * time.Second)
			resultC <- Result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC

		log.Printf("[main] Task %d has been finished with result %d:", res.id, res.value)

	}

}
