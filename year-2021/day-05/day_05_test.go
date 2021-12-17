package day_05

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test05Part1(t *testing.T) {
	vectors := parse("day_05.in")
	filter := func(v vector) bool { return v.x1 == v.x2 || v.y1 == v.y2 }
	result := countPoints(vectors, filter)
	assert.Equal(t, 3990, result)
}

func Test05Part2(t *testing.T) {
	vectors := parse("day_05.in")
	filter := func(vector) bool { return true }
	result := countPoints(vectors, filter)
	assert.Equal(t, 21305, result)
}

func countPoints(vectors []vector, filter func(v vector) bool) int {
	m := map[string]int{}
	for _, v := range vectors {
		if !filter(v) {
			continue
		}
		for _, c := range v.coords() {
			m[c]++
		}
	}
	counter := 0
	for _, v := range values(m) {
		if v >= 2 {
			counter++
		}
	}
	return counter
}

type vector struct {
	x1, y1, x2, y2 int
}

func (v vector) coords() []string {
	xs := make([]string, 0)
	for x, y := v.x1, v.y1; x != v.x2 || y != v.y2; {
		xs = append(xs, fmt.Sprintf("%d,%d", x, y))
		if v.x2 > x {
			x += 1
		} else if v.x2 < x {
			x -= 1
		}
		if v.y2 > y {
			y += 1
		} else if v.y2 < y {
			y -= 1
		}
		if x == v.x2 && y == v.y2 {
			xs = append(xs, fmt.Sprintf("%d,%d", x, y))
		}
	}
	return xs
}

func parse(filepath string) []vector {
	xs := make([]vector, 0)
	for _, line := range readLines(filepath) {
		var v vector
		fmt.Sscanf(line, "%d,%d -> %d,%d", &v.x1, &v.y1, &v.x2, &v.y2)
		xs = append(xs, v)
	}
	return xs
}

func values(m map[string]int) []int {
	xs := make([]int, 0)
	for _, v := range m {
		xs = append(xs, v)
	}
	return xs
}

func readLines(filepath string) []string {
	xs := make([]string, 0)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, scanner.Text())
	}
	return xs
}
