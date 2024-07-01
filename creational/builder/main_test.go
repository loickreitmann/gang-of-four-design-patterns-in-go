package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseBuilder(t *testing.T) {
	tests := []struct {
		name          string
		windowType    string
		doorType      string
		floor         int
		expectedHouse House
	}{
		{
			name:       "Wooden House",
			windowType: "Wooden Window",
			doorType:   "Wooden Door",
			floor:      2,
			expectedHouse: House{
				WindowType: "Wooden Window",
				DoorType:   "Wooden Door",
				Floor:      2,
			},
		},
		{
			name:       "Glass House",
			windowType: "Glass Window",
			doorType:   "Glass Door",
			floor:      1,
			expectedHouse: House{
				WindowType: "Glass Window",
				DoorType:   "Glass Door",
				Floor:      1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := &HouseBuilder{}
			house := builder.SetWindowType(tt.windowType).SetDoorType(tt.doorType).SetFloor(tt.floor).Build()
			assert.Equal(t, tt.expectedHouse, house)
		})
	}
}
