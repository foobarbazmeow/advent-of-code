package year_2021

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"sort"
	"strings"
	"testing"
)

var (
	scores1 = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	scores2 = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
)

func Test10Part1(t *testing.T) {
	file2, _ := os.Open("day_10.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	result2 := sumScores(lines2)
	assert.Equal(t, 278475, result2)
}

func Test10Part2(t *testing.T) {
	file2, _ := os.Open("day_10.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	result2 := middleScore(lines2)
	assert.Equal(t, 3015539998, result2)
}

func sumScores(lines []string) int {
	result := 0
	for _, line := range lines {
		result += findIllegalCharacter(line)
	}
	return result
}

func middleScore(lines []string) int {
	result := make([]int, 0)
	for _, line := range lines {
		score := findIllegalCharacterAndComplete(line)
		if score > 0 {
			result = append(result, score)
		}
	}
	sort.Ints(result)
	return result[len(result)/2]
}

func findIllegalCharacter(line string) int {
	stack := make([]rune, 0)
	for _, r := range line {
		if r == '(' {
			stack = append(stack, ')')
		} else if r == '[' {
			stack = append(stack, ']')
		} else if r == '{' {
			stack = append(stack, '}')
		} else if r == '<' {
			stack = append(stack, '>')
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if r != last {
				return scores1[r]
			}
		}
	}
	return 0
}

func findIllegalCharacterAndComplete(line string) int {
	stack := make([]rune, 0)
	for _, r := range line {
		if r == '(' {
			stack = append(stack, ')')
		} else if r == '[' {
			stack = append(stack, ']')
		} else if r == '{' {
			stack = append(stack, '}')
		} else if r == '<' {
			stack = append(stack, '>')
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if r != last {
				return 0
			}
		}
	}

	if len(stack) > 0 {
		return calculateScore(stack)
	}

	return 0
}

func calculateScore(xs []rune) int {
	result := 0
	for i := len(xs) - 1; i >= 0; i-- {
		result = result*5 + scores2[xs[i]]
	}
	return result
}
