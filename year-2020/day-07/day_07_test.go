package day_07

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	TargetBag = "shiny gold"
	EmptyBag  = "no other bags."
)

func Test07Part1(t *testing.T) {
	file2, _ := os.Open("day_07.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	result2 := countBagColors(lines2)
	assert.Equal(t, 254, result2)
}

func Test07Part2(t *testing.T) {
	file2, _ := os.Open("day_07.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	result2 := countNestedBags(lines2)
	assert.Equal(t, 6006, result2)
}

func countBagColors(lines []string) int {
	m := parseInputBags(lines)

	result := 0
	for k, _ := range m {
		if hasTargetBag(m, k) {
			result += 1
		}
	}
	return result
}

func countNestedBags(lines []string) int {
	m := parseInputBagsWithCount(lines)
	result := 0
	for k, v := range m[TargetBag] {
		result += v * countTargetBags(m, k)
	}
	return result
}

func parseInputBags(lines []string) map[string][]string {
	m := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, " bags ")
		parent := parts[0]
		nested := make([]string, 0)
		snd := strings.TrimPrefix(parts[1], "contain ")
		if snd != EmptyBag {
			for _, part := range strings.Split(snd, ",") {
				clr := strings.Join(strings.Split(strings.TrimSpace(part), " ")[1:3], " ")
				nested = append(nested, clr)
			}
		}
		m[parent] = nested
	}
	return m
}

func parseInputBagsWithCount(lines []string) map[string]map[string]int {
	m := map[string]map[string]int{}
	for _, line := range lines {
		parts := strings.Split(line, " bags ")
		parent := parts[0]
		nested := make(map[string]int)
		snd := strings.TrimPrefix(parts[1], "contain ")
		if snd != EmptyBag {
			for _, part := range strings.Split(snd, ",") {
				split := strings.Split(strings.TrimSpace(part), " ")
				count, _ := strconv.Atoi(split[0])
				clr := strings.Join(split[1:3], " ")
				nested[clr] = count
			}
		}
		m[parent] = nested
	}
	return m
}

func hasTargetBag(m map[string][]string, color string) bool {
	result := false
	if colors, ok := m[color]; ok && len(colors) > 0 {
		for _, v := range colors {
			if v == TargetBag {
				result = true
			} else {
				result = result || hasTargetBag(m, v)
			}
		}
	}
	return result
}

func countTargetBags(m map[string]map[string]int, color string) int {
	result := 1
	if colors, ok := m[color]; ok && len(colors) > 0 {
		for k, v := range colors {
			result += v * countTargetBags(m, k)
		}
	}
	return result
}
