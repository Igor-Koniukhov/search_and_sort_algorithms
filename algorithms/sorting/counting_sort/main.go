package main

import (
	"fmt"
	"time"
)

// Time Complexity O(n+k)
// Space Complexity O(k)

func countingSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	maxN, minN := arr[0], arr[0]
	for _, x := range arr {
		if x > maxN {
			maxN = x
		}
		if x < minN {
			minN = x
		}
	}

	counts := make([]int, maxN-minN+1)
	for _, x := range arr {
		counts[x-minN]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	sortedArr := make([]int, len(arr))
	//items := []int{4, 2, 2, 1, 1, 1, 10, 10, 1, 3, 44}
	for i := len(arr) - 1; i >= 0; i-- {
		counts[arr[i]-minN]--
		sortedArr[counts[arr[i]-minN]] = arr[i]
	}

	return sortedArr
}

func main() {
	items := []int{4, 2, 2, 1, 1, 1, 10, 10, 1, 3, 44}

	sortItems := countingSort(items)
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
