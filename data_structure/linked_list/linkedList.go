package main

import "fmt"

// Node represents an element of the linked list
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	Head *Node
}

// Append adds a new node with the specified data at the end of the list
func (l *LinkedList) Append(data int) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	lastNode := l.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
}

// Prepend adds a new node with the specified data at the beginning of the list
func (l *LinkedList) Prepend(data int) {
	newNode := &Node{Data: data}
	newNode.Next = l.Head
	l.Head = newNode
}

// DeleteWithValue deletes the first node with the specified data
func (l *LinkedList) DeleteWithValue(data int) {
	if l.Head == nil {
		return
	}
	if l.Head.Data == data {
		l.Head = l.Head.Next
		return
	}
	previousNode := l.Head
	for previousNode.Next != nil {
		if previousNode.Next.Data == data {
			previousNode.Next = previousNode.Next.Next
			return
		}
		previousNode = previousNode.Next
	}
}

// Display prints out the elements of the list
func (l *LinkedList) Display() {
	currentNode := l.Head
	for currentNode != nil {
		fmt.Printf("%d -> ", currentNode.Data)
		currentNode = currentNode.Next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Prepend(0)
	list.Display() // Expected: 0 -> 1 -> 2 -> nil

	list.DeleteWithValue(1)
	list.Display() // Expected: 0 -> 2 -> nil
}
