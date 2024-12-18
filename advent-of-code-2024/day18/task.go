package day18

import (
	"math"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type pos = utils.Coordinate2D

func part1(input string, drops, ex, ey int) int {
	blocked := map[pos]bool{}
	lines := strings.Split(input, "\n")
	for i := range drops {
		n := utils.GetAllNumbers(lines[i])
		x, y := n[0], n[1]
		blocked[pos{X: x, Y: y}] = true
	}

	open := []pos{{X: 0, Y: 0}}
	closed := map[pos]int{{X: 0, Y: 0}: 0}

	shortest := math.MaxInt
	ep := pos{X: ex, Y: ey}
	for len(open) > 0 {
		e := open[0]
		dist := closed[e]
		open = open[1:]

		if e == ep && shortest > dist {
			shortest = dist
		}

		for _, d := range utils.Get2DDirections() {
			np := e.Add(d)
			if np.X >= 0 && np.X <= ex && np.Y >= 0 && np.Y <= ey {
				if _, e := closed[np]; !e && !blocked[np] {
					open = append(open, np)
					closed[np] = dist + 1
				}
			}
		}
	}
	return shortest
}

func part2(input string, ex, ey int) []int {
	lines := strings.Split(input, "\n")

	low, high := 0, len(lines)
	for high-low > 1 {
		mid := low + (high-low)/2

		v := part1(input, mid, ex, ey)
		if v == math.MaxInt {
			high = mid
		} else {
			low = mid
		}
	}

	n := utils.GetAllNumbers(lines[low])
	return []int{n[0], n[1]}
}
