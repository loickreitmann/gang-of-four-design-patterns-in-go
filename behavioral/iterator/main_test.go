package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionIterator(t *testing.T) {
	collection := &Collection{
		items: []interface{}{"A", "B", "C"},
	}
	iterator := collection.Iterator()

	assert.True(t, iterator.HasNext(), "Iterator should have next element")
	assert.Equal(t, "A", iterator.Next(), "First element should be 'A'")

	assert.True(t, iterator.HasNext(), "Iterator should have next element")
	assert.Equal(t, "B", iterator.Next(), "Second element should be 'B'")

	assert.True(t, iterator.HasNext(), "Iterator should have next element")
	assert.Equal(t, "C", iterator.Next(), "Third element should be 'C'")

	assert.False(t, iterator.HasNext(), "Iterator should not have next element")
	assert.Nil(t, iterator.Next(), "Next element should be nil when no more elements")
}

func TestEmptyCollectionIterator(t *testing.T) {
	collection := &Collection{
		items: []interface{}{},
	}
	iterator := collection.Iterator()

	assert.False(t, iterator.HasNext(), "Iterator should not have next element for empty collection")
	assert.Nil(t, iterator.Next(), "Next element should be nil for empty collection")
}
