package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrototype(t *testing.T) {

	shoes := NewShoes(40, "black")
	otherShoes := shoes.Clone()

	assert.Equal(t, shoes.GetSize(), otherShoes.GetSize())
	assert.Equal(t, shoes.GetColor(), otherShoes.GetColor())
}
