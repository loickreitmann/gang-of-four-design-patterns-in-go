package main

import "fmt"

// Abstract Product
type Shirt interface {
	GetColor() string
}

// Concrete Product A
type RedShirt struct{}

func (r *RedShirt) GetColor() string {
	return "Red"
}

// Concrete Product B
type BlueShirt struct{}

func (b *BlueShirt) GetColor() string {
	return "Blue"
}

// Abstract Factory
type ShirtFactory interface {
	CreateShirt() Shirt
}

// Concrete Factory A
type RedShirtFactory struct{}

func (r *RedShirtFactory) CreateShirt() Shirt {
	return &RedShirt{}
}

// Concrete Factory B
type BlueShirtFactory struct{}

func (b *BlueShirtFactory) CreateShirt() Shirt {
	return &BlueShirt{}
}

func main() {
	var factory ShirtFactory
	var shirt Shirt

	factory = &RedShirtFactory{}
	shirt = factory.CreateShirt()
	fmt.Println("Shirt color:", shirt.GetColor())

	factory = &BlueShirtFactory{}
	shirt = factory.CreateShirt()
	fmt.Println("Shirt color:", shirt.GetColor())
}
