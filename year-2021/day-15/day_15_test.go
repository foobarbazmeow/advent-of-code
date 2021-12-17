package day_15

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test15Part1(t *testing.T) {
	assert.Equal(t, 562, dijkstra(parse("day_15.in")))
}

func Test15Part2(t *testing.T) {
	assert.Equal(t, 2874, dijkstra(createBig(parse("day_15.in"))))
}

func parse(filepath string) [][]int {
	file, _ := os.Open(filepath)
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	content := strings.Split(string(bytes), "\n")

	xs := make([][]int, len(content))

	for y := 0; y < len(xs); y++ {
		xs[y] = make([]int, len(content[y]))
		for x := 0; x < len(content[y]); x++ {
			num, _ := strconv.Atoi(string(content[y][x]))
			xs[y][x] = num
		}
	}

	return xs
}

func dijkstra(xs [][]int) int {
	unvisited := map[pair]struct{}{}
	distances := map[pair]int{}

	q := &pq{}
	qItems := map[pair]*pqItem{}

	for y := 0; y < len(xs); y++ {
		for x := 0; x < len(xs[y]); x++ {
			p := pair{x, y}
			unvisited[p] = struct{}{}
			distances[p] = math.MaxInt

			qItem := &pqItem{
				priority: math.MaxInt,
				p:        p,
			}
			qItems[p] = qItem
			q.Push(qItem)
		}
	}

	distances[pair{0, 0}] = 0

	current := pair{0, 0}
	for len(unvisited) > 0 {
		for _, n := range _neighbours(current.x, current.y) {
			if sliceHasElement(xs, n.x, n.y) {
				if _, ok := unvisited[n]; ok {
					dist := distances[current] + xs[n.y][n.x]
					if dist < distances[n] {
						distances[n] = dist
						q.Update(qItems[n], dist)
					}
				}
			}
		}
		delete(unvisited, current)
		item := heap.Pop(q).(*pqItem)
		current = item.p
	}

	return distances[pair{len(xs[len(xs)-1]) - 1, len(xs) - 1}]
}

func _neighbours(x, y int) []pair {
	return []pair{
		// ↑
		{x, y - 1},
		// ->
		{x + 1, y},
		// ↓
		{x, y + 1},
		// <-
		{x - 1, y},
	}
}

type pair struct {
	x, y int
}

func sliceHasElement(xs [][]int, x, y int) bool {
	if y >= 0 && y < len(xs) {
		if x >= 0 && x < len(xs[0]) {
			return true
		}
	}
	return false
}

func createBig(small [][]int) [][]int {
	height := len(small)
	width := len(small[0])

	big := make([][]int, height*5)
	for y := 0; y < len(big); y++ {
		big[y] = make([]int, width*5)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			big[y][x] = small[y][x]
		}
	}

	for y := 0; y < height; y++ {
		for x := width; x < len(big[y]); x++ {
			val := big[y][((x/width-1)*width+x%width)] + 1
			if val > 9 {
				val = 1
			}
			big[y][x] = val
		}
	}

	for y := height; y < len(big); y++ {
		for x := 0; x < len(big[y]); x++ {
			val := big[(y/height-1)*height+y%height][x] + 1
			if val > 9 {
				val = 1
			}
			big[y][x] = val
		}
	}

	return big
}

type pqItem struct {
	index    int
	priority int

	p pair
}

type pq []*pqItem

func (h pq) Len() int           { return len(h) }
func (h pq) Less(i, j int) bool { return h[i].priority < h[j].priority }

func (h pq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *pq) Push(x interface{}) {
	idx := len(*h)
	newItem := x.(*pqItem)
	newItem.index = idx
	*h = append(*h, newItem)
}

func (h *pq) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	x.index = -1
	*h = old[0 : n-1]
	return x
}

func (h *pq) Update(item *pqItem, priority int) {
	item.priority = priority
	heap.Fix(h, item.index)
}
