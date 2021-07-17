package main

import (
	"fmt"
	"time"
)

//works when the array is sorted
func interpolationSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	for (arr[low] < x) && (x < arr[high]) {
		var mid = low + ((x-arr[low])*(high-low))/(arr[high]-arr[low])

		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}

	if arr[low] == x {
		return low
	}
	if arr[high] == x {
		return high
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

