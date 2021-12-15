package year_2020

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test08Part1(t *testing.T) {
	file2, _ := os.Open("day_08.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	instructions2 := parse(lines2)
	result2 := execute(instructions2)
	assert.Equal(t, 1801, result2)
}

func Test08Part2(t *testing.T) {
	file2, _ := os.Open("day_08_fixed.in")
	defer file2.Close()
	bytes2, _ := io.ReadAll(file2)
	lines2 := strings.Split(string(bytes2), "\n")
	instructions2 := parse(lines2)
	result2 := executeAll(instructions2)
	// line #211
	// replace jmp -31 => nop -31
	assert.Equal(t, 2060, result2)
}

type instr struct {
	op  string
	arg int
}

func parse(lines []string) []instr {
	result := make([]instr, len(lines))
	for i := 0; i < len(lines); i++ {
		ins := instr{}
		fmt.Sscanf(lines[i], "%s %d", &ins.op, &ins.arg)
		result[i] = ins
	}
	return result
}

func execute(xs []instr) int {
	acc, current := 0, 0
	executed := map[int]struct{}{}
	for {
		if _, ok := executed[current]; ok {
			return acc
		} else {
			executed[current] = struct{}{}
		}
		instr := xs[current]
		if instr.op == "nop" {
			current += 1
		} else if instr.op == "acc" {
			acc += instr.arg
			current += 1
		} else if instr.op == "jmp" {
			current += instr.arg
		}
	}
}

func executeAll(xs []instr) int {
	acc, prev, current := 0, 0, 0
	executed := map[int][]int{}

	detect := func() int {
		for k, v := range executed {
			if len(v) > 1 && v[len(v)-2] == v[len(v)-1] {
				return k
			}
		}
		return -1
	}

	print := func() string {
		sb := strings.Builder{}
		for k, v := range executed {
			if len(v) > 1 {
				sb.WriteString(fmt.Sprintf("idx=%d,prevs=%v\n", k, v))
			}
		}
		return sb.String()
	}

	for i := 0; ; i++ {
		if current > len(xs)-1 {
			return acc
		}
		executed[current] = append(executed[current], prev)
		prev = current
		instr := xs[current]
		if instr.op == "nop" {
			current += 1
		} else if instr.op == "acc" {
			acc += instr.arg
			current += 1
		} else if instr.op == "jmp" {
			current += instr.arg
			if d := detect(); d != -1 {
				fmt.Println(print())
				return -1
			}
		}
	}
}
