package _2023

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pauldolden/advent-go/config"
)

func TestEightOne(t *testing.T) {
	o := config.NewTestOptions()
	res := EightOne(o)

	assert.Equal(t, 6, res, "Should be equal")
}

// func TestEightTwo(t *testing.T) {
// 	o := config.NewTestOptions()
// 	res := EightTwo(o)
//
// 	assert.Equal(t, 71503, res, "Should be equal")
// }
