package main

import (
	"fmt"
	"sync"
)

// Singleton represents the singleton instance in this example.
type Singleton struct{}

var instance *Singleton
var once sync.Once

// GetInstance returns the singleton instance, creating it if necessary.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	// Attempt to create two instances.
	instance1 := GetInstance()
	instance2 := GetInstance()

	// Check if instances are the same.
	if instance1 == instance2 {
		fmt.Println("Both instances are the same.")
	} else {
		fmt.Println("Instances are different.")
	}
}
