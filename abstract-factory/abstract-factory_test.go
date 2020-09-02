package abstract_factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	material = "cotton"
	size     = 46
)

func TestAbstractFactory(t *testing.T) {

	nikeFactory := NewNikeFactory()
	adidasFactory := NewAdidasFactory()

	nikeShirt := nikeFactory.CreateShirt(material, size)
	nikeShorts := nikeFactory.CreateShorts(material, size)

	adidasShirt := adidasFactory.CreateShorts(material, size)
	adidasShorts := adidasFactory.CreateShorts(material, size)

	assert.NotEqual(t, nikeShirt.GetMaterial(), adidasShirt.GetMaterial())
	assert.NotEqual(t, nikeShirt.GetSize(), adidasShirt.GetSize())
	assert.NotEqual(t, nikeShorts.GetMaterial(), adidasShorts.GetMaterial())
	assert.NotEqual(t, nikeShorts.GetSize(), adidasShorts.GetSize())
}
