package day1

import (
	"slices"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	l1, l2 := parseLists(input)

	d := 0
	for i := range l1 {
		d += abs(l1[i] - l2[i])
	}

	return d
}

func part2(input string) int {
	l1, l2 := parseLists(input)

	occ := map[int]int{}
	for _, n := range l2 {
		occ[n] += 1
	}

	s := 0
	for _, n := range l1 {
		s += n * occ[n]
	}
	return s
}

func parseLists(input string) ([]int, []int) {
	l1, l2 := []int{}, []int{}
	for i, n := range utils.GetAllNumbers(input) {
		if i%2 == 0 {
			l1 = append(l1, n)
		} else {
			l2 = append(l2, n)
		}
	}
	slices.Sort(l1)
	slices.Sort(l2)
	return l1, l2
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
