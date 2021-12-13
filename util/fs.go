package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInts(filepath string) []int {
	xs := make([]int, 0)
	for str := range ReadLines(filepath) {
		num, _ := strconv.Atoi(str)
		xs = append(xs, num)
	}
	return xs
}

func ReadBytes(filepath string) [][]byte {
	xs := make([][]byte, 0)
	for str := range ReadLines(filepath) {
		xs = append(xs, []byte(str))
	}
	return xs
}

func ReadLines(filepath string) <-chan string {
	ch := make(chan string)
	go func() {
		file, _ := os.Open(filepath)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()
	return ch
}
