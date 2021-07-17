package main

import (
"fmt"
"time"
)

//works when the array is sorted
func binarySearch(arr []int, x int) int {
	i := 0
	j := len(arr)
	for i != j {
		var m = (i + j) / 2
		if x == arr[m] {
			return m
		}
		if x < arr[m] {
			j = m
		} else {
			i = m + 1
		}
	}
	return -1
}

func main() {
	items := []int{2, 3, 5, 7, 11, 13, 17}

	fmt.Println(binarySearch(items, 1))
	//print -1
	fmt.Println(binarySearch(items, 7))
	//print 3
	fmt.Println(binarySearch(items, 19))
	//print -1

	// *** simplified speed test ***

	items = make([]int, 10000000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	count := 100

	start := time.Now()

	for i := 0; i < count; i++ {
		binarySearch(items, 7777777)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds() / int64(count)

	fmt.Println(nanoseconds)
	// about 88 nanoseconds
}