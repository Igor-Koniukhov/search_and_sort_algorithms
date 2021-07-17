package main

import (
	"fmt"
	"time"
)

// Time Complexity O(n^2)
// Space Complexity O(1)

func bubbleSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if items[j] < items[i] {
				var tmp = items[j]
				items[j] = items[i]
				items[i] = tmp
			}
		}
	}
	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := bubbleSort(items)
	// sortItems is {1, 2, 3, 4, 5}

	// *** simplified speed test ***

	items = make([]int, 200)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	var tmp = items[5]
	items[5] = items[6]
	items[6] = tmp
	count := 10000
	start := time.Now()

	for i := 0; i < count; i++ {
		bubbleSort(items)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanoseconds)
	// about 683 000 000 nanoseconds
}