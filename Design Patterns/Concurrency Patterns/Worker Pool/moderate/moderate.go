package main

import (
	"fmt"
	"sync"
	"time"
)

// task definition
type Task interface {
	Process()
}

// Email Task Definition
type EmailTask struct {
	Email       string
}

// Process Email Task
func (t *EmailTask) Process() {
	fmt.Printf("Email Processing task %s\n", t.Email)
	// simulate a time consuming process
	time.Sleep(1 * time.Second)
}

// Image Task Definition
type ImageTask struct {
	ImageURL string
}

// Process Image Task
func (t *ImageTask) Process() {
	fmt.Printf("Image Processing task %s\n", t.ImageURL)
	// simulate a time consuming process
	time.Sleep(5 * time.Second)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// functions to execute the worker pool

func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	// wait for all tasks to finish
	wp.wg.Wait()
}

func main() {
	// create new tasks
	tasks := []Task{
		&EmailTask{Email: "1"},
		&ImageTask{ImageURL: "1"},
		&EmailTask{Email: "2"},
		&ImageTask{ImageURL: "2"},
		&EmailTask{Email: "3"},
		&ImageTask{ImageURL: "3"},
		&EmailTask{Email: "4"},
		&ImageTask{ImageURL: "4"},
		&EmailTask{Email: "5"},
		&ImageTask{ImageURL: "5"},
		&EmailTask{Email: "6"},
		&ImageTask{ImageURL: "6"},
		&EmailTask{Email: "7"},
		&ImageTask{ImageURL: "7"},
	}
	

	// Create a worker pool
	wp := WorkerPool{
		Tasks:  tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All tasks have been processed")
}
