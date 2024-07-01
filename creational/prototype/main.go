/*
Prototype Example:
Specifies the kinds of objects to create using a
prototypical instance and creates new objects by
copying this prototype.
*/
package main

import "fmt"

// Prototype Interface
type Cloneable interface {
	Clone() Cloneable
}

// Concrete Prototype
type Item struct {
	Name  string
	Price float64
}

func (i *Item) Clone() Cloneable {
	return &Item{
		Name:  i.Name,
		Price: i.Price,
	}
}

func main() {
	item1 := &Item{Name: "Item1", Price: 10.0}
	item2 := item1.Clone().(*Item)
	item2.Name = "Item2"

	fmt.Printf("Original: %+v\n", item1)
	fmt.Printf("Clone: %+v\n", item2)
}
