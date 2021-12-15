package year_2020

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test05Part1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_05.in")
	max := 0
	for _, str := range strings.Split(string(bytes), "\n") {
		row := binarySearch(str[:7], 0, 127)
		column := binarySearch(str[7:], 0, 7)
		id := row*8 + column
		if id > max {
			max = id
		}
	}
	assert.Equal(t, 864, max)
}

func Test05Part2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_05.in")
	ids := newRange(75, 864)
	for _, str := range strings.Split(string(bytes), "\n") {
		row := binarySearch(str[:7], 0, 127)
		column := binarySearch(str[7:], 0, 7)
		id := row*8 + column
		delete(ids, id)
	}
	assert.Equal(t, 1, len(ids))
	assert.Contains(t, ids, 739)
}

func binarySearch(str string, l, r int) int {
	if len(str) == 0 {
		return l
	}

	m := (l + r) / 2
	if str[0] == 'F' || str[0] == 'L' {
		return binarySearch(str[1:], l, m-1)
	}
	return binarySearch(str[1:], m+1, r)
}

func newRange(min, max int) map[int]struct{} {
	m := map[int]struct{}{}
	for i := min; i <= max; i++ {
		m[i] = struct{}{}
	}
	return m
}
