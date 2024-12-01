package _2024

import (
	"testing"

	"github.com/pauldolden/advent-go/config"
	"github.com/stretchr/testify/assert"
)

func TestOneOne(t *testing.T) {
	o := config.Options{
		Test:        true,
		SplitInputs: false,
	}
	res := OneOne(o)

	assert.Equal(t, 11, res, "Should be equal")
}

func TestOneTwo(t *testing.T) {
	o := config.Options{
		Test:        true,
		SplitInputs: false,
	}
	res := OneTwo(o)

	assert.Equal(t, 31, res, "Should be equal")
}
