package util

import (
	"crypto/sha1"
	"encoding/base64"
)

func SumInts(xs []int) int {
	result := 0
	for _, v := range xs {
		result += v
	}
	return result
}

func ByteSliceHasElement(x, y int, xs [][]byte) bool {
	if y >= 0 && y <= len(xs)-1 {
		if x >= 0 && x <= len(xs[0])-1 {
			return true
		}
	}
	return false
}

func AdjacentWithDiagonals(x, y int) []IntPair {
	return []IntPair{
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

func ByteSliceHash(xs [][]byte) string {
	h := sha1.New()
	for _, row := range xs {
		h.Write(row)
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func CountInByteSlice(val byte, x, y int, xs [][]byte) int {
	result := 0
	for _, p := range AdjacentWithDiagonals(x, y) {
		if ByteSliceHasElement(p.X, p.Y, xs) {
			if xs[p.Y][p.X] == val {
				result += 1
			}
		}
	}
	return result
}
