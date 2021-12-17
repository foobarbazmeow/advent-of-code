package day_09

import (
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test09Part1(t *testing.T) {
	file, _ := os.Open("day_09.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	lines := strings.Split(string(bytes), "\n")

	xs := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		xs[i] = num
	}

	assert.Equal(t, 88311122, findWeakNumber(xs, 25))
}

func Test09Part2(t *testing.T) {
	file, _ := os.Open("day_09.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	lines := strings.Split(string(bytes), "\n")

	xs := make([]int, len(lines))
	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		xs[i] = num
	}

	assert.Equal(t, 13549369, findContiguousNumbers(xs, 88311122))
}

func findWeakNumber(xs []int, skip int) int {
	for n := skip; n < len(xs); n++ {
		found := false
		n3 := xs[n]
		for i := n - skip; i < n; i++ {
			for j := i + 1; j < n; j++ {
				n1 := xs[i]
				n2 := xs[j]
				if n1+n2 == n3 {
					found = true
					goto ret
				}
			}
		}
	ret:
		if !found {
			return xs[n]
		}
	}
	return -1
}

func findContiguousNumbers(xs []int, num int) int {
	h, t := 0, 0
	for i := 0; i < len(xs); i++ {
		sum := xs[i]
		for j := i + 1; j < len(xs); j++ {
			sum += xs[j]
			if sum == num {
				h, t = i, j
				goto ret
			} else if sum > num {
				break
			}
		}
	}
ret:
	min, max := math.MaxInt, math.MinInt
	for i := h; i <= t; i++ {
		if xs[i] > max {
			max = xs[i]
		}
		if xs[i] < min {
			min = xs[i]
		}
	}
	return max + min
}
