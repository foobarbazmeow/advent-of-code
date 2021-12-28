package day_25

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	str := read("day_25.in")
	lines := strings.Split(str, "\n")

	xs1 := make([][]byte, len(lines))
	for i, line := range lines {
		xs1[i] = []byte(line)
	}

	result := 1
	for {
		xs2, mod := step(xs1)
		if !mod {
			break
		} else {
			xs1 = xs2
			result++
		}
	}
	assert.Equal(t, 441, result)
}

func read(filepath string) string {
	file, _ := os.Open(filepath)
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	return string(bytes)
}

func step(src [][]byte) ([][]byte, bool) {
	move1, mod1 := immutableMove(src, moveRight)
	move2, mod2 := immutableMove(move1, moveDown)
	return move2, mod1 || mod2
}

func immutableMove(src [][]byte, fn func(int, int, [][]byte) []cmd) ([][]byte, bool) {
	commands := make([]cmd, 0)
	for y := 0; y < len(src); y++ {
		for x := 0; x < len(src[y]); x++ {
			cmds := fn(x, y, src)
			if cmds != nil {
				commands = append(commands, cmds...)
			}
		}
	}
	for _, c := range commands {
		src[c.p.y][c.p.x] = c.b
	}
	return src, len(commands) > 0
}

func moveRight(x, y int, src [][]byte) []cmd {
	if src[y][x] != '>' {
		return nil
	}

	if x == len(src[0])-1 {
		if src[y][0] == '.' {
			return []cmd{
				{pair{0, y}, '>'},
				{pair{x, y}, '.'},
			}
		}
	} else {
		if src[y][x+1] == '.' {
			return []cmd{
				{pair{x + 1, y}, '>'},
				{pair{x, y}, '.'},
			}
		}
	}

	return nil
}

func moveDown(x, y int, src [][]byte) []cmd {
	if src[y][x] != 'v' {
		return nil
	}

	if y == len(src)-1 {
		if src[0][x] == '.' {
			return []cmd{
				{pair{x, 0}, 'v'},
				{pair{x, y}, '.'},
			}
		}
	} else {
		if src[y+1][x] == '.' {
			return []cmd{
				{pair{x, y + 1}, 'v'},
				{pair{x, y}, '.'},
			}
		}
	}

	return nil
}

type cmd struct {
	p pair
	b byte
}

type pair struct {
	x, y int
}
