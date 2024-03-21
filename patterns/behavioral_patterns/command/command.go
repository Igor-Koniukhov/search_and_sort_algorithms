package main

import "fmt"

type Command interface {
	Execute()
}

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Light is off")
}

// ConcreteCommand
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// ConcreteCommand
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

type Switch struct {
	command Command
}

func (s *Switch) StoreAndExecute(command Command) {
	s.command = command
	s.command.Execute()
}

func main() {
	light := &Light{}
	lightOnCommand := &LightOnCommand{light: light}
	lightOffCommand := &LightOffCommand{light: light}

	switchButton := &Switch{}

	switchButton.StoreAndExecute(lightOnCommand)
	switchButton.StoreAndExecute(lightOffCommand)
}
