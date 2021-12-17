package day_02

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test02Part1(t *testing.T) {
	hor, depth := 0, 0
	for _, str := range readLines("day_02.in") {
		parts := strings.Split(str, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			hor += val
		case "up":
			depth -= val
		case "down":
			depth += val
		}
	}
	assert.Equal(t, 1604850, hor*depth)
}

func Test02Part2(t *testing.T) {
	hor, depth, aim := 0, 0, 0
	for _, str := range readLines("day_02.in") {
		parts := strings.Split(str, " ")
		val, _ := strconv.Atoi(parts[1])
		switch parts[0] {
		case "forward":
			{
				hor += val
				depth += aim * val
			}
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}
	assert.Equal(t, 1685186100, hor*depth)
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
