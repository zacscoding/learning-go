package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicAssertions(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "msgAndArgs...")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil
	var obj interface{}
	assert.Nil(t, obj)

	obj = "aa"
	if assert.NotNil(t, obj) {
		assert.Equal(t, "aa", obj)
	}
}

func TestAssertionsWithoutT(t *testing.T) {
	assert := assert.New(t)

	// without t
	assert.Equal( 123, 123, "msgAndArgs...")
}


