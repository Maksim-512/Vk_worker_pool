package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestAddWorker(t *testing.T) {
	wp := newWorkerPool()

	wp.addWorker()
	if wp.countWorker != 1{
		t.Errorf("Expected 1 worker, actual %d", wp.countWorker)
	}

	wp.addWorker()
	if wp.countWorker != 2{
		t.Errorf("Expected 2 worker, actual %d", wp.countWorker)
	}
}


func TestRemoveWorker(t *testing.T) {
	wp := newWorkerPool()

	wp.addWorker()
	wp.addWorker()

	wp.removeWorker()

	if wp.countWorker != 1{
		t.Errorf("Expected 1 worker, actual %d", wp.countWorker)
	}
}


func TestAddTasks(t *testing.T) {
	wp := newWorkerPool()

	var taskCounter int64
	wp.addWorker()

	go func() {
		for range wp.textChan{
			atomic.AddInt64(&taskCounter, 1)
		}
	}()

	wp.addTasks("first task")
	wp.addTasks("second task")
	wp.addTasks("third task")

	time.Sleep(time.Millisecond * 500)

	if atomic.LoadInt64(&taskCounter) != 3{
		t.Errorf("Expected 1 tasks, actual %d", taskCounter)
	}

	wp.wait()
}


func TestWair(t *testing.T) {
	wp := newWorkerPool()

	wp.addWorker()

	wp.addTasks("first task")
	wp.addTasks("second task")

	done := make(chan struct{})
	go func() {
		wp.wait()
		close(done)
	}()

	select {
	case <- done:
		t.Log("Success")
	case <- time.After(time.Second * 1):
		t.Error("wait did not complete in time")
	}	
}