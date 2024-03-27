package main

import "fmt"

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents the binary tree itself, with a reference to the root node
type BinaryTree struct {
	Root *TreeNode
}

// Insert inserts a new node into the binary tree
func (t *BinaryTree) Insert(value int) {
	newNode := &TreeNode{Value: value}
	if t.Root == nil {
		t.Root = newNode
	} else {
		t.Root.insertNode(newNode)
	}
}

// insertNode is a recursive helper function for inserting a new node
func (n *TreeNode) insertNode(newNode *TreeNode) {
	if newNode.Value < n.Value {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.insertNode(newNode)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.insertNode(newNode)
		}
	}
}

// InOrderTraversal traverses the tree in in-order
func (t *BinaryTree) InOrderTraversal() {
	t.Root.inOrderTraversal()
}

// inOrderTraversal is a recursive helper function for in-order traversal
func (n *TreeNode) inOrderTraversal() {
	if n != nil {
		n.Left.inOrderTraversal()
		fmt.Printf("%d ", n.Value)
		n.Right.inOrderTraversal()
	}
}

// PreOrderTraversal traverses the tree in pre-order
func (t *BinaryTree) PreOrderTraversal() {
	t.Root.preOrderTraversal()
}

// preOrderTraversal is a recursive helper function for pre-order traversal
func (n *TreeNode) preOrderTraversal() {
	if n != nil {
		fmt.Printf("%d ", n.Value)
		n.Left.preOrderTraversal()
		n.Right.preOrderTraversal()
	}
}

// PostOrderTraversal traverses the tree in post-order
func (t *BinaryTree) PostOrderTraversal() {
	t.Root.postOrderTraversal()
}

// postOrderTraversal is a recursive helper function for post-order traversal
func (n *TreeNode) postOrderTraversal() {
	if n != nil {
		n.Left.postOrderTraversal()
		n.Right.postOrderTraversal()
		fmt.Printf("%d ", n.Value)
	}
}

func main() {
	tree := &BinaryTree{}
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	fmt.Println("In-Order Traversal:")
	tree.InOrderTraversal()
	fmt.Println("\nPre-Order Traversal:")
	tree.PreOrderTraversal()
	fmt.Println("\nPost-Order Traversal:")
	tree.PostOrderTraversal()
}
