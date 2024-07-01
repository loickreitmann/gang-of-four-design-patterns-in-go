package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemClone(t *testing.T) {
	tests := []struct {
		name             string
		original         *Item
		modifyClone      func(*Item)
		expectedOriginal *Item
		expectedClone    *Item
	}{
		{
			name:     "Clone Item1 and modify name to Item2",
			original: &Item{Name: "Item1", Price: 10.0},
			modifyClone: func(clone *Item) {
				clone.Name = "Item2"
				clone.Price = 5.0
			},
			expectedOriginal: &Item{Name: "Item1", Price: 10.0},
			expectedClone:    &Item{Name: "Item2", Price: 5.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clone := tt.original.Clone().(*Item)
			tt.modifyClone(clone)
			assert.Equal(t, tt.expectedOriginal, tt.original)
			assert.Equal(t, tt.expectedClone, clone)
		})
	}
}
