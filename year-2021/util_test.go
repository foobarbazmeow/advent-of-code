package year_2021

import (
	"strconv"
	"strings"
)

type pair struct {
	x, y int
}

func readIntSlice(in string) [][]int {
	rows := strings.Split(in, "\n")
	result := make([][]int, len(rows))
	for y, str := range rows {
		row := make([]int, len(str))
		for x, r := range str {
			num, _ := strconv.Atoi(string(r))
			row[x] = num
		}
		result[y] = row
	}
	return result
}

func sliceHasElement(xs [][]int, x, y int) bool {
	if y >= 0 && y <= len(xs)-1 {
		if x >= 0 && x <= len(xs[0])-1 {
			return true
		}
	}
	return false
}

func neighbours(x, y int) []pair {
	return []pair{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},

		{x - 1, y},
		{x + 1, y},

		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
}

type intMatrix [][]int

func (m intMatrix) String() string {
	sb := strings.Builder{}
	for _, row := range m {
		for _, n := range row {
			sb.WriteString(strconv.Itoa(n))
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (m intMatrix) AllIs(num int) bool {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] != num {
				return false
			}
		}
	}
	return true
}
