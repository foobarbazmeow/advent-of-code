package day_17

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, _ := os.Open("day_17.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	a := area{}
	fmt.Sscanf(string(bytes), "target area: x=%d..%d, y=%d..%d", &a.x1, &a.x2, &a.y1, &a.y2)

	max := 0
	for y := -1000; y < 1000; y++ {
		for x := -1000; x < 1000; x++ {
			if ok, maxY := fire(x, y, a); ok {
				if maxY > max {
					max = maxY
				}
			}
		}
	}
	assert.Equal(t, 6903, max)
}

func TestPart2(t *testing.T) {
	file, _ := os.Open("day_17.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	a := area{}
	fmt.Sscanf(string(bytes), "target area: x=%d..%d, y=%d..%d", &a.x1, &a.x2, &a.y1, &a.y2)

	m := map[pair]struct{}{}
	for y := -1000; y < 1000; y++ {
		for x := -1000; x < 1000; x++ {
			if ok, _ := fire(x, y, a); ok {
				m[pair{x, y}] = struct{}{}
			}
		}
	}

	assert.Equal(t, 2351, len(m))
}

func fire(xVelocity, yVelocity int, a area) (bool, int) {
	maxY := 0
	x, y := 0, 0
	xV, yV := xVelocity, yVelocity
	for {
		if a.contains(x, y) {
			return true, maxY
		}

		x += xV
		y += yV
		if y > maxY {
			maxY = y
		}

		if xV > 0 {
			xV -= 1
		} else if xV < 0 {
			xV += 1
		}

		yV -= 1

		if a.miss(x, y) {
			return false, maxY
		}
	}
}

type area struct {
	x1, x2, y1, y2 int
}

func (a area) contains(x, y int) bool {
	return x >= a.x1 && x <= a.x2 && y >= a.y1 && y <= a.y2
}

func (a area) miss(x, y int) bool {
	return y < a.y1 || x > a.x2
}

type pair struct {
	x, y int
}
