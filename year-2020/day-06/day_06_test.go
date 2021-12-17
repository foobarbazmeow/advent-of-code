package day_06

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test06Part1(t *testing.T) {
	parts := read("day_06.in")
	questions := 0
	for _, lines := range parts {
		m := map[int32]int{}
		people := strings.Split(lines, "\n")

		for _, person := range people {
			for _, c := range person {
				m[c]++
			}
		}

		questions += len(m)
	}
	assert.Equal(t, 6542, questions)
}

func Test06Part2(t *testing.T) {
	parts := read("day_06.in")
	questions := 0
	for _, lines := range parts {
		m := map[int32]int{}
		people := strings.Split(lines, "\n")

		for _, person := range people {
			for _, c := range person {
				m[c]++
			}
		}

		sum := 0
		for _, v := range m {
			if v == len(people) {
				sum += 1
			}
		}

		questions += sum
	}
	assert.Equal(t, 3299, questions)
}

func read(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	bytes, _ := io.ReadAll(file)
	str := string(bytes)

	return strings.Split(str, "\n\n")
}
