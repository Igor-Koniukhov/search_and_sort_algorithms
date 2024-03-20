package main

import "fmt"

type Prototype interface {
	Clone() Prototype
}

type File struct {
	Name    string
	Content string
}

func (f *File) Clone() Prototype {
	// Create a new File object and copy the values.
	return &File{Name: f.Name + "_copy", Content: f.Content}
}

func main() {
	originalFile := &File{Name: "original", Content: "This is the original file"}
	clonedFile := originalFile.Clone()

	fmt.Printf("Original: %#v\n", originalFile)
	fmt.Printf("Cloned: %#v\n", clonedFile)
	fmt.Printf("Cloned clone: %#v\n", clonedFile.Clone())
}
