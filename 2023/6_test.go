package _2023

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pauldolden/advent-go/config"
)

func TestSixOne(t *testing.T) {
	o := config.NewTestOptions()
	res := SixOne(o)

	assert.Equal(t, 288, res, "Should be equal")
}

func TestSixTwo(t *testing.T) {
	o := config.NewTestOptions()
	res := SixTwo(o)

	assert.Equal(t, 71503, res, "Should be equal")
}
