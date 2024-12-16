package day16

import (
	"math"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type pos = utils.Coordinate2D

type posAndDir struct {
	pos pos
	dir pos
}

type scored struct {
	pos   pos
	dir   pos
	score int
}

func part1(input string) int {
	reindeerPos, endPos, lines := parse(input)
	score, _ := bfs(lines, reindeerPos, endPos)
	return score
}

func part2(input string) int {
	reindeerPos, endPos, lines := parse(input)
	bfs(lines, reindeerPos, endPos)
	return 0
}

func parse(input string) (utils.Coordinate2D, utils.Coordinate2D, []string) {
	var reindeerPos pos
	var endPos pos

	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, r := range line {
			if r == 'S' {
				reindeerPos = pos{X: x, Y: y}
			} else if r == 'E' {
				endPos = pos{X: x, Y: y}
			}
		}
	}
	return reindeerPos, endPos, lines
}

func bfs(lines []string, reindeerPos, endPos pos) (int, map[posAndDir]int) {
	canMove := func(p pos) bool {
		return lines[p.Y][p.X] != '#'
	}

	bestScore := math.MaxInt
	reindeer := posAndDir{pos: reindeerPos, dir: pos{X: 1, Y: 0}}
	closed := map[posAndDir]int{reindeer: 0}
	open := []scored{{pos: reindeer.pos, dir: reindeer.dir, score: 0}}
	for len(open) > 0 {
		e := open[0]
		open = open[1:]

		if e.pos == endPos && e.score < bestScore {
			bestScore = e.score
		}

		for _, d := range utils.Get2DDirections() {
			n := e.pos.Add(d)
			if canMove(n) {
				nextScore := e.score + 1
				if d != e.dir {
					nextScore += 1000
				}

				prevScore, exists := closed[posAndDir{pos: n, dir: d}]
				if !exists || prevScore >= nextScore {
					closed[posAndDir{pos: n, dir: d}] = nextScore
					open = append(open, scored{pos: n, dir: d, score: nextScore})
				}
			}
		}
	}
	return bestScore, closed
}
