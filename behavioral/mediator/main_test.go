package main

import (
	"testing"

	"gang_of_four_design_patterns_in_go/helper"

	"github.com/stretchr/testify/assert"
)

func TestConcreteColleagueSend(t *testing.T) {
	mediator := &ConcreteMediator{}
	colleague1 := &ConcreteColleague{name: "Obi Wan"}
	colleague2 := &ConcreteColleague{name: "Anakin"}

	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)

	mediator.SetColleague1(colleague1)
	mediator.SetColleague2(colleague2)

	output := helper.CaptureOutput(func() {
		colleague1.Send("Hello, there.")
		colleague2.Send("This is where the fun begins!")
	})

	expectedOutput := "Obi Wan sending message: Hello, there.\nAnakin received message: Hello, there.\nAnakin sending message: This is where the fun begins!\nObi Wan received message: This is where the fun begins!\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestConcreteMediatorSend(t *testing.T) {
	mediator := &ConcreteMediator{}
	colleague1 := &ConcreteColleague{name: "Obi Wan"}
	colleague2 := &ConcreteColleague{name: "Anakin"}

	colleague1.SetMediator(mediator)
	colleague2.SetMediator(mediator)

	mediator.SetColleague1(colleague1)
	mediator.SetColleague2(colleague2)

	output := helper.CaptureOutput(func() {
		mediator.Send("Message from Obi Wan", colleague1)
		mediator.Send("Message from Anakin", colleague2)
	})

	expectedOutput := "Anakin received message: Message from Obi Wan\nObi Wan received message: Message from Anakin\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}

func TestColleagueReceive(t *testing.T) {
	colleague := &ConcreteColleague{name: "Yoda"}

	output := helper.CaptureOutput(func() {
		colleague.Receive("May the Force be with you.")
	})

	expectedOutput := "Yoda received message: May the Force be with you.\n"
	assert.Equal(t, expectedOutput, output, "Output should match the expected format")
}
