package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumInts(t *testing.T) {
	xs := []int{1, 2, 3}
	assert.Equal(t, 6, SumInts(xs))
}
