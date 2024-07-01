package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	assert.Equal(t, instance1, instance2, "Instances should be the same")

	instance1.Name = "Updated Singleton Instance"
	assert.Equal(t, "Updated Singleton Instance", instance2.Name, "Instance2 name should reflect the update from instance1")
}

func TestSingletonConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	var instance1, instance2 *Singleton

	wg.Add(2)

	go func() {
		defer wg.Done()
		instance1 = GetInstance()
	}()

	go func() {
		defer wg.Done()
		instance2 = GetInstance()
	}()

	wg.Wait()

	assert.Equal(t, instance1, instance2, "Instances should be the same when accessed concurrently")
}
