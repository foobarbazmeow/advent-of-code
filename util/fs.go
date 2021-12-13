package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(filepath string) []int {
	xs := make([]int, 0)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		xs = append(xs, num)
	}
	return xs
}

func ReadBytes(filepath string) [][]byte {
	xs := make([][]byte, 0)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, []byte(scanner.Text()))
	}
	return xs
}

func ReadLines(filepath string) []string {
	xs := make([]string, 0)
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xs = append(xs, scanner.Text())
	}
	return xs
}
