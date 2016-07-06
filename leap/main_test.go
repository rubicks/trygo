package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeap(t *testing.T) {
	if isleap(1700) {
		t.Fatalf("isleap(1700) == %v", true)
	}
}
