package day_11

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test11Part1(t *testing.T) {
	file2, _ := os.Open("day_11.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	assert.Equal(t, 1757, countFlashes(readIntSlice(string(bytes2))))
}

func Test11Part2(t *testing.T) {
	file2, _ := os.Open("day_11.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	assert.Equal(t, 422, allOctopusFlashAt(readIntSlice(string(bytes2))))
}

func countFlashes(xs [][]int) int {
	flashes := 0
	iteration := 100
	for iteration > 0 {
		flashed := map[pair]struct{}{}
		for y := 0; y < len(xs); y++ {
			for x := 0; x < len(xs[0]); x++ {
				flashes += countFlashesIteration(x, y, xs, flashed)
			}
		}
		iteration -= 1
	}
	return flashes
}

func allOctopusFlashAt(xs [][]int) int {
	iteration := 0
	for iteration < 1000 {
		flashed := map[pair]struct{}{}
		for y := 0; y < len(xs); y++ {
			for x := 0; x < len(xs[0]); x++ {
				countFlashesIteration(x, y, xs, flashed)
			}
		}
		iteration += 1
		if allIs(xs, 0) {
			return iteration
		}
	}
	return -1
}

func countFlashesIteration(x, y int, xs [][]int, flashed map[pair]struct{}) int {
	flashes := 0
	if sliceHasElement(xs, x, y) {
		p := pair{x, y}
		if _, ok := flashed[p]; ok {
			return flashes
		} else {
			if xs[y][x] == 9 {
				flashed[p] = struct{}{}
				xs[y][x] = 0
				flashes += 1
				for _, n := range neighbours(x, y) {
					flashes += countFlashesIteration(n.x, n.y, xs, flashed)
				}
			} else {
				xs[y][x] += 1
			}
		}
	}
	return flashes
}

func allIs(xs [][]int, num int) bool {
	for y := 0; y < len(xs); y++ {
		for x := 0; x < len(xs[0]); x++ {
			if xs[y][x] != num {
				return false
			}
		}
	}
	return true
}

func readIntSlice(in string) [][]int {
	rows := strings.Split(in, "\n")
	result := make([][]int, len(rows))
	for y, str := range rows {
		row := make([]int, len(str))
		for x, r := range str {
			num, _ := strconv.Atoi(string(r))
			row[x] = num
		}
		result[y] = row
	}
	return result
}

func sliceHasElement(xs [][]int, x, y int) bool {
	if y >= 0 && y <= len(xs)-1 {
		if x >= 0 && x <= len(xs[0])-1 {
			return true
		}
	}
	return false
}

func neighbours(x, y int) []pair {
	return []pair{
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

type pair struct {
	x, y int
}
