package main

import "fmt"

type Logger interface {
	Log(message string, severity int)
	SetNext(logger Logger)
}

const (
	Info = iota
	Warning
	Error
)

type BaseLogger struct {
	severity int
	next     Logger
}

func (l *BaseLogger) SetNext(next Logger) {
	l.next = next
}

type InfoLogger struct {
	BaseLogger
}

func NewInfoLogger() *InfoLogger {
	return &InfoLogger{BaseLogger{severity: Info}}
}

func (l *InfoLogger) Log(message string, severity int) {
	if severity == Info {
		fmt.Println("Info:", message)
		return
	}
	if l.next != nil {
		l.next.Log(message, severity)
	}
}

type WarningLogger struct {
	BaseLogger
}

func NewWarningLogger() *WarningLogger {
	return &WarningLogger{BaseLogger{severity: Warning}}
}

func (l *WarningLogger) Log(message string, severity int) {
	if severity == Warning {
		fmt.Println("Warning:", message)
		return
	}
	if l.next != nil {
		l.next.Log(message, severity)
	}
}

type ErrorLogger struct {
	BaseLogger
}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{BaseLogger{severity: Error}}
}

func (l *ErrorLogger) Log(message string, severity int) {
	if severity == Error {
		fmt.Println("Error:", message)
		return
	}
	if l.next != nil {
		l.next.Log(message, severity)
	}
}

func main() {
	infoLogger := NewInfoLogger()
	warningLogger := NewWarningLogger()
	errorLogger := NewErrorLogger()

	infoLogger.SetNext(warningLogger)
	warningLogger.SetNext(errorLogger)

	// The request will be handled by the InfoLogger.
	infoLogger.Log("This is an informational message.", Info)

	// The request will be passed to the WarningLogger by the InfoLogger.
	infoLogger.Log("This is a warning message.", Warning)

	// The request will be passed all the way to the ErrorLogger.
	infoLogger.Log("This is an error message.", Error)
}
