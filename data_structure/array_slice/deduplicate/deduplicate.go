package main

import (
	"fmt"
	"sort"
)

func main() {
	in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1} // any item can be sorted
	sort.Ints(in)
	fmt.Println(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		fmt.Println(j, i, " continue")
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}
	result := in[:j+1]
	fmt.Println(result) // [1 2 3 4]
}
