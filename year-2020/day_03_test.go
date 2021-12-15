package year_2020

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func Test03Part1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_03.in")
	forest := strings.Split(string(bytes), "\n")
	trees := findTrees(forest, 3, 1)
	assert.Equal(t, 280, trees)
}

func Test03Part2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("day_03.in")
	forest := strings.Split(string(bytes), "\n")
	trees1 := findTrees(forest, 1, 1)
	trees2 := findTrees(forest, 3, 1)
	trees3 := findTrees(forest, 5, 1)
	trees4 := findTrees(forest, 7, 1)
	trees5 := findTrees(forest, 1, 2)
	assert.Equal(t, 4355551200, trees1*trees2*trees3*trees4*trees5)
}

func findTrees(forest []string, columnInc, rowInc int) int {
	trees := 0
	for column, row := 0, 0; row < len(forest); column, row = column+columnInc, row+rowInc {
		if column >= len(forest[0]) {
			column = column % len(forest[0])
		}
		if forest[row][column] == '#' {
			trees += 1
		}
	}
	return trees
}
