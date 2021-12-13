package year_2018

import (
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
		ops := map[util.IntPair]byte{}
		for y := range forest {
			for x, b := range forest[y] {
				switch b {
				case ground:
					{
						if util.CountInByteSlice(tree, x, y, forest) >= 3 {
							ops[util.IntPair{x, y}] = tree
						}
					}
				case tree:
					{
						if util.CountInByteSlice(lumberyard, x, y, forest) >= 3 {
							ops[util.IntPair{x, y}] = lumberyard
						}
					}
				case lumberyard:
					{
						if util.CountInByteSlice(tree, x, y, forest) < 1 || util.CountInByteSlice(lumberyard, x, y, forest) < 1 {
							ops[util.IntPair{x, y}] = ground
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
		ops := map[util.IntPair]byte{}
		for y := range forest {
			for x, b := range forest[y] {
				switch b {
				case ground:
					{
						if util.CountInByteSlice(tree, x, y, forest) >= 3 {
							ops[util.IntPair{x, y}] = tree
						}
					}
				case tree:
					{
						if util.CountInByteSlice(lumberyard, x, y, forest) >= 3 {
							ops[util.IntPair{x, y}] = lumberyard
						}
					}
				case lumberyard:
					{
						if util.CountInByteSlice(tree, x, y, forest) < 1 || util.CountInByteSlice(lumberyard, x, y, forest) < 1 {
							ops[util.IntPair{x, y}] = ground
						}
					}
				}
			}
		}
		for k, v := range ops {
			forest[k.Y][k.X] = v
		}
		hash := util.ByteSliceHash(forest)
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
