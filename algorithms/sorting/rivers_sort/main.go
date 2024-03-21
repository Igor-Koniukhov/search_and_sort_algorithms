package main

import (
	"fmt"
)

func riversSort(arr []int) []int {
	var n = len(arr) - 1
	fmt.Println(n)

	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i] = arr[n-i], arr[i]
	}
	return arr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(riversSort(arr1))

}
