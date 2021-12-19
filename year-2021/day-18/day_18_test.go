package day_18

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

/*
func TestExplode(t *testing.T) {
	tcs := []struct {
		in1  string
		out1 bool
		out2 string
	}{
		{"[[[[[9,8],1],2],3],4]", true, "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", true, "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", true, "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", true, "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", true, "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
		{"[1,[2,3]]", false, "[1,[2,3]]"},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			e, _ := parse(tc.in1, 0, 0)
			assert.Equal(t, tc.out1, e.explode())
			assert.Equal(t, tc.out2, e.String())
		})
	}
}

func TestMagnitude(t *testing.T) {
	e4, _ := parse("[[1,2],[[3,4],5]]", 0, 0)
	assert.Equal(t, 143, e4.magnitude())
}

func TestCreateSplit(t *testing.T) {
	e11 := split(0, 10, nil)
	assert.Equal(t, 0, e11.level)
	assert.Nil(t, e11.parent)
	assert.Equal(t, 1, e11.left.level)
	assert.Equal(t, 5, e11.left.val)
	assert.Equal(t, e11, e11.left.parent)
	assert.Equal(t, 1, e11.right.level)
	assert.Equal(t, 5, e11.right.val)
	assert.Equal(t, e11, e11.right.parent)

	e12 := split(0, 11, nil)
	assert.Equal(t, 0, e12.level)
	assert.Nil(t, e12.parent)
	assert.Equal(t, 1, e12.left.level)
	assert.Equal(t, 5, e12.left.val)
	assert.Equal(t, e12, e12.left.parent)
	assert.Equal(t, 1, e12.right.level)
	assert.Equal(t, 6, e12.right.val)
	assert.Equal(t, e12, e12.right.parent)
}

func TestSeq(t *testing.T) {
	e15, _ := parse("[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", 0, 0)
	e15.explode()
	assert.Equal(t, "[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", e15.String())
	e15.explode()
	assert.Equal(t, "[[[[0,7],4],[15,[0,13]]],[1,1]]", e15.String())
	e15.split()
	assert.Equal(t, "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", e15.String())
	e15.split()
	assert.Equal(t, "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", e15.String())
	e15.explode()
	assert.Equal(t, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", e15.String())
}

func TestSum(t *testing.T) {
	tcs := []struct {
		in  string
		out string
	}{
		{"[[[[4,3],4],4],[7,[[8,4],9]]]\n[1,1]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
		{"[1,1]\n[2,2]\n[3,3]\n[4,4]", "[[[[1,1],[2,2]],[3,3]],[4,4]]"},
		{"[1,1]\n[2,2]\n[3,3]\n[4,4]\n[5,5]", "[[[[3,0],[5,3]],[4,4]],[5,5]]"},
		{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]\n[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]\n[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]\n[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]\n[7,[5,[[3,8],[1,4]]]]\n[[2,[2,2]],[8,[8,1]]]\n[2,9]\n[1,[[[9,3],9],[[9,0],[0,7]]]]\n[[[5,[7,4]],7],1]\n[[[[4,2],2],6],[8,7]]", "[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]"},
		{"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]\n[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]", "[[[[7,8],[6,6]],[[6,0],[7,7]]],[[[7,8],[8,8]],[[7,9],[0,6]]]]"},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			lines := strings.Split(tc.in, "\n")
			xs := make([]*expr, len(lines))
			for i, line := range lines {
				e, _ := parse(line, 0, 0)
				xs[i] = e
			}
			e := xs[0]
			for i := 1; i < len(xs); i++ {
				e = e.sum(xs[i])
			}
			assert.Equal(t, tc.out, e.String())
		})
	}
}

func TestMaxPair(t *testing.T) {
	str := "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]\n[[[5,[2,8]],4],[5,[[9,9],0]]]\n[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]\n[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]\n[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]\n[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]\n[[[[5,4],[7,7]],8],[[8,3],8]]\n[[9,3],[[9,9],[6,[4,9]]]]\n[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]\n[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]"
	lines := strings.Split(str, "\n")

	max := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			// Maybe...
			// One day i'll rewrite it for value-types
			xs := make([]*expr, len(lines))
			for k, line := range lines {
				e, _ := parse(line, 0, 0)
				xs[k] = e
			}

			m := xs[i].sum(xs[j]).magnitude()
			if m > max {
				max = m
			}
		}
	}

	assert.Equal(t, 3993, max)
}
*/

func Test18Part1(t *testing.T) {
	file, _ := os.Open("day_18.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	lines := strings.Split(string(bytes), "\n")
	xs := make([]*expr, len(lines))
	for i, line := range lines {
		e, _ := parse(line, 0, 0)
		xs[i] = e
	}
	e := xs[0]
	for i := 1; i < len(xs); i++ {
		e = e.sum(xs[i])
	}
	assert.Equal(t, 3869, e.magnitude())
}

func Test18Part2(t *testing.T) {
	file, _ := os.Open("day_18.in")
	defer file.Close()
	bytes, _ := io.ReadAll(file)
	lines := strings.Split(string(bytes), "\n")

	max := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			// Maybe...
			// One day i'll rewrite it for value-types
			xs := make([]*expr, len(lines))
			for k, line := range lines {
				e, _ := parse(line, 0, 0)
				xs[k] = e
			}

			m := xs[i].sum(xs[j]).magnitude()
			if m > max {
				max = m
			}
		}
	}
	assert.Equal(t, 4671, max)
}

func parse(str string, idx int, level int) (*expr, int) {
	e := &expr{level: level}

	if str[idx] >= '0' && str[idx] <= '9' {
		e.val = int(str[idx] - '0')
		return e, idx + 1
	}

	l, idx := parse(str, idx+1, level+1)
	l.parent = e
	e.left = l

	r, idx := parse(str, idx+1, level+1)
	r.parent = e
	e.right = r

	return e, idx + 1
}

type expr struct {
	left, right, parent *expr
	val                 int
	level               int
}

func split(level, val int, parent *expr) *expr {
	p := &expr{level: level, parent: parent}
	p.left = &expr{val: val / 2, parent: p, level: level + 1}
	p.right = &expr{val: int(math.Ceil(float64(val) / 2.0)), parent: p, level: level + 1}
	return p
}

func (e *expr) fromLeft() bool {
	c, prev := e, e
	for ; c.parent != nil; prev, c = c, c.parent {
	}
	return c.left == prev
}

func (e *expr) modifyLevel(delta int) *expr {
	e.level += delta
	if e.left != nil {
		e.left.modifyLevel(delta)
	}
	if e.right != nil {
		e.right.modifyLevel(delta)
	}
	return e
}

func (e *expr) magnitude() int {
	r := e.val
	if e.left != nil {
		r += 3 * e.left.magnitude()
	}
	if e.right != nil {
		r += 2 * e.right.magnitude()
	}
	return r
}

func (e *expr) isLiteral() bool {
	return e.left == nil && e.right == nil
}

func (e *expr) explode() bool {
	if e.level == 4 && !e.isLiteral() {
		tmp, p := e, e.parent
		if p.left == e {
			// Replace
			// [[1,2],3] => [0,3]
			p.left = &expr{val: 0, level: p.level + 1, parent: p}

			// Add value to the right
			// [[1,2],3] => [0,3+2]
			if p.right.isLiteral() {
				p.right.val += tmp.right.val
			} else {
				l := p.right.left
				for ; !l.isLiteral(); l = l.left {
				}
				l.val += tmp.right.val
			}

			// Add left value
			for prev, cur := p, p.parent; cur != nil; prev, cur = cur, cur.parent {
				if cur.left != prev {
					l := cur.left
					for ; !l.isLiteral(); l = l.right {
					}
					l.val += tmp.left.val
					break
				}
			}
		} else if p.right == e {
			// Replace
			// [1,[2,3]] => [1,0]
			p.right = &expr{val: 0, level: p.level + 1, parent: p}

			// Add value to the left
			// [1,[2,3]] => [1+2,0]
			if p.left.isLiteral() {
				p.left.val += tmp.left.val
			} else {
				r := p.left.right
				for ; !r.isLiteral(); r = r.right {
				}
				r.val += tmp.left.val
			}

			// Add right value
			for prev, cur := p, p.parent; cur != nil; prev, cur = cur, cur.parent {
				if cur.right != prev {
					l := cur.right
					for ; !l.isLiteral(); l = l.left {
					}
					l.val += tmp.right.val
					break
				}
			}
		}
		return true
	}
	if e.left != nil && e.left.explode() {
		return true
	}
	if e.right != nil {
		return e.right.explode()
	}
	return false
}

func (e *expr) split() bool {
	if e.isLiteral() && e.val > 9 {
		tmp, p := e, e.parent
		if p.left == e {
			p.left = split(tmp.level, tmp.val, tmp.parent)
		} else if p.right == e {
			p.right = split(tmp.level, tmp.val, tmp.parent)
		}
		return true
	}
	if e.left != nil && e.left.split() {
		return true
	}
	if e.right != nil {
		return e.right.split()
	}
	return false
}

func (l *expr) sum(r *expr) *expr {
	root := &expr{level: 0}
	l.parent = root
	r.parent = root
	root.left = l.modifyLevel(1)
	root.right = r.modifyLevel(1)
	for {
		if root.explode() {
			continue
		}
		if root.split() {
			continue
		}
		break
	}
	return root
}

func (e *expr) String() string {
	if e.isLiteral() {
		return strconv.Itoa(e.val)
	}
	return fmt.Sprintf("[%s,%s]", e.left.String(), e.right.String())
}
