package main

import "fmt"

// Logger is the product interface
type Logger interface {
	Log(message string)
}

// ConsoleLogger is a concrete product
type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
	fmt.Println("ConsoleLogger:", message)
}

// FileLogger is another concrete product
type FileLogger struct{}

func (f *FileLogger) Log(message string) {
	fmt.Println("FileLogger:", message)
}

// LoggerFactory is the creator interface
type LoggerFactory interface {
	CreateLogger() Logger
}

// ConsoleLoggerFactory is a concrete creator
type ConsoleLoggerFactory struct{}

func (c *ConsoleLoggerFactory) CreateLogger() Logger {
	return &ConsoleLogger{}
}

// FileLoggerFactory is another concrete creator
type FileLoggerFactory struct{}

func (f *FileLoggerFactory) CreateLogger() Logger {
	return &FileLogger{}
}

func main() {
	consoleFactory := &ConsoleLoggerFactory{}
	consoleLogger := consoleFactory.CreateLogger()
	consoleLogger.Log("This is a console log.")

	fileFactory := &FileLoggerFactory{}
	fileLogger := fileFactory.CreateLogger()
	fileLogger.Log("This is a file log.")
}
