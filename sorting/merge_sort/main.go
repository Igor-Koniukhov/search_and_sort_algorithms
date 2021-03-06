package main

import (
	"fmt"
	"time"
)

// Time Complexity O(n log(n)))
// Space Complexity O(n)

func mergeSort(arr []int) []int {
	items := make([]int, len(arr))
	copy(items, arr)
	doMergeSort(items)
	return items
}

func doMergeSort(items []int) {
	length := len(items)
	if length == 1 {
		return
	}

	var lLeft = length / 2
	var left = make([]int, lLeft)
	copy(left, items[:lLeft])
	var lRight = length - lLeft
	var right = make([]int, lRight)
	copy(right, items[lLeft:])

	doMergeSort(left)
	doMergeSort(right)

	merge(left, right, items)
}

func merge(left []int, right []int, result []int) {
	l := 0
	r := 0
	i := 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[i] = left[l]
			l++
		} else {
			result[i] = right[r]
			r++
		}
		i++
	}
	var length = len(left) - l
	copy(result[i:i+length], left[l:])
	i = i + length
	length = len(right) - r
	copy(result[i:i+length], right[r:])
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := mergeSort(items)
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
		mergeSort(items)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanoseconds)
	// about 439 000 000 nanoseconds
}