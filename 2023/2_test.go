package _2023

import (
	"testing"

	"github.com/pauldolden/advent-go/config"
	"github.com/stretchr/testify/assert"
)

func TestTwoOne(t *testing.T) {
	o := config.NewTestOptions()
	res := TwoOne(o)

	assert.Equal(t, 8, res, "Should be equal")
}

func TestTwoTwo(t *testing.T) {
	o := config.NewTestOptions()
	res := TwoTwo(o)

	assert.Equal(t, 2286, res, "Should be equal")
}
