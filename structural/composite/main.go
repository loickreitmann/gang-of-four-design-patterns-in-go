/*
Composite Example:
Composes objects into tree structures to represent
part-whole hierarchies, allowing clients to treat
individual objects and compositions uniformly.
*/
package main

import "fmt"

// Component Interface
type Employee interface {
	ShowDetails()
}

// Leaf
type Developer struct {
	Name string
}

func (d *Developer) ShowDetails() {
	fmt.Println("Developer:", d.Name)
}

// Leaf
type Manager struct {
	Name string
}

func (m *Manager) ShowDetails() {
	fmt.Println("Manager:", m.Name)
}

// Composite
type Directory struct {
	Employees []Employee
}

func (d *Directory) ShowDetails() {
	for _, e := range d.Employees {
		e.ShowDetails()
	}
}

func main() {
	dev1 := &Developer{Name: "John"}
	dev2 := &Developer{Name: "Jane"}
	manager := &Manager{Name: "Alice"}

	dir := &Directory{
		Employees: []Employee{dev1, dev2, manager},
	}
	dir.ShowDetails()
}
