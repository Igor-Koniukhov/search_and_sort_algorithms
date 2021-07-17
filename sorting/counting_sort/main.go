package main

import (
	"fmt"
	"math"
	"time"
)

// Time Complexity O(n+k)
// Space Complexity O(k)

func countingSort(arr []int) []int {
	length := len(arr)
	items := make([]int, length)
	copy(items, arr)

	var min = math.MaxInt32
	var max = math.MinInt32
	for _, x := range arr {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	var counts = make([]int, max-min+1)

	for _, x := range arr {
		counts[x-min]++
	}

	var total = 0
	for i := min; i <= max; i++ {
		var oldCount = counts[i-min]
		counts[i-min] = total
		total += oldCount
	}

	for _, x := range arr {
		items[counts[x-min]] = x
		counts[x-min]++
	}
	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := countingSort(items)
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
		countingSort(items)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanoseconds)
	// about 58 000 000 nanoseconds
}