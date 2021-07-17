package main
import (
	"fmt"
	"time"
)

func linearSearch(arr []int, x int) int {
	i := 0
	count := len(arr)
	for i < count {
		if arr[i] == x {
			return i
		}
		i++
	}
	return -1
}

func main() {
	items := []int{2, 3, 5, 7, 11, 13, 17}

	fmt.Println(linearSearch(items, 1))
	//print -1
	fmt.Println(linearSearch(items, 7))
	//print 3
	fmt.Println(linearSearch(items, 19))
	//print -1

	// *** simplified speed test ***

	items = make([]int, 10000000)
	for i := 0; i < len(items); i++ {
		items[i] = i
	}
	count := 100

	start := time.Now()

	for i := 0; i < count; i++ {
		linearSearch(items, 7777777)
	}

	delta := time.Now().Sub(start)
	nanoseconds := delta.Nanoseconds() / int64(count)

	fmt.Println(nanoseconds)
	// about 20 910 000 nanoseconds
}