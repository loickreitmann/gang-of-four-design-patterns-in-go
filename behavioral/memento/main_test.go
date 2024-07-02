package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOriginatorSetAndGetState(t *testing.T) {
	originator := &Originator{}
	originator.SetState("Shock")

	state := originator.GetState()
	assert.Equal(t, "Shock", state, "The state should be 'Shock'")
}

func TestOriginatorSaveAndRestoreState(t *testing.T) {
	originator := &Originator{}
	originator.SetState("Shock")

	memento := originator.SaveState()
	assert.Equal(t, "Shock", memento.state, "The saved state should be 'Shock'")

	originator.SetState("Denial")
	originator.RestoreState(memento)
	assert.Equal(t, "Shock", originator.GetState(), "The restored state should be 'Shock'")
}

func TestCaretakerAddAndRetrieveMementos(t *testing.T) {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("Shock")
	memento1 := originator.SaveState()
	caretaker.AddMemento(memento1)

	originator.SetState("Denial")
	memento2 := originator.SaveState()
	caretaker.AddMemento(memento2)

	retrievedMemento1 := caretaker.GetMemento(0)
	assert.Equal(t, "Shock", retrievedMemento1.state, "The first retrieved state should be 'Shock'")

	retrievedMemento2 := caretaker.GetMemento(1)
	assert.Equal(t, "Denial", retrievedMemento2.state, "The second retrieved state should be 'Denial'")
}

func TestMainScenario(t *testing.T) {
	originator := &Originator{}
	caretaker := &Caretaker{}

	originator.SetState("Shock")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Denial")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Anger")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Bargaining")
	caretaker.AddMemento(originator.SaveState())

	originator.SetState("Depression")
	assert.Equal(t, "Depression", originator.GetState(), "Current state should be 'Depression'")

	originator.RestoreState(caretaker.GetMemento(0))
	assert.Equal(t, "Shock", originator.GetState(), "Restored state should be 'Shock'")

	originator.RestoreState(caretaker.GetMemento(1))
	assert.Equal(t, "Denial", originator.GetState(), "Restored state should be 'Denial'")

	originator.RestoreState(caretaker.GetMemento(2))
	assert.Equal(t, "Anger", originator.GetState(), "Restored state should be 'Anger'")

	originator.RestoreState(caretaker.GetMemento(3))
	assert.Equal(t, "Bargaining", originator.GetState(), "Restored state should be 'Bargaining'")
}
