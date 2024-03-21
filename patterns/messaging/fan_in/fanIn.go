package main

import "fmt"

func produce(numbers []int, channel chan<- int) {
	for _, n := range numbers {
		channel <- n
	}
	close(channel)
}

func fanIn(input1, input2 <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for {
			select {
			case n, ok := <-input1:
				if !ok {
					input1 = nil
				} else {
					output <- n
				}
			case n, ok := <-input2:
				if !ok {
					input2 = nil
				} else {
					output <- n
				}
			}
			if input1 == nil && input2 == nil {
				break
			}
		}
		close(output)
	}()
	return output
}

func main() {
	input1 := make(chan int)
	input2 := make(chan int)

	// Start producers
	go produce([]int{1, 3, 5, 7, 9}, input1)
	go produce([]int{2, 4, 6, 8, 10}, input2)

	// Fan-in the inputs
	output := fanIn(input1, input2)

	// Consume the merged output
	for n := range output {
		fmt.Println(n)
	}
}
