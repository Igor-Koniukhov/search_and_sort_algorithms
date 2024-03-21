package main

import (
	"fmt"
)

// Stack represents a stack that holds a slice.
type Stack struct {
	items []int
}

// Push adds an item to the top of the stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top item of the stack.
// If the stack is empty, it returns -1
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex] // Resize the slice to remove the last element
	return item
}

// Peek returns the top item of the stack without removing it.
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		return -1 // Or handle the error as per your requirement
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack.IsEmpty()) // true

	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Peek())    // 2
	fmt.Println(myStack.Pop())     // 2
	fmt.Println(myStack.Size())    // 1
	fmt.Println(myStack.IsEmpty()) // false
}
