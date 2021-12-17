package day_02

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test02Part1(t *testing.T) {
	file, _ := os.Open("day_02.in")
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		var l, r int
		var ch rune
		var pass string

		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &l, &r, &ch, &pass)

		count := 0
		for _, s := range pass {
			if s == ch {
				count++
			}
		}

		if count >= l && count <= r {
			result++
		}
	}
	t.Log(result)
}

func Test02Part2(t *testing.T) {
	file, _ := os.Open("day_02.in")
	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		var l, r int
		var ch rune
		var pass string

		fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &l, &r, &ch, &pass)

		if pass[l-1] == uint8(ch) || pass[r-1] == uint8(ch) {
			if pass[l-1] != pass[r-1] {
				result++
			}
		}
	}
	t.Log(result)
}
