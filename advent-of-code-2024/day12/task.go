package day12

import (
	"fmt"
	"strings"
)

func part1(input string) int {
	visited := map[int]bool{}
	price := 0
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			i := x + y*height
			if !visited[i] {
				areaVisit := map[int]bool{}
				area, perimeter := visit(lines, r, x, y, height, width, areaVisit)
				price += area * perimeter

				for k := range areaVisit {
					visited[k] = true
				}
			}
		}
	}
	return price
}

func visit(input []string, r rune, x, y, height, width int, visited map[int]bool) (int, int) {
	if x < 0 || y < 0 || y >= height || x >= width || input[y][x] != byte(r) {
		return 0, 1
	}

	visited[x+y*height] = true

	area := 1
	perimeter := 0
	for _, d := range [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
		nx := x + d[0]
		ny := y + d[1]
		if !visited[nx+ny*height] {
			a, p := visit(input, r, nx, ny, height, width, visited)
			area += a
			perimeter += p
		}
	}

	return area, perimeter
}

func part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}
	return 0
}
