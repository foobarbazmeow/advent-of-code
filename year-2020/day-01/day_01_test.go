package day_01

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"testing"
)

func Test01Part1(t *testing.T) {
	const target = 2020

	file, err := os.Open("day_01.in")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	m := map[int]struct{}{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num1, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error(err)
		}
		num2 := target - num1
		if _, ok := m[num2]; ok {
			t.Log(num1 * num2)
			return
		} else {
			m[num1] = struct{}{}
		}
	}
}

func Test01Part2(t *testing.T) {
	const target = 2020

	file, err := os.Open("day_01.in")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	xs := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			t.Error(err)
		}
		xs = append(xs, num)
	}

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	sort.Ints(xs)

	// 1, 2, 3, 4, 5
	// i  j        z
	for i := 0; i < len(xs); i++ {
		j, z := i+1, len(xs)-1
		for j < z {
			num := xs[i] + xs[j] + xs[z]
			if num < target {
				j++
			} else if num > target {
				z--
			} else {
				t.Log(xs[i] * xs[j] * xs[z])
				return
			}
		}
	}
}
