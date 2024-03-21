package main

import (
	"fmt"
	"math"
	"time"
)

// Time Complexity O(nk)
// Space Complexity O(n+k)

func listToBuckets(items []int, cBase int, i int) [][]int {
	var buckets = make([][]int, cBase)

	var pBase = int(math.Pow(
		float64(cBase), float64(i)))
	for _, x := range items {
		//Isolate the base-digit from the number
		var digit = (x / pBase) % cBase
		//Drop the number into the correct bucket
		buckets[digit] = append(buckets[digit], x)
	}

	return buckets
}

func bucketsToList(buckets [][]int) []int {
	result := []int{}

	for _, bucket := range buckets {
		result = append(result, bucket...)
	}

	return result
}

func radixSort(arr []int, cBase int) []int {
	maxVal := 0
	for i, value := range arr {
		if i == 0 || value > maxVal {
			maxVal = value
		}
	}

	length := len(arr)
	items := make([]int, length)
	copy(items, arr)

	i := 0
	for math.Pow(float64(cBase), float64(i)) <= float64(maxVal) {
		items = bucketsToList(listToBuckets(items, cBase, i))
		i++
	}

	return items
}

func main() {
	items := []int{4, 1, 5, 3, 2}

	sortItems := radixSort(items, 10)
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
		radixSort(items, 10)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds()

	fmt.Println(sortItems)
	fmt.Println(nanoseconds)
	// about 532 200 000 nanoseconds
}