package main

import (
	"fmt"
	"time"
)

// works when the array is sorted
func interpolationSearch(arr []int, value int) int {
	lowInd := 0
	highInd := len(arr) - 1

	for (arr[lowInd] < value) && (value < arr[highInd]) {
		var midIndex = lowInd + ((value-arr[lowInd])*(highInd-lowInd))/(arr[highInd]-arr[lowInd])

		if arr[midIndex] < value {
			lowInd = midIndex + 1
		} else if arr[midIndex] > value {
			highInd = midIndex - 1
		} else {
			return midIndex
		}
	}

	if arr[lowInd] == value {
		return lowInd
	}
	if arr[highInd] == value {
		return highInd
	}
	return -1
}

func main() {
	items := []int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(interpolationSearch(items, 1))
	//print -1
	fmt.Println(interpolationSearch(items, 7))
	//print 3
	fmt.Println(interpolationSearch(items, 19))
	//print -1

	// *** simplified speed test ***

	items = make([]int, 10000000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	count := 100

	start := time.Now()

	for i := 0; i < count; i++ {
		interpolationSearch(items, 7777777)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds() / int64(count)

	fmt.Println(nanoseconds)
	// about 41 nanoseconds
}
