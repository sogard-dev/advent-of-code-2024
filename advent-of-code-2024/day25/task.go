package day25

import (
	"fmt"
	"strings"
)

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")

	keys := [][5]int{}
	locks := [][5]int{}
	for _, block := range blocks {
		lines := strings.Split(block, "\n")

		teeth := [5]int{}
		for x := range 5 {
			t := 0
			for y := 1; y < 6; y++ {
				if lines[y][x] == '#' {
					t += 1
				}
			}
			teeth[x] = t
		}

		if lines[0][0] == '#' {
			locks = append(locks, teeth)
		} else {
			keys = append(keys, teeth)
		}
	}

	combis := 0
	for _, k := range keys {
		for _, l := range locks {
			fits := true
			for i := range len(k) {
				if k[i]+l[i] > 5 {
					fits = false
					break
				}
			}

			if fits {
				combis += 1
			}
		}
	}

	return combis
}

func part2(input string) int {
	for _, line := range strings.Split(input, "\n") {
		fmt.Println(line)
	}
	return 0
}
