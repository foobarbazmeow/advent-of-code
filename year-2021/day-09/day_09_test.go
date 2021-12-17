package day_09

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test09Part1(t *testing.T) {
	file, _ := os.Open("day_09.in")
	bytes, _ := io.ReadAll(file)
	lines2 := strings.Split(string(bytes), "\n")
	assert.Equal(t, 591, calculateRiskLevel(lines2))
}

func Test09Part2(t *testing.T) {
	file, _ := os.Open("day_09.in")
	bytes, _ := io.ReadAll(file)
	lines2 := strings.Split(string(bytes), "\n")
	basins2 := calculateBasins(lines2)
	// 6874560
	// 1751120
	// 1113424
	assert.Equal(t, 1113424, basins2[0]*basins2[1]*basins2[2])
}

func peek(xs []string, x, y int) int {
	if x >= 0 && x <= len(xs[0])-1 {
		if y >= 0 && y <= len(xs)-1 {
			num, _ := strconv.Atoi(string(xs[y][x]))
			return num
		}
	}
	return -1
}

func calculateRiskLevel(xs []string) int {
	result := 0
	for y := 0; y < len(xs); y++ {
		for x := 0; x < len(xs[0]); x++ {
			current, _ := strconv.Atoi(string(xs[y][x]))
			if top := peek(xs, x, y-1); top != -1 && top <= current {
				continue
			}
			if right := peek(xs, x+1, y); right != -1 && right <= current {
				continue
			}
			if bottom := peek(xs, x, y+1); bottom != -1 && bottom <= current {
				continue
			}
			if left := peek(xs, x-1, y); left != -1 && left <= current {
				continue
			}
			result += current + 1
		}
	}
	return result
}

func calculateBasins(xs []string) []int {
	m := map[int][]pair{}

	for y := 0; y < len(xs); y++ {
		for x := 0; x < len(xs[0]); x++ {
			arr := traverse(xs, x, y, map[pair]struct{}{})
			if len(arr) > 0 {
				h := hash(arr)
				if _, ok := m[h]; !ok {
					m[h] = arr
				}
			}
		}
	}

	basins := make([]int, 0)
	for _, v := range m {
		basins = append(basins, len(v))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	return basins
}

func hash(xs []pair) int {
	result := 0
	for _, p := range xs {
		result += p.x + p.y
	}
	return result
}

func traverse(xs []string, x, y int, scanned map[pair]struct{}) []pair {
	m := map[pair]struct{}{{x, y}: {}}
	result := make([]pair, 0)

	pairs := scan(xs, x, y)
	scanned[pair{x, y}] = struct{}{}
	if len(pairs) == 0 {
		return result
	} else {
		for _, p := range pairs {
			m[p] = struct{}{}
		}

		for _, pair := range around(x, y) {
			if _, ok := scanned[pair]; !ok {
				for _, p := range traverse(xs, pair.x, pair.y, scanned) {
					m[p] = struct{}{}
				}
			}
		}
	}

	if len(m) == 0 {
		return result
	}

	for k, _ := range m {
		result = append(result, k)
	}
	return result
}

func scan(xs []string, x, y int) []pair {
	result := make([]pair, 0)

	var current int
	if current = peek(xs, x, y); current == -1 || current == 9 {
		return result
	}

	if top := peek(xs, x, y-1); top != -1 && top != 9 && top > current {
		result = append(result, pair{x, y - 1})
	}
	if right := peek(xs, x+1, y); right != -1 && right != 9 && right > current {
		result = append(result, pair{x + 1, y})
	}
	if bottom := peek(xs, x, y+1); bottom != -1 && bottom != 9 && bottom > current {
		result = append(result, pair{x, y + 1})
	}
	if left := peek(xs, x-1, y); left != -1 && left != 9 && left > current {
		result = append(result, pair{x - 1, y})
	}

	return result
}

func around(x, y int) []pair {
	return []pair{
		{x - 1, y},
		{x + 1, y},
		{x, y - 1},
		{x, y + 1},
	}
}

type pair struct {
	x, y int
}
