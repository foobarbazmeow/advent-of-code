package day_21

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	tcs := []struct {
		p1, p2 int
		out    int
	}{
		// Example
		{4, 8, 739785},

		// Part 1
		{8, 2, 513936},
	}

	for _, tc := range tcs {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			var (
				player1 = tc.p1
				player2 = tc.p2
				score1  = 0
				score2  = 0

				limit = 1000

				dice        = 0
				result      = 0
				rollCounter = 0
			)

			for {
				dice, player1, score1 = turn(player1, score1, dice, &rollCounter)
				if score1 >= limit {
					result = rollCounter * score2
					break
				}
				dice, player2, score2 = turn(player2, score2, dice, &rollCounter)
				if score2 >= limit {
					result = rollCounter * score1
					break
				}
			}

			assert.Equal(t, tc.out, result)
		})
	}
}

func turn(position, score, dice int, counter *int) (int, int, int) {
	dice = roll100(dice, counter)
	a := dice
	dice = roll100(dice, counter)
	a += dice
	dice = roll100(dice, counter)
	a += dice

	position = move(position, a)
	score += position

	return dice, position, score
}

func move(c, a int) int {
	return (((c + a) - 1) % 10) + 1
}

func roll100(d int, counter *int) int {
	*counter++
	return d%100 + 1
}
