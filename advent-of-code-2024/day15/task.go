package day15

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type pos = utils.Coordinate2D

func part1(input string) int {
	blocks := parseBlocks(input)
	m, robot := parseMap(blocks)
	runInstructions(blocks, m, robot)
	return calculateSumOfGps(m)
}

func part2(input string) int {
	input = strings.ReplaceAll(input, "#", "##")
	input = strings.ReplaceAll(input, "O", "[]")
	input = strings.ReplaceAll(input, ".", "..")
	input = strings.ReplaceAll(input, "@", "@.")
	return part1(input)
}

func calculateSumOfGps(m map[pos]rune) int {
	width := 0
	height := 0
	for p := range m {
		width = max(width, p.X+1)
		height = max(height, p.Y+1)
	}

	check := 0
	for p, v := range m {
		if v == 'O' || v == '[' {
			check += p.X + p.Y*100
		}
	}
	return check
}

func runInstructions(blocks []string, m map[pos]rune, robot pos) {
	for _, line := range strings.Split(blocks[1], "\n") {
		for _, op := range line {
			d := dirToPos(op)
			if canMove(m, robot, d) {
				move(m, robot, d, map[pos]bool{})
				robot = robot.Add(d)
			}
		}
	}
}

func parseMap(blocks []string) (map[pos]rune, pos) {
	m := map[pos]rune{}
	var robot pos
	lines := strings.Split(blocks[0], "\n")

	for y, line := range lines {
		for x, r := range line {
			p := pos{X: x, Y: y}
			if r == '@' {
				robot = p
			}
			m[p] = r

		}
	}
	return m, robot
}

func parseBlocks(input string) []string {
	input = strings.ReplaceAll(input, "\r", "")
	blocks := strings.Split(input, "\n\n")
	return blocks
}

func move(m map[pos]rune, obj, d pos, visited map[pos]bool) {
	r := m[obj]
	if r == '.' || visited[obj] {
		return
	}
	visited[obj] = true

	next := obj.Add(d)
	move(m, next, d, visited)

	if d.X == 0 {
		if r == '[' {
			move(m, obj.Add(pos{X: 1, Y: 0}), d, visited)
		} else if r == ']' {
			move(m, obj.Add(pos{X: -1, Y: 0}), d, visited)
		}
	}

	m[next] = m[obj]
	m[obj] = '.'
}

func canMove(m map[pos]rune, obj pos, d pos) bool {
	r := m[obj]
	if r == '.' {
		return true
	}
	if r == '#' {
		return false
	}

	n := obj.Add(d)

	if r == '@' || r == 'O' || d.Y == 0 {
		return canMove(m, n, d)
	}
	if r == '[' {
		return canMove(m, n, d) && canMove(m, n.Add(pos{X: 1, Y: 0}), d)
	}
	if r == ']' {
		return canMove(m, n, d) && canMove(m, n.Add(pos{X: -1, Y: 0}), d)
	}

	panic("donkey, missing: " + string(r))
}

func dirToPos(r rune) pos {
	if r == '^' {
		return pos{X: 0, Y: -1}
	}
	if r == 'v' {
		return pos{X: 0, Y: 1}
	}
	if r == '<' {
		return pos{X: -1, Y: 0}
	}
	if r == '>' {
		return pos{X: 1, Y: 0}
	}
	panic("donkey")
}
