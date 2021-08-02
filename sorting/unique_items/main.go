package main

import (
	"fmt"
	"sort"
)

var arr = []int{23, 23, 34, 34, 4, 6, 2, 3, 78}

func main() {
	p := Unique(arr)
	fmt.Println(p)
}

func Unique(arr []int) []int {
	var m = make(map[int]bool)
	var newArr []int
	for _, valueArr := range arr {
		if _, value := m[valueArr]; !value {
			m[valueArr] = true
			newArr = append(newArr, valueArr)
		}
	}
	sort.Ints(newArr)
	return newArr
}
