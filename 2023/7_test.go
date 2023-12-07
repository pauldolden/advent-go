package _2023

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pauldolden/advent-go/config"
)

func TestSevenOne(t *testing.T) {
	o := config.NewTestOptions()
	res := SevenOne(o)

	assert.Equal(t, 6440, res, "Should be equal")
}

func TestSevenTwo(t *testing.T) {
	o := config.NewTestOptions()
	res := SevenTwo(o)

	assert.Equal(t, 71503, res, "Should be equal")
}
