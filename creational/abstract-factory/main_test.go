package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShirtFactory(t *testing.T) {
	tests := []struct {
		name          string
		factory       ShirtFactory
		expectedColor string
	}{
		{
			name:          "Red Shirt Factory",
			factory:       &RedShirtFactory{},
			expectedColor: "Red",
		},
		{
			name:          "Blue Shirt Factory",
			factory:       &BlueShirtFactory{},
			expectedColor: "Blue",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shirt := tt.factory.CreateShirt()
			assert.Equal(t, tt.expectedColor, shirt.GetColor())
		})
	}
}
