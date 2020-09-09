package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleton(t *testing.T) {

	instance1 := GetInstance()
	instance2 := GetInstance()

	assert.Equal(t, instance1, instance2)
}
