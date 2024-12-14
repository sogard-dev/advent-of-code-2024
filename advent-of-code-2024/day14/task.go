package day14

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type robot struct {
	x, y, vx, vy int
}

type pos struct {
	x, y int
}

func part1(input string, width, height int) int {
	q := []int{0, 0, 0, 0}
	for _, line := range strings.Split(input, "\n") {
		x, y, vx, vy := parse(line, width, height)

		x = (x + vx*100) % width
		y = (y + vy*100) % height
		qx, qy := getQuadrant(x, width, y, height)
		if qx >= 0 && qy >= 0 {
			q[qx+qy*2]++
		}
	}

	s := 1
	for _, v := range q {
		s *= v
	}

	return s
}

func getQuadrant(x int, width int, y int, height int) (int, int) {
	qx := -1
	if x < width/2 {
		qx = 0
	}
	if x > width/2 {
		qx = 1
	}

	qy := -1
	if y < height/2 {
		qy = 0
	}
	if y > height/2 {
		qy = 1
	}

	return qx, qy
}

func part2(input string, width, height int) int {
	robots := []robot{}
	for _, line := range strings.Split(input, "\n") {
		x, y, vx, vy := parse(line, width, height)
		robots = append(robots, robot{x, y, vx, vy})
	}

	directions := []pos{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

	bc := 0
	for i := range 100000000000 {
		positions := map[pos]bool{}
		for i, r := range robots {
			robots[i].x = (r.x + r.vx) % width
			robots[i].y = (r.y + r.vy) % height
			positions[pos{robots[i].x, robots[i].y}] = true
		}

		seen := map[pos]bool{}
		for p := range positions {
			if !seen[p] {
				open := []pos{p}
				clSize := 0
				for len(open) > 0 {
					e := open[0]
					open = open[1:]
					clSize++
					for _, d := range directions {
						np := pos{e.x + d.x, e.y + d.y}
						if positions[np] && !seen[np] {
							seen[np] = true
							open = append(open, np)
						}
					}
				}
				bc = max(bc, clSize)
			}
		}

		if bc > 200 {
			printThis(i, height, width, positions)
			return i + 1
		}
	}
	return 0
}

func parse(line string, width int, height int) (int, int, int, int) {
	nums := utils.GetAllNumbers(line)
	x, y := nums[0], nums[1]
	vx, vy := nums[2], nums[3]
	for vx < 0 {
		vx += width
	}
	for vy < 0 {
		vy += height
	}
	return x, y, vx, vy
}

func printThis(round int, height int, width int, positions map[pos]bool) {
	println("Round", round+1)
	for y := range height {
		for x := range width {
			if positions[pos{x, y}] {
				print("X")
			} else {
				print(" ")
			}
		}
		println()
	}
}
