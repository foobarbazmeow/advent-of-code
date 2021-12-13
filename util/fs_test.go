package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadLines(t *testing.T) {
	filepath := tmp()
	defer os.Remove(filepath)
	expected := []string{"1", "2", "3"}
	actual := []string{}
	for line := range ReadLines(filepath) {
		actual = append(actual, line)
	}
	assert.Equal(t, expected, actual)
}

func TestReadInts(t *testing.T) {
	filepath := tmp()
	defer os.Remove(filepath)
	expected := []int{1, 2, 3}
	actual := ReadInts(filepath)
	assert.Equal(t, expected, actual)
}

func tmp() string {
	file, _ := os.CreateTemp("", "aoc")
	defer file.Close()
	file.WriteString("1\n")
	file.WriteString("2\n")
	file.WriteString("3")
	return file.Name()
}
