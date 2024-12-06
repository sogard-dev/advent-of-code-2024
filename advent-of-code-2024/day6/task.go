package day6

import (
	"strings"
)

type pos struct {
	x, y int
}

type posAndDir struct {
	x, y, d int
}

func part1(input string) int {
	visited, _ := solve(input)
	return len(visited)
}

func part2(input string) int {
	sum := 0
	for i := range input {
		if input[i] == '.' {
			replaced := replaceAtIndex(input, '#', i)
			_, looped := solve(replaced)
			if looped {
				sum += 1
			}
		}
	}
	return sum
}

func solve(input string) (map[pos]bool, bool) {
	obstructions := map[pos]bool{}
	guard := posAndDir{}
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, r := range line {
			if r == '#' {
				obstructions[pos{x: x, y: y}] = true
			} else if r == '^' {
				guard.x = x
				guard.y = y
			}
		}
	}

	directions := []pos{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}

	visited := map[posAndDir]bool{}
	looped := false
	for guard.x >= 0 && guard.x < width && guard.y >= 0 && guard.y < height {
		if visited[guard] {
			looped = true
			break
		}
		visited[guard] = true

		for turns := 0; turns < 4; turns++ {
			thisDir := (guard.d + turns) % 4
			dir := directions[thisDir]
			nextPos := posAndDir{x: guard.x + dir.x, y: guard.y + dir.y, d: thisDir}
			if !obstructions[pos{x: nextPos.x, y: nextPos.y}] {
				guard = nextPos
				break
			}
		}
	}

	positions := map[pos]bool{}
	for k := range visited {
		positions[pos{x: k.x, y: k.y}] = true
	}

	return positions, looped
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
