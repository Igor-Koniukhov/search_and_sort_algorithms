package main

import "fmt"

func worker(id int, input <-chan int, output chan<- int) {
	for n := range input {
		fmt.Printf("Worker %d received %d\n", id, n)
		output <- n * n
	}
}

func fanOut(input <-chan int, numWorkers int) []<-chan int {
	outputs := make([]<-chan int, numWorkers)

	for i := 0; i < numWorkers; i++ {
		output := make(chan int)
		outputs[i] = output

		go worker(i, input, output)
	}

	return outputs
}

func main() {
	input := make(chan int)
	numWorkers := 4

	// Start Fan-Out
	outputs := fanOut(input, numWorkers)

	// Send tasks to the input channel
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
		}
		close(input)
	}()

	// Optionally, you can implement a Fan-In pattern here to aggregate results
	// For simplicity, we'll just read from the first worker's output in this example
	for n := range outputs[0] {
		fmt.Println("Result:", n)
	}
}
