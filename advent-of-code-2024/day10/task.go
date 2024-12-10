package day10

import (
	"strings"
)

type pos struct {
	x, y int
}

func part1(input string) int {
	return solve(input, true)
}

func part2(input string) int {
	return solve(input, false)
}

func solve(input string, cached bool) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, r := range line {
			if r == '0' {
				fromThis := routesFrom(lines, x, y, byte(r), map[pos]bool{}, cached)
				sum += fromThis
			}
		}
	}
	return sum
}

func routesFrom(line []string, x, y int, cur byte, walked map[pos]bool, cached bool) int {
	if cached {
		if walked[pos{x: x, y: y}] {
			return 0
		}
		walked[pos{x: x, y: y}] = true
	}

	if line[y][x] == '9' {
		return 1
	}

	opt := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	routes := 0
	for _, d := range opt {
		nx := x + d[0]
		ny := y + d[1]
		if nx >= 0 && ny >= 0 && ny < len(line) && nx < len(line[0]) && line[ny][nx] == cur+1 {
			routes += routesFrom(line, nx, ny, cur+1, walked, cached)
		}
	}

	return routes
}
