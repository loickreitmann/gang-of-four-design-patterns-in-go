/*
Iterator Example:
Provides a way to access the elements of an aggregate object
sequentially without exposing its underlying representation.
*/
package main

import "fmt"

// Aggregate
type Aggregate interface {
	Iterator() Iterator
}

// Iterator
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Concrete Aggregate
type Collection struct {
	items []interface{}
}

func (c *Collection) Iterator() Iterator {
	return &CollectionIterator{collection: c, index: 0}
}

// Concrete Iterator
type CollectionIterator struct {
	collection *Collection
	index      int
}

func (i *CollectionIterator) HasNext() bool {
	return i.index < len(i.collection.items)
}

func (i *CollectionIterator) Next() interface{} {
	if i.HasNext() {
		item := i.collection.items[i.index]
		i.index++
		return item
	}
	return nil
}

func main() {
	collection := &Collection{
		items: []interface{}{"A", "B", "C"},
	}
	iterator := collection.Iterator()

	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
