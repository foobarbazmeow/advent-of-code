package day_04

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

const (
	BoardWidth  = 5
	BoardHeight = 5
)

func Test04Part1(t *testing.T) {
	input, boards := parseInput("day_04.in")
	scores := play(input, boards)
	assert.Equal(t, 16674, scores[0])
}

func Test04Part2(t *testing.T) {
	input, boards := parseInput("day_04.in")
	scores := play(input, boards)
	assert.Equal(t, 7075, scores[len(scores)-1])
}

func parseInput(filepath string) ([]int, [][][]int) {
	bytes, _ := ioutil.ReadFile(filepath)
	parts := strings.Split(string(bytes), "\n\n")

	in := mapToInt(strings.Split(parts[0], ","))

	boards := make([][][]int, len(parts)-1)
	for i := 1; i < len(parts); i++ {
		str := strings.Split(strings.ReplaceAll(strings.Trim(strings.Join(strings.Split(parts[i], "\n"), " "), " "), "  ", " "), " ")
		boards[i-1] = chunk(mapToInt(str), 5)
	}

	return in, boards
}

func play(in []int, boards [][][]int) []int {
	scores := make([]int, 0)
	for _, num := range in {
		for _, b := range boards {
			if hasWin(b) {
				continue
			}
			move(b, num)
			if hasWin(b) {
				scores = append(scores, score(b, num))
			}
		}
	}
	return scores
}

func mapToInt(in []string) []int {
	xs := make([]int, len(in))
	for i := 0; i < len(in); i++ {
		num, _ := strconv.Atoi(in[i])
		xs[i] = num
	}
	return xs
}

func hasWin(xs [][]int) bool {
	for i := 0; i < BoardHeight; i++ {
		// row
		if allNegatives(xs[i]) {
			return true
		}

		// column
		column := make([]int, 5)
		for j := 0; j < BoardHeight; j++ {
			column[j] = xs[j][i]
		}
		if allNegatives(column) {
			return true
		}
	}
	return false
}

func move(xs [][]int, num int) {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			if xs[i][j] == num {
				xs[i][j] = -1
			}
		}
	}
}

func score(xs [][]int, num int) int {
	sum := 0
	for i := 0; i < BoardHeight; i++ {
		for j := 0; j < BoardWidth; j++ {
			if xs[i][j] > 0 {
				sum += xs[i][j]
			}
		}
	}
	return sum * num
}

func allNegatives(in []int) bool {
	for _, b := range in {
		if b >= 0 {
			return false
		}
	}
	return true
}

func chunk(xs []int, size int) [][]int {
	chunks := len(xs) / size
	if len(xs)%size != 0 {
		chunks += 1
	}
	lastChunk := chunks - 1
	result := make([][]int, chunks)
	for i, from := 0, 0; i < len(result); i, from = i+1, from+size {
		if i == lastChunk {
			result[i] = xs[from:]
		} else {
			result[i] = xs[from : from+size]
		}
	}
	return result
}
