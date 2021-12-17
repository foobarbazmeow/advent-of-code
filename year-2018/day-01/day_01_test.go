package day_01

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func Test01Part1(t *testing.T) {
	xs := readInts("day_01.in")
	assert.Equal(t, 435, sumInts(xs))
}

func Test01Part2(t *testing.T) {
	xs := readInts("day_01.in")
	current := 0
	m := map[int]int{current: 1}
	for i := 0; ; i = (i + 1) % len(xs) {
		current += xs[i]
		m[current] += 1
		if c := m[current]; c == 2 {
			break
		}
	}
	assert.Equal(t, 245, current)
}

func sumInts(xs []int) int {
	result := 0
	for _, v := range xs {
		result += v
	}
	return result
}

func readInts(filepath string) []int {
	xs := make([]int, 0)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		xs = append(xs, num)
	}
	return xs
}
