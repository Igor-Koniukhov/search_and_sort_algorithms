package main

import "fmt"

type SortStrategy interface {
	Sort([]int) []int
}

// BubbleSort implements the SortStrategy interface using bubble sort algorithm
type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) []int {
	// Simple bubble sort implementation
	sorted := make([]int, len(data))
	copy(sorted, data)
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}
	return sorted
}

// QuickSort implements the SortStrategy interface using quick sort algorithm
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) []int {
	// Placeholder for quick sort implementation
	// For simplicity, assume it's implemented
	return data // Return the data as-is for demonstration purposes
}

type SortedArray struct {
	sorter SortStrategy
}

func (s *SortedArray) SetSortStrategy(sorter SortStrategy) {
	s.sorter = sorter
}

func (s *SortedArray) Sort(data []int) []int {
	if s.sorter == nil {
		fmt.Println("Sort strategy not set")
		return data
	}
	return s.sorter.Sort(data)
}

func main() {
	data := []int{5, 3, 4, 1, 2}

	sortedArray := &SortedArray{}

	// Set strategy to BubbleSort and sort data
	sortedArray.SetSortStrategy(&BubbleSort{})
	fmt.Println("Sorted with BubbleSort:", sortedArray.Sort(data))

	// Change strategy to QuickSort and sort data
	sortedArray.SetSortStrategy(&QuickSort{})
	// Assuming QuickSort is implemented, data would be sorted differently
	fmt.Println("Sorted with QuickSort:", sortedArray.Sort(data))
}
