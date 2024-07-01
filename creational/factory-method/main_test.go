package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVehicleFactory(t *testing.T) {
	tests := []struct {
		name           string
		factory        VehicleFactory
		expectedOutput string
	}{
		{
			name:           "Car Factory",
			factory:        &CarFactory{},
			expectedOutput: "Driving a car",
		},
		{
			name:           "Bike Factory",
			factory:        &BikeFactory{},
			expectedOutput: "Riding a bike",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vehicle := tt.factory.NewVehicle()
			assert.Equal(t, tt.expectedOutput, vehicle.Drive())
		})
	}
}
