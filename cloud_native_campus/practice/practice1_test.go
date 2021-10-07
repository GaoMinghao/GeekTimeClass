package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	input := []string{"I", "am", "stupid", "and", "weak"}
	t.Log("start testing")
	result := replace(input)
	assert.Equal(t, result, []string{"I", "am", "smart", "and", "strong"})
}
