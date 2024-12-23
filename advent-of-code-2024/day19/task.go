package day19

import (
	"strings"
)

func part1(input string) int {
	available, lines := parse(input)
	s := 0
	for _, line := range lines {
		if canMake(available, line) > 0 {
			s += 1
		}
	}
	return s
}

func part2(input string) int {
	available, lines := parse(input)
	s := 0
	for _, line := range lines {
		s += canMake(available, line)
	}
	return s
}

func parse(input string) ([]string, []string) {
	blocks := strings.Split(input, "\n\n")
	available := strings.Split(blocks[0], ", ")
	lines := strings.Split(blocks[1], "\n")
	return available, lines
}

func canMake(available []string, line string) int {
	cache := map[int]int{}
	var check func(int) int
	check = func(idx int) int {
		if idx == len(line) {
			return 1
		}
		if v, e := cache[idx]; e {
			return v
		}

		canDo := 0
		for _, a := range available {
			la := len(a)
			if la+idx <= len(line) && line[idx:idx+la] == a {
				canDo += check(idx + la)
			}
		}

		cache[idx] = canDo

		return canDo
	}

	return check(0)
}
