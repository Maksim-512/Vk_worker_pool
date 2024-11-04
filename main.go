package main

import (
	"fmt"
	"sync"
)

type workerPool struct{
	textChan chan string
	controlChan chan struct{}
	wg sync.WaitGroup
	countWorker int
}


func newWorkerPool() *workerPool{
	return &workerPool{
		textChan: make(chan string),
		controlChan: make(chan struct{}),
	}
}


func (w *workerPool) addWorker() {
	w.wg.Add(1)
	workerID := w.countWorker + 1
	w.countWorker++

	go func(id int){
		defer w.wg.Done()
		fmt.Printf("Worker %d started\n", id)

		for{
			select{
			case task, ok := <- w.textChan:
				if !ok{
					fmt.Printf("Worker %d exiting due to closed channel\n", id)
					return
				}
				// time.Sleep(time.Millisecond * 500)
				fmt.Printf("Worker %d performs the task %s\n", id, task)
			case <- w.controlChan:
				fmt.Printf("Worker %d stop\n", id)
				return
			}
		}
	}(workerID)
}


func (w *workerPool) removeWorker(){
	w.controlChan <- struct{}{}
	w.countWorker--
}

func (w *workerPool) addTasks(task string) {
	w.textChan <- task
}


func (w *workerPool) wait(){
	close(w.textChan)
	w.wg.Wait()
}


func main() {
	// wp := newWorkerPool()

	// wp.addWorker()
	// wp.addWorker()

	// wp.addTasks("first")
	// wp.addTasks("second")
	// wp.addTasks("third")
	// wp.addTasks("four")
	// wp.addTasks("five")

	// wp.wait()
}