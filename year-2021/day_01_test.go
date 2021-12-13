package year_2021

import (
	"github.com/antigravity/advent-of-code/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test01Part1(t *testing.T) {
	xs := util.ReadInts("day_01.in")
	result, buf := 0, xs[0]
	for i := 1; i < len(xs); i++ {
		if xs[i] > buf {
			result++
		}
		buf = xs[i]
	}
	assert.Equal(t, 1215, result)
}

func Test01Part2(t *testing.T) {
	xs := util.ReadInts("day_01.in")
	result, buf := 0, xs[0]+xs[1]+xs[2]
	for i := 1; i < len(xs)-2; i++ {
		cand := xs[i] + xs[i+1] + xs[i+2]
		if cand > buf {
			result++
		}
		buf = cand
	}
	assert.Equal(t, 1150, result)
}
