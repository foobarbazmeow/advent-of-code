package year_2021

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	ZeroSectors  = []byte{0, 1, 2, 4, 5, 6}
	OneSectors   = []byte{2, 5}
	TwoSectors   = []byte{0, 2, 3, 4, 6}
	ThreeSectors = []byte{0, 2, 3, 5, 6}
	FourSectors  = []byte{1, 2, 3, 5}
	FiveSectors  = []byte{0, 1, 3, 5, 6}
	SixSectors   = []byte{0, 1, 3, 4, 5, 6}
	SevenSectors = []byte{0, 2, 5}
	EightSectors = []byte{0, 1, 2, 3, 4, 5, 6}
	NineSectors  = []byte{0, 1, 2, 3, 5, 6}

	Numbers = [][]byte{
		ZeroSectors,
		OneSectors,
		TwoSectors,
		ThreeSectors,
		FourSectors,
		FiveSectors,
		SixSectors,
		SevenSectors,
		EightSectors,
		NineSectors,
	}
)

func getNumber(xs []byte) int {
	cmp := func(a, b []byte) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})
	for i, num := range Numbers {
		if !cmp(num, xs) {
			continue
		}
		return i
	}
	return -1
}

func Test08Part1(t *testing.T) {
	lines2 := readLines("day_08.in")
	counter2 := countNumbersWithUniqueLength(lines2)
	assert.Equal(t, 392, counter2)
}

func countNumbersWithUniqueLength(lines []string) int {
	m := map[int]int{}
	for _, line := range lines {
		numbers := strings.Split(strings.TrimSpace(strings.Split(line, "|")[1]), " ")
		for _, num := range numbers {
			m[len(num)]++
		}
	}
	return m[len(OneSectors)] + m[len(FourSectors)] + m[len(SevenSectors)] + m[len(EightSectors)]
}

func Test08Part2(t *testing.T) {
	str := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
	assert.Equal(t, 5353, encodeNumbers(str))

	lines2 := readLines("day_08.in")
	sum2 := 0
	for _, s := range lines2 {
		sum2 += encodeNumbers(s)
	}
	assert.Equal(t, 1004688, sum2)
}

func encodeNumbers(str string) int {
	parts := strings.Split(str, "|")
	digits := strings.Split(strings.TrimSpace(parts[0]), " ")
	number := strings.Split(strings.TrimSpace(parts[1]), " ")

	sectors := make([]byte, 7)

	// (7) - (1) = sectors[0]
	one := single(digits, func(str string) bool { return len(str) == 2 })
	seven := single(digits, func(str string) bool { return len(str) == 3 })
	sectors[0] = byte(sumRunes(seven) - sumRunes(one))

	// (4) - (1) = sectors[1] & sectors[3]
	four := single(digits, func(str string) bool { return len(str) == 4 })
	sectors13 := filter(four, one...)

	// (8) - sectors[0] - (1) - sectors[1] - sectors[3] = sectors[4] & sectors[6]
	eight := single(digits, func(str string) bool { return len(str) == 7 })
	sectors46 := filter(filter(filter(eight, sectors[0]), one...), sectors13...)

	// (3) - sectors[0] - (1) - sectors[1] - sectors[3] = sectors[3]
	tree := single(digits, func(str string) bool {
		if len(str) != 5 {
			return false
		}

		m := map[byte]struct{}{}
		for i := 0; i < len(str); i++ {
			m[str[i]] = struct{}{}
		}

		xs := []byte{sectors[0], one[0], one[1]}
		for _, v := range xs {
			if _, ok := m[v]; !ok {
				return false
			}
		}

		return true
	})

	sectors36 := filter(filter(tree, sectors[0]), one...)
	sectors[6] = filter(sectors36, sectors13...)[0]
	sectors[3] = byte(sumRunes(sectors36) - int(sectors[6]))
	sectors[1] = byte(sumRunes(sectors13) - int(sectors[3]))
	sectors[4] = byte(sumRunes(sectors46) - int(sectors[6]))

	// (3) - sectors[0] - (1) - sectors[1] - sectors[3] = sectors[3]
	six := single(digits, func(str string) bool {
		if len(str) != 6 {
			return false
		}

		m := map[byte]struct{}{}
		for i := 0; i < len(str); i++ {
			m[str[i]] = struct{}{}
		}

		xs := []byte{sectors[0], sectors[1], sectors[3], sectors[4], sectors[6]}
		for _, v := range xs {
			if _, ok := m[v]; !ok {
				return false
			}
		}

		return true
	})
	sectors[5] = byte(sumRunes(six) - (sumRunes([]byte{sectors[0], sectors[1], sectors[3], sectors[4], sectors[6]})))
	sectors[2] = byte(sumRunes(one) - int(sectors[5]))

	sum := 0
	rev := reverseSlice(sectors)
	for _, n := range number {
		na := make([]byte, 0)
		for i := 0; i < len(n); i++ {
			u := n[i]
			i2 := rev[u]
			na = append(na, byte(i2))
		}
		sum = sum*10 + getNumber(na)
	}
	return sum
}

/*
   0000
  1    2
  1    2
   3333
  4    5
  4    5
   6666
*/

func sumRunes(str []byte) int {
	sum := 0
	for i := 0; i < len(str); i++ {
		sum += int(str[i])
	}
	return sum
}

func single(xs []string, predicate func(string) bool) []byte {
	for _, str := range xs {
		if predicate((str)) {
			return []byte(str)
		}
	}
	panic("string not found")
}

func filter(xs1 []byte, xs2 ...byte) []byte {
	m := map[byte]struct{}{}
	for _, v := range xs1 {
		m[v] = struct{}{}
	}
	for _, v := range xs2 {
		if _, ok := m[v]; ok {
			delete(m, v)
		}
	}
	xs := []byte{}
	for k, _ := range m {
		xs = append(xs, k)
	}
	return xs
}

func reverseSlice(xs []byte) map[byte]int {
	m := map[byte]int{}
	for i, v := range xs {
		m[v] = i
	}
	return m
}
