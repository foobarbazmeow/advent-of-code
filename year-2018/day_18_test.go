package year_2018

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/antigravity/advent-of-code/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	ground     = '.'
	tree       = '|'
	lumberyard = '#'
)

func Test18Part1(t *testing.T) {
	assert.Equal(t, 558960, countResources(util.ReadBytes("day_18.in"), 10))
}

func Test18Part2(t *testing.T) {
	minutes := 1000000000
	left, segment := countMinutes(util.ReadBytes("day_18.in"), minutes)
	count := (minutes - left) / segment
	right := minutes - left - (segment * count)

	l := countResources(util.ReadBytes("day_18.in"), left)
	r := countResources(util.ReadBytes("day_18.in"), left+right) - l

	assert.Equal(t, 207900, l+r)
}

func countResources(forest [][]byte, minutes int) int {
	for minute := 0; minute < minutes; minute++ {
		ops := map[intPair]byte{}
		for y := range forest {
			for x, b := range forest[y] {
				switch b {
				case ground:
					{
						if countInByteSlice(tree, x, y, forest) >= 3 {
							ops[intPair{x, y}] = tree
						}
					}
				case tree:
					{
						if countInByteSlice(lumberyard, x, y, forest) >= 3 {
							ops[intPair{x, y}] = lumberyard
						}
					}
				case lumberyard:
					{
						if countInByteSlice(tree, x, y, forest) < 1 || countInByteSlice(lumberyard, x, y, forest) < 1 {
							ops[intPair{x, y}] = ground
						}
					}
				}
			}
		}
		for k, v := range ops {
			forest[k.Y][k.X] = v
		}
	}
	return resourceValue(forest)
}

func countMinutes(forest [][]byte, minutes int) (int, int) {
	cache := map[string]int{}
	for minute := 1; minute <= minutes; minute++ {
		ops := map[intPair]byte{}
		for y := range forest {
			for x, b := range forest[y] {
				switch b {
				case ground:
					{
						if countInByteSlice(tree, x, y, forest) >= 3 {
							ops[intPair{x, y}] = tree
						}
					}
				case tree:
					{
						if countInByteSlice(lumberyard, x, y, forest) >= 3 {
							ops[intPair{x, y}] = lumberyard
						}
					}
				case lumberyard:
					{
						if countInByteSlice(tree, x, y, forest) < 1 || countInByteSlice(lumberyard, x, y, forest) < 1 {
							ops[intPair{x, y}] = ground
						}
					}
				}
			}
		}
		for k, v := range ops {
			forest[k.Y][k.X] = v
		}
		hash := byteSliceHash(forest)
		if v, ok := cache[hash]; ok {
			return v, minute - v
		} else {
			cache[hash] = minute
		}
	}
	return -1, -1
}

func resourceValue(forest [][]byte) int {
	t, l := 0, 0
	for y := range forest {
		for _, b := range forest[y] {
			switch b {
			case tree:
				t += 1
			case lumberyard:
				l += 1
			}
		}
	}
	return t * l
}

func byteSliceHasElement(x, y int, xs [][]byte) bool {
	if y >= 0 && y <= len(xs)-1 {
		if x >= 0 && x <= len(xs[0])-1 {
			return true
		}
	}
	return false
}

type intPair struct {
	X, Y int
}

func adjacentWithDiagonals(x, y int) []intPair {
	return []intPair{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},

		{x - 1, y},
		{x + 1, y},

		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

func byteSliceHash(xs [][]byte) string {
	h := sha1.New()
	for _, row := range xs {
		h.Write(row)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func countInByteSlice(val byte, x, y int, xs [][]byte) int {
	result := 0
	for _, p := range adjacentWithDiagonals(x, y) {
		if byteSliceHasElement(p.X, p.Y, xs) {
			if xs[p.Y][p.X] == val {
				result += 1
			}
		}
	}
	return result
}
