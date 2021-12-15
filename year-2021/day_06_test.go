package year_2021

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func Test06Part1(t *testing.T) {
	arr := func(str string) []int {
		parts := strings.Split(str, ",")
		xs := make([]int, len(parts))
		for i, s := range parts {
			num, _ := strconv.Atoi(s)
			xs[i] = num
		}
		return xs
	}
	bytes, _ := ioutil.ReadFile("day_06.in")
	xs := arr(string(bytes))
	assert.Equal(t, 383160, simulate(xs, 80))
}

func Test06Part2(t *testing.T) {
	arr := func(str string) []int {
		parts := strings.Split(str, ",")
		xs := make([]int, len(parts))
		for i, s := range parts {
			num, _ := strconv.Atoi(s)
			xs[i] = num
		}
		return xs
	}
	bytes, _ := ioutil.ReadFile("day_06.in")
	xs := arr(string(bytes))

	days := make([]int, 9)
	for i := 0; i < len(xs); i++ {
		days[xs[i]] += 1
	}

	for i := 0; i < 256; i++ {
		d0 := days[0]
		for j := 0; j < len(days)-1; j++ {
			days[j] = days[j+1]
		}
		days[6] += d0
		days[8] = d0
	}

	summ := 0
	for i := 0; i < len(days); i++ {
		summ += days[i]
	}

	t.Log(summ)
}

func simulate(xs []int, days int) int {
	left := days
	for left > 0 {
		add := 0
		for i := 0; i < len(xs); i++ {
			if xs[i] == 0 {
				add++
				xs[i] = 6
			} else {
				xs[i]--
			}
		}

		if add > 0 {
			for i := 0; i < add; i++ {
				xs = append(xs, 8)
			}
		}

		left--
	}
	return len(xs)
}
