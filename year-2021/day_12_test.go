package year_2021

import (
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

var (
	smallCavePattern = regexp.MustCompile("^[a-z]+$")
)

func Test12Part1(t *testing.T) {
	file4, _ := os.Open("day_12.in")
	defer file4.Close()
	bytes4, _ := io.ReadAll(file4)
	assert.Equal(t, 5104, buildPaths(string(bytes4)))
}

func Test12Part2(t *testing.T) {
	file4, _ := os.Open("day_12.in")
	defer file4.Close()
	bytes4, _ := io.ReadAll(file4)
	assert.Equal(t, 149220, buildPaths2(string(bytes4)))
}

type cave struct {
	name       string
	neighbours []*cave
}

func newCave(name string) *cave {
	return &cave{name, make([]*cave, 0)}
}

func (c *cave) addNeighbour(n *cave) {
	c.neighbours = append(c.neighbours, n)
}

func parseCaves(content string) map[string]*cave {
	caves := map[string]*cave{}
	for _, s := range strings.Split(content, "\n") {
		parts := strings.Split(s, "-")
		a := parts[0]
		b := parts[1]
		ac := newCave(a)
		bc := newCave(b)
		if _, ok := caves[a]; !ok {
			caves[a] = ac
		}
		if _, ok := caves[b]; !ok {
			caves[b] = bc
		}
		caves[a].addNeighbour(caves[b])
		caves[b].addNeighbour(caves[a])
	}
	return caves
}

func buildPaths(content string) int {
	caves := parseCaves(content)
	failed := map[string]struct{}{}
	success := map[string]struct{}{}
	caveIteration([]*cave{caves["start"]}, failed, success)
	return len(success)
}

func buildPaths2(content string) int {
	caves := parseCaves(content)
	failed := map[string]struct{}{}
	success := map[string]struct{}{}
	caveIteration2([]*cave{caves["start"]}, failed, success)
	return len(success)
}

func caveIteration(route []*cave, failed map[string]struct{}, success map[string]struct{}) {
	str := formatRoute(route)
	if _, ok := failed[str]; ok {
		return
	}
	last := route[len(route)-1]
	if last.name == "end" {
		if _, ok := success[str]; !ok {
			success[str] = struct{}{}
		}
		return
	}
	if alreadyVisited(route[:len(route)-1], last) {
		failed[str] = struct{}{}
		return
	}
	for _, neighbour := range last.neighbours {
		caveIteration(append(route, neighbour), failed, success)
	}
}

func caveIteration2(route []*cave, failed map[string]struct{}, success map[string]struct{}) {
	str := formatRoute(route)
	if _, ok := failed[str]; ok {
		return
	}
	last := route[len(route)-1]
	if last.name == "end" {
		if _, ok := success[str]; !ok {
			success[str] = struct{}{}
		}
		return
	}
	if alreadyVisited2(route[:len(route)-1], last) {
		failed[str] = struct{}{}
		return
	}
	for _, neighbour := range last.neighbours {
		caveIteration2(append(route, neighbour), failed, success)
	}
}

func formatRoute(route []*cave) string {
	xs := make([]string, len(route))
	for i, c := range route {
		xs[i] = c.name
	}
	return strings.Join(xs, ",")
}

func alreadyVisited(route []*cave, c *cave) bool {
	if smallCavePattern.MatchString(c.name) {
		for _, r := range route {
			if r.name == c.name {
				return true
			}
		}
	}
	return false
}

func alreadyVisited2(route []*cave, c *cave) bool {
	if len(route) == 0 {
		return false
	}

	if c.name == "start" {
		return true
	}

	if smallCavePattern.MatchString(c.name) {
		visits := map[string]int{}
		for _, r := range route {
			visits[r.name]++
		}

		alreadyReentered := false
		for k, v := range visits {
			if smallCavePattern.MatchString(k) {
				if v >= 2 {
					alreadyReentered = true
				}
			}
		}

		for _, r := range route {
			if r.name == c.name {
				if alreadyReentered {
					return true
				}
			}
		}
	}
	return false
}
