package _2023

import (
	"testing"

	"github.com/pauldolden/advent-go/config"
	"github.com/stretchr/testify/assert"
)

func TestOneOne(t *testing.T) {
	o := config.Options{
		Test:        true,
		SplitInputs: true,
		TestPart:    1,
	}
	res := OneOne(o)

	assert.Equal(t, 142, res, "Should be equal")
}

func TestOneTwo(t *testing.T) {
	o := config.Options{
		Test:        true,
		SplitInputs: true,
		TestPart:    2,
	}
	res := OneTwo(o)

	assert.Equal(t, 281, res, "Should be equal")
}
