package day8

import (
	"strings"
)

type pos struct {
	x, y int
}

func part1(input string) int {
	return solve(input, 1, 2)
}

func part2(input string) int {
	return solve(input, 0, 1000)
}

func solve(input string, mStart, mEnd int) int {
	antennas := map[rune][]pos{}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, r := range line {
			if r != '.' {
				if _, exists := antennas[r]; !exists {
					antennas[r] = []pos{}
				}
				antennas[r] = append(antennas[r], pos{x: x, y: y})
			}
		}
	}

	height := len(lines)
	width := len(lines[0])
	antinodes := map[pos]bool{}

	for _, list := range antennas {
		for _, a := range list {
			for _, b := range list {
				if a != b {
					for m := mStart; m < mEnd; m++ {
						nx := (b.x-a.x)*m + b.x
						ny := (b.y-a.y)*m + b.y
						if nx >= 0 && nx < width && ny >= 0 && ny < height {
							antinodes[pos{x: nx, y: ny}] = true
						} else {
							break
						}
					}
				}
			}
		}
	}

	return len(antinodes)
}
