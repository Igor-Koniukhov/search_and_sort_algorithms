package main

import "fmt"

func moveZerosToEnd(arr []int) []int {

	position := 0
	for _, num := range arr {
		if num != 0 {
			arr[position] = num
			position++
		}
	}

	for i := position; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

func main() {
	arr1 := []int{0, 1, 0, 3, 7, 0, 8, 9}
	fmt.Println(moveZerosToEnd(arr1))

}
