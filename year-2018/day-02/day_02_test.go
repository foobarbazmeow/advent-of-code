package day_02

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"strings"
	"testing"
)

func Test02Part1(t *testing.T) {
	file, _ := os.Open("day_02.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	two, three := 0, 0
	for _, str := range strings.Split(string(bytes), "\n") {
		m := map[byte]int{}
		for i := 0; i < len(str); i++ {
			m[str[i]]++
		}
		addTwo, addThree := false, false
		for _, v := range m {
			switch v {
			case 2:
				{
					if !addTwo {
						addTwo = true
					}
				}
			case 3:
				{
					if !addThree {
						addThree = true
					}
				}
			}
		}
		if addTwo {
			two++
		}
		if addThree {
			three++
		}
	}
	assert.Equal(t, 7688, two*three)
}

func Test02Part2(t *testing.T) {
	file, _ := os.Open("day_02.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)

	xs := []string{}
	for _, str := range strings.Split(string(bytes), "\n") {
		xs = append(xs, str)
	}

	findSameStrings := func() (string, int) {
		for i := 0; i < len(xs); i++ {
			str1 := xs[i]
			for j := 0; j < len(xs); j++ {
				diff, idx := 0, 0
				str2 := xs[j]
				for k := 0; k < len(str1); k++ {
					if str1[k] != str2[k] {
						diff += 1
						idx = k
					}
					if diff > 1 {
						goto ret
					}
				}
				if diff == 1 {
					return str1, idx
				}
			ret:
			}
		}
		return "", -1
	}

	str, idx := findSameStrings()

	assert.Equal(t, "lsrivmotzbdxpkxnaqmuwcchj", str[:idx]+str[idx+1:])
}
