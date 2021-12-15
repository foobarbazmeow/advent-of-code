package year_2021

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

func Test07Part1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_07.in")
	parts := strings.Split(string(bytes), ",")
	xs := make([]int, len(parts))
	for i, c := range parts {
		num, _ := strconv.Atoi(c)
		xs[i] = num
	}

	lowest := math.MaxInt32
	for cand := min(xs); cand < max(xs); cand++ {
		val := 0
		for j := 0; j < len(xs); j++ {
			if cand == xs[j] {
				continue
			}

			if cand > xs[j] {
				val += cand - xs[j]
			} else {
				val += xs[j] - cand
			}
		}

		if val < lowest {
			lowest = val
		}
	}

	t.Log(lowest)
}

func Test07Part2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_07.in")
	parts := strings.Split(string(bytes), ",")
	xs := make([]int, len(parts))
	for i, c := range parts {
		num, _ := strconv.Atoi(c)
		xs[i] = num
	}

	lowest := math.MaxInt32
	for cand := min(xs); cand < max(xs); cand++ {
		val := 0
		for j := 0; j < len(xs); j++ {
			if cand == xs[j] {
				continue
			}

			if cand > xs[j] {
				val += gauss(cand - xs[j])
			} else {
				val += gauss(xs[j] - cand)
			}
		}

		if val < lowest {
			lowest = val
		}
	}

	assert.Equal(t, 98231647, lowest)
}

func gauss(n int) int {
	return ((n * n) + n) / 2
}

func min(xs []int) int {
	min := xs[0]
	for _, v := range xs {
		if v < min {
			min = v
		}
	}
	return min
}

func max(xs []int) int {
	max := xs[0]
	for _, v := range xs {
		if v > max {
			max = v
		}
	}
	return max
}
