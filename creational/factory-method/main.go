/*
Factory Method Example:
Defines an interface for creating an object but lets
subclasses alter the type of objects that will be created.
*/
package main

import "fmt"

// Product Interface
type Vehicle interface {
	Drive() string
}

// Concrete Product A
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

// Concrete Product B
type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a bike"
}

// Creator Interface
type VehicleFactory interface {
	NewVehicle() Vehicle
}

// Concrete Creator A
type CarFactory struct{}

func (c *CarFactory) NewVehicle() Vehicle {
	return &Car{}
}

// Concrete Creator B
type BikeFactory struct{}

func (b *BikeFactory) NewVehicle() Vehicle {
	return &Bike{}
}

func main() {
	var factory VehicleFactory
	var vehicle Vehicle

	factory = &CarFactory{}
	vehicle = factory.NewVehicle()
	fmt.Println(vehicle.Drive())

	factory = &BikeFactory{}
	vehicle = factory.NewVehicle()
	fmt.Println(vehicle.Drive())
}
