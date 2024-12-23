package day20

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type pos = utils.Coordinate2D

func part1(input string, saveAtleast int) int {
	return solve(input, saveAtleast, 2)
}

func part2(input string, saveAtleast int) int {
	return solve(input, saveAtleast, 20)
}

func solve(input string, saveAtleast, cheatDistance int) int {
	lines := strings.Split(input, "\n")
	remaining := map[pos]int{}
	for y, line := range lines {
		for x, r := range line {
			if r == 'E' {
				fillRemaining(lines, pos{X: x, Y: y}, remaining, 0)
			}
		}
	}

	s := 0
	for ap1, ad1 := range remaining {
		for ap2, ad2 := range remaining {
			cheatSteps := ap1.Manhatten(ap2)
			if ad2 < ad1 && cheatSteps <= cheatDistance {
				saved := ad1 - ad2 - cheatSteps
				if saved >= saveAtleast {
					s += 1
				}
			}
		}
	}

	return s
}

func fillRemaining(lines []string, p pos, remaining map[utils.Coordinate2D]int, dist int) {
	remaining[p] = dist
	for _, d := range utils.Get2DDirections() {
		np := d.Add(p)
		if _, e := remaining[np]; !e && lines[np.Y][np.X] != '#' {
			fillRemaining(lines, np, remaining, dist+1)
		}
	}
}
