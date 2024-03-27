package main

import (
	"fmt"
)

func shufle(a []int) []int {
	j := 0

	for i := (len(a) - 1); i > 0; i-- {
		fmt.Println(j)
		//j := rand.Intn(i + 1)
		temp := a[j]
		a[j] = a[i]
		a[i] = temp
		j++
	}
	return a
}

func main() {
	arr := []int{3, 2, 5, 23, 45, 87, 98}
	fmt.Println(shufle(arr))

}
