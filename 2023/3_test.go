package _2023

import (
	"testing"

	"github.com/pauldolden/advent-go/config"
	"github.com/stretchr/testify/assert"
)

func TestThreeOne(t *testing.T) {
	o := config.NewTestOptions()
	res := ThreeOne(o)

	assert.Equal(t, 4361, res, "Should be equal")
}

func TestThreeTwo(t *testing.T) {
	o := config.NewTestOptions()
	res := ThreeTwo(o)

	assert.Equal(t, 467835, res, "Should be equal")
}
