package main

import "fmt"

func main() {
	actions := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	batchSize := 3
	fmt.Println(makeBatch(actions, batchSize))
}

func makeBatch(arr []int, batchSize int) [][]int {
	batches := make([][]int, 0, (len(arr)+batchSize-1)/batchSize)

	for batchSize < len(arr) {
		arr, batches = arr[batchSize:], append(batches, arr[0:batchSize])
	}

	batches = append(batches, arr)
	return batches
}
