package _2023

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pauldolden/advent-go/config"
)

func TestFourOne(t *testing.T) {
	o := config.NewTestOptions()
	res := FourOne(o)

	assert.Equal(t, 13, res, "Should be equal")
}

func TestFourTwo(t *testing.T) {
	o := config.NewTestOptions()
	res := FourTwo(o)

	assert.Equal(t, 30, res, "Should be equal")
}
