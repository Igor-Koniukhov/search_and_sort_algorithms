package main

import "fmt"

type Task int
type Result int

func worker(id int, tasks <-chan Task, results chan<- Result) {
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		output := Result(task * 2) // Processing task
		results <- output
	}
}

func startWorkerPool(numWorkers int, tasks <-chan Task, results chan<- Result) {
	for i := 0; i < numWorkers; i++ {
		go worker(i, tasks, results)
	}
}

func main() {
	numTasks := 10
	numWorkers := 3

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)

	// Start the worker pool.
	startWorkerPool(numWorkers, tasks, results)

	// Send tasks to be processed.
	for i := 0; i < numTasks; i++ {
		tasks <- Task(i + 1)
	}
	close(tasks) // No more tasks are coming.

	// Collect the results.
	for i := 0; i < numTasks; i++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
	close(results)
}
