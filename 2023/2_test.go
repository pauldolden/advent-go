package _2023

import (
	"testing"

	"github.com/pauldolden/advent-go/config"
	"github.com/stretchr/testify/assert"
)

var o = config.NewTestOptions()

func TestTwoOne(t *testing.T) {
	res := TwoOne(o)

	assert.Equal(t, 8, res, "Should be equal")
}
