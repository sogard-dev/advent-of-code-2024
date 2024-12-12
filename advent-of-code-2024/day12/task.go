package day12

import (
	"strings"
)

type pos struct {
	x, y int
}

func part1(input string) int {
	price := 0
	solve(input, func(a int, p int, _ map[pos]map[pos]bool) {
		price += a * p
	})
	return price
}

func part2(input string) int {
	price := 0
	solve(input, func(a int, _ int, p map[pos]map[pos]bool) {
		sidesForThis := 0
		for _, d := range getDirections() {
			s := sides(p[d])
			sidesForThis += s
		}
		price += a * sidesForThis

	})
	return price
}

func sides(positions map[pos]bool) int {
	seen := map[pos]bool{}
	sides := 0
	for p := range positions {
		if !seen[p] {
			seen[p] = true

			sides += 1
			open := []pos{p}
			for len(open) > 0 {
				e := open[0]
				open = open[1:]
				for _, d := range getDirections() {
					np := pos{e.x + d.x, e.y + d.y}

					if positions[np] && !seen[np] {
						seen[np] = true
						open = append(open, np)
					}
				}
			}
		}
	}

	return sides
}

func solve(input string, f func(a int, p int, pe map[pos]map[pos]bool)) {
	visited := map[pos]bool{}
	lines := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			if !visited[pos{x, y}] {
				areaVisit := map[pos]bool{}
				area, perimeter, perimeters := visit(lines, r, x, y, height, width, areaVisit)
				f(area, perimeter, perimeters)

				for k := range areaVisit {
					visited[k] = true
				}
			}
		}
	}
}

func visit(input []string, r rune, initX, initY, height, width int, visited map[pos]bool) (int, int, map[pos]map[pos]bool) {
	var rec func(int, int) (int, int)

	perimeters := map[pos]map[pos]bool{}

	rec = func(x, y int) (int, int) {
		if x < 0 || y < 0 || y >= height || x >= width || input[y][x] != byte(r) {
			return 0, 1
		}

		visited[pos{x, y}] = true

		area := 1
		perimeter := 0
		for _, d := range getDirections() {
			np := pos{x + d.x, y + d.y}
			if !visited[np] {
				a, p := rec(np.x, np.y)
				if p == 1 && a == 0 {
					if _, e := perimeters[d]; !e {
						perimeters[d] = map[pos]bool{}
					}
					perimeters[d][pos{x, y}] = true
				}
				area += a
				perimeter += p
			}
		}

		return area, perimeter
	}

	a, p := rec(initX, initY)

	return a, p, perimeters
}

func getDirections() [4]pos {
	return [4]pos{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
}
