package main

import (
	"fmt"
	"sync"
	"time"
)

func funIn(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func produce(arr []int) <-chan int {
	ch := make(chan int)
	go func() {
		for n := range arr {
			ch <- n
			time.Sleep(100)
		}
		close(ch)
	}()

	return ch

}

func main() {
	ch1 := produce([]int{1, 2, 3, 4})
	ch2 := produce([]int{5, 6, 7, 8})

	m := funIn(ch1, ch2)

	for n := range m {
		fmt.Println(n)
	}

}
