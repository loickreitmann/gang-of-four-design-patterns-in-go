package main

import "fmt"

type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

type HouseBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (b *HouseBuilder) SetWindowType(windowType string) *HouseBuilder {
	b.windowType = windowType
	return b
}

func (b *HouseBuilder) SetDoorType(doorType string) *HouseBuilder {
	b.doorType = doorType
	return b
}

func (b *HouseBuilder) SetFloor(floor int) *HouseBuilder {
	b.floor = floor
	return b
}

func (b *HouseBuilder) Build() House {
	return House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		Floor:      b.floor,
	}
}

func main() {
	builder := &HouseBuilder{}
	house := builder.SetWindowType("Wooden Window").SetDoorType("Wooden Door").SetFloor(2).Build()
	fmt.Printf("House: %+v\n", house)
}
