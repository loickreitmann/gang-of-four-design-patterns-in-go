/*
Command Example:
Encapsulates a request as an object, thereby allowing for
parameterization of clients with queues, requests, and operations.
*/
package main

import "fmt"

// Command Interface
type Command interface {
	Execute()
}

// Concrete Command
type PrintCommand struct {
	message string
}

func (p *PrintCommand) Execute() {
	fmt.Println(p.message)
}

// Invoker
type Invoker struct {
	commands []Command
}

func (i *Invoker) StoreCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		command.Execute()
	}
}

func main() {
	invoker := &Invoker{}
	invoker.StoreCommand(&PrintCommand{message: "Hello"})
	invoker.StoreCommand(&PrintCommand{message: "World"})

	invoker.ExecuteCommands()
}
