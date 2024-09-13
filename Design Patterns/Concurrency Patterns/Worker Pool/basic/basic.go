package main

import (
	"fmt"
	"sync"
	"time"
)

// task definition
type Task struct {
	ID int
}

// way to process the tasks
func (t *Task) Process() {
	fmt.Printf("Processing task %d\n", t.ID)
	// simulate a time consuming process
	time.Sleep(2 * time.Second)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan    chan Task
	wg          sync.WaitGroup
}

// functions to excute the worker pool

func (wp *WorkerPool) worker(){
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// initalise the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks))

	// start workers 
	for i:=0; i< wp.concurrency; i++{
		go wp.worker()
	}

	// send tasks to the tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _,task := range wp.Tasks {
		wp.tasksChan <- task
	}

	close(wp.tasksChan)

	// wait for all tasks to finish
	wp.wg.Wait()
}


func main(){
	// create new tasks 
	tasks := make([]Task, 20)
	for i:=0; i<20; i++{
		tasks[i] = Task{ID: i+1}
	}

	// Create a worker Pool
	wp := WorkerPool{
		Tasks: tasks,
		concurrency: 5,
	}

	wp.Run()
	fmt.Println("All task has been proccessed")
}	