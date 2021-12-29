package day_20

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	code, image := parse(read("day_20.in"), 2)
	for i := 0; i < 2; i++ {
		if !code[0] {
			image = enhance(code, image, false)
		} else {
			image = enhance(code, image, i%2 == 1)
		}
	}
	assert.Equal(t, 5316, count(image))
}

func TestPart2(t *testing.T) {
	code, image := parse(read("day_20.in"), 50)
	for i := 0; i < 50; i++ {
		if !code[0] {
			image = enhance(code, image, false)
		} else {
			image = enhance(code, image, i%2 == 1)
		}
	}
	assert.Equal(t, 16728, count(image))
}

func read(filepath string) string {
	file, _ := os.Open(filepath)
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	return string(bytes)
}

func parse(in string, ext int) ([]bool, [][]bool) {
	parts := strings.Split(in, "\n\n")

	lines := strings.Split(parts[1], "\n")
	height := len(lines) + ext*2
	width := len(lines[0]) + ext*2

	image := empty(height, width)
	for y := 0; y < len(lines); y++ {
		image[y+ext] = append(append(repeat(ext, false), toBin(lines[y])...), repeat(ext, false)...)
	}

	return toBin(parts[0]), image
}

func repeat(n int, b bool) []bool {
	r := make([]bool, n)
	for i := 0; i < n; i++ {
		r[i] = b
	}
	return r
}

func empty(height, width int) [][]bool {
	r := make([][]bool, height)
	for y := 0; y < len(r); y++ {
		r[y] = make([]bool, width)
	}
	return r
}

func toBin(xs string) []bool {
	r := make([]bool, len(xs))
	for i, ch := range xs {
		r[i] = ch == '#'
	}
	return r
}

func enhance(code []bool, image [][]bool, infValue bool) [][]bool {
	height := len(image)
	width := len(image[0])
	r := empty(height, width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			n := neighbours9(x, y)
			index := idx(n, image, infValue)
			r[y][x] = code[index]
		}
	}
	return r
}

type pair struct{ x, y int }

func neighbours9(x, y int) []pair {
	return []pair{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},

		{x - 1, y},
		{x, y},
		{x + 1, y},

		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

func idx(xs []pair, image [][]bool, infValue bool) int {
	str := ""
	height := len(image)
	width := len(image[0])
	for _, p := range xs {
		if in(p, height, width) {
			if image[p.y][p.x] {
				str += "1"
			} else {
				str += "0"
			}
		} else {
			if infValue {
				str += "1"
			} else {
				str += "0"
			}
		}
	}
	num, _ := strconv.ParseInt(str, 2, 32)
	return int(num)
}

func in(p pair, height, width int) bool {
	if p.y >= 0 && p.y < height {
		if p.x >= 0 && p.x < width {
			return true
		}
	}
	return false
}

func count(xs [][]bool) int {
	r := 0
	for y := 0; y < len(xs); y++ {
		for x := 0; x < len(xs[y]); x++ {
			if xs[y][x] {
				r += 1
			}
		}
	}
	return r
}
