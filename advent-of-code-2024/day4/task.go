package day4

import (
	"strings"
)

type coord struct {
	x, y int
}
type subword struct {
	dx, dy int
	r      byte
}

type word struct {
	s []subword
}

func wordInDirection(dx, dy int, w string) word {
	l := []subword{}

	x, y := 0, 0
	for i := range w {
		l = append(l, subword{
			r:  w[i],
			dx: x,
			dy: y,
		})
		x += dx
		y += dy
	}

	return word{s: l}
}

func part1(input string) int {
	toFind := []word{
		wordInDirection(1, 0, "XMAS"),
		wordInDirection(-1, 0, "XMAS"),
		wordInDirection(0, 1, "XMAS"),
		wordInDirection(0, -1, "XMAS"),
		wordInDirection(1, 1, "XMAS"),
		wordInDirection(1, -1, "XMAS"),
		wordInDirection(-1, 1, "XMAS"),
		wordInDirection(-1, -1, "XMAS"),
	}

	hits := 0
	solve(input, toFind, func(_ coord) {
		hits++
	})

	return hits
}

func part2(input string) int {
	toFind := []word{
		wordInDirection(1, 1, "MAS"),
		wordInDirection(1, -1, "MAS"),
		wordInDirection(-1, 1, "MAS"),
		wordInDirection(-1, -1, "MAS"),
	}

	centerHits := map[coord]int{}
	solve(input, toFind, func(c coord) {
		centerHits[c] += 1
	})

	hits := 0
	for _, v := range centerHits {
		if v == 2 {
			hits++
		}
	}
	return hits
}

func solve(input string, toFind []word, match func(coord)) {
	spl := strings.Split(input, "\n")
	for y := range spl {
		for x := range spl[y] {
			for _, f := range toFind {
				matches := true
				for _, r := range f.s {
					nx := x + r.dx
					ny := y + r.dy
					xo := nx < 0 || nx >= len(spl[y])
					yo := ny < 0 || ny >= len(spl)
					if xo || yo || spl[ny][nx] != r.r {
						matches = false
						break
					}
				}

				if matches {
					match(coord{
						x: x + f.s[1].dx,
						y: y + f.s[1].dy,
					})
				}
			}
		}
	}
}
