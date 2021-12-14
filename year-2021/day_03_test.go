package year_2021

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test03Part1(t *testing.T) {
	strings := readLines("day_03.in")
	gamma, epsilon := make([]uint8, 12), make([]uint8, 12)
	for i := 0; i < 12; i++ {
		xs := make([]uint8, len(strings))
		for j := 0; j < len(strings); j++ {
			xs[j] = strings[j][i]
		}
		z, o := 0, 0
		for _, x := range xs {
			switch x {
			case '0':
				z++
			case '1':
				o++
			}
		}
		if z > o {
			gamma[i] = '0'
			epsilon[i] = '1'
		} else {
			gamma[i] = '1'
			epsilon[i] = '0'
		}
	}
	gammaRate, _ := strconv.ParseInt(string(gamma), 2, 32)
	epsilonRate, _ := strconv.ParseInt(string(epsilon), 2, 32)
	assert.Equal(t, int64(4174964), gammaRate*epsilonRate)
}

func Test03Part2(t *testing.T) {
	strings := readLines("day_03.in")

	data := make([]int, len(strings))
	for i := 0; i < len(strings); i++ {
		data[i] = i
	}

	oxygen := search(strings, 0, data, true)
	co2 := search(strings, 0, data, false)

	oxygenRate, _ := strconv.ParseInt(oxygen, 2, 32)
	co2Rate, _ := strconv.ParseInt(co2, 2, 32)
	assert.Equal(t, int64(4474944), oxygenRate*co2Rate)
}

func search(strings []string, idx int, data []int, greater bool) string {
	if len(data) == 1 {
		return strings[data[0]]
	}

	z, o := make([]int, 0), make([]int, 0)
	for i := 0; i < len(data); i++ {
		u := strings[data[i]][idx]
		if u == '0' {
			z = append(z, data[i])
		} else {
			o = append(o, data[i])
		}
	}

	if greater {
		if len(o) >= len(z) {
			return search(strings, idx+1, o, greater)
		}
		return search(strings, idx+1, z, greater)
	} else {
		if len(z) <= len(o) {
			return search(strings, idx+1, z, greater)
		}
		return search(strings, idx+1, o, greater)
	}
}
