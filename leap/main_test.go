package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeap(t *testing.T) {
	assert := assert.New(t)
	assert.False(isleap(1700))
	assert.False(isleap(1800))
	assert.False(isleap(1900))
	assert.True(isleap(2000))
}
