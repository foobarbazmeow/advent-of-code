package year_2021

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

func Test14Part1(t *testing.T) {
	lines := readLines("day_14.in")
	result := applyRules(lines[0], parseRules(lines[2:]), 10)
	assert.Equal(t, 2584, result)
}

func Test14Part2(t *testing.T) {
	lines := readLines("day_14.in")
	result := applyRules(lines[0], parseRules(lines[2:]), 40)
	assert.Equal(t, 3816397135460, result)
}

func applyRules(tmpl string, rules map[string]byte, steps int) int {
	xs := map[string]int{}
	c := map[byte]int{}

	for i := 0; i < len(tmpl)-1; i++ {
		xs[tmpl[i:i+2]]++
		c[tmpl[i]]++
		if i == len(tmpl)-2 {
			c[tmpl[i+1]]++
		}
	}

	for ; steps > 0; steps-- {
		ops := map[string]int{}
		for k, v := range xs {
			r := rules[k]
			c[r] += v
			ops[string([]byte{k[0], r})] += v
			ops[string([]byte{r, k[1]})] += v
		}
		xs = ops
	}

	min, max := math.MaxInt, math.MinInt
	for _, v := range c {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return max - min
}

func parseRules(xs []string) map[string]byte {
	m := map[string]byte{}
	for _, s := range xs {
		parts := strings.Split(s, " -> ")
		m[parts[0]] = []byte(parts[1])[0]
	}
	return m
}
