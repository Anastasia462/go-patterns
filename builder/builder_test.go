package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilder(t *testing.T) {

	expect := "smart head\n" +
		"flawless body\n" +
		"strong foot\n"

	human := new(Human)

	director := Director{&ConcreteBuilder{human}}
	director.Construct()

	result := human.Born()

	assert.Equal(t, expect, result)
}
