package day_13

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test13Part1(t *testing.T) {
	fn := func(filepath string, expected int) {
		file1, _ := os.Open(filepath)
		defer file1.Close()
		bytes1, _ := io.ReadAll(file1)
		content1 := strings.Split(string(bytes1), "\n\n")
		dots1 := parseDots(content1[0])
		instructions1 := parseInstructions(content1[1])
		assert.Equal(t, expected, countDotsAfterFirstFold(dots1, instructions1))
	}

	fn("day_13.in", 850)
}

func Test13Part2(t *testing.T) {
	fn := func(filepath string, expected int) {
		file1, _ := os.Open(filepath)
		defer file1.Close()
		bytes1, _ := io.ReadAll(file1)
		content1 := strings.Split(string(bytes1), "\n\n")
		dots1 := parseDots(content1[0])
		instructions1 := parseInstructions(content1[1])
		assert.Equal(t, expected, countDotsAfterFolds(dots1, instructions1))
	}

	fn("day_13.in", 102) // prints AHGCPGAU
}

type instr struct {
	axis string
	val  int
}

func parseDots(in string) [][]bool {
	xs := make([][]bool, 0)
	for _, str := range strings.Split(in, "\n") {
		var x, y int
		fmt.Sscanf(str, "%d,%d", &x, &y)
		for i := len(xs) - 1; i < y; i++ {
			xs = append(xs, make([]bool, 0))
		}
		if len(xs[y])-1 < x {
			tmp := make([]bool, x+1)
			copy(tmp, xs[y])
			xs[y] = tmp
		}
		xs[y][x] = true
	}
	return xs
}

func parseInstructions(in string) []*instr {
	xs := make([]*instr, 0)
	for _, str := range strings.Split(in, "\n") {
		parts := strings.Split(strings.TrimPrefix(str, "fold along "), "=")
		val, _ := strconv.Atoi(parts[1])
		xs = append(xs, &instr{parts[0], val})
	}
	return xs
}

func countDotsAfterFolds(dots [][]bool, instructions []*instr) int {
	for _, fold := range instructions {
		switch fold.axis {
		case "x":
			{
				for y := 0; y < len(dots); y++ {
					for xr := fold.val + 1; xr < len(dots[y]); xr++ {
						xl := fold.val - (xr - fold.val)
						if xl >= 0 && xl < fold.val {
							if dots[y][xr] && !dots[y][xl] {
								dots[y][xl] = dots[y][xr]
							}
						}
					}
					if len(dots[y]) > fold.val {
						dots[y] = dots[y][:fold.val]
					}
				}
			}
		case "y":
			{
				for yb := fold.val + 1; yb < len(dots); yb++ {
					yt := fold.val - (yb - fold.val)
					if yt >= 0 && yt < fold.val {
						if len(dots[yt]) < len(dots[yb]) {
							tmp := make([]bool, len(dots[yb]))
							copy(tmp, dots[yt])
							dots[yt] = tmp
						}
						for x := 0; x < len(dots[yb]); x++ {
							if dots[yb][x] && !dots[yt][x] {
								dots[yt][x] = dots[yb][x]
							}
						}
					}
				}
				dots = dots[:fold.val]
			}
		}
	}
	fmt.Println(formatDots(dots))
	return countDots(dots)
}

func countDotsAfterFirstFold(dots [][]bool, instructions []*instr) int {
	fold := instructions[0]
	switch fold.axis {
	case "x":
		{
			for y := 0; y < len(dots); y++ {
				for xr := fold.val + 1; xr < len(dots[y]); xr++ {
					xl := fold.val - (xr - fold.val)
					if xl >= 0 && xl < fold.val {
						if dots[y][xr] && !dots[y][xl] {
							dots[y][xl] = dots[y][xr]
						}
					}
				}
				if len(dots[y]) > fold.val {
					dots[y] = dots[y][:fold.val]
				}
			}
		}
	case "y":
		{
			for yb := fold.val + 1; yb < len(dots); yb++ {
				yt := fold.val - (yb - fold.val)
				if yt >= 0 && yt < fold.val {
					if len(dots[yt]) < len(dots[yb]) {
						tmp := make([]bool, len(dots[yb]))
						copy(tmp, dots[yt])
						dots[yt] = tmp
					}
					for x := 0; x < len(dots[yb]); x++ {
						if dots[yb][x] && !dots[yt][x] {
							dots[yt][x] = dots[yb][x]
						}
					}
				}
			}
			dots = dots[:fold.val]
		}
	}
	return countDots(dots)
}

func countDots(in [][]bool) int {
	result := 0
	for y := 0; y < len(in); y++ {
		for x := 0; x < len(in[y]); x++ {
			if in[y][x] {
				result += 1
			}
		}
	}
	return result
}

func formatDots(dots [][]bool) string {
	sb := strings.Builder{}
	for y := 0; y < len(dots); y++ {
		for x := 0; x < len(dots[y]); x++ {
			if dots[y][x] {
				sb.WriteString("◼")
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
