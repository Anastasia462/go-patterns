package factory_method

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	words := []string{"hi", "meow", "i'm stupid"}

	factory := NewCreator()
	creatures := []Creature{
		factory.CreateCreature("human"),
		factory.CreateCreature("cat"),
		factory.CreateCreature("nastya"),
	}

	for i, creature := range creatures {
		assert.Equal(t, creature.SayHi(), words[i])
	}

}