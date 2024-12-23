package day22

import (
	"fmt"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	s := 0
	for _, n := range utils.GetAllNumbers(input) {
		for range 2000 {
			n = nextSecretNumber(n)
		}

		s += n
	}

	return s
}

func part2(input string) int {
	changes := map[string]int{}

	for _, n := range utils.GetAllNumbers(input) {
		bananas := map[int]int{}
		for i := range 2000 {
			n = nextSecretNumber(n)
			bananas[i] = n % 10
		}

		thisChanges := map[string]int{}
		for i := 4; i < len(bananas); i++ {
			c1 := bananas[i-1] - bananas[i-0]
			c2 := bananas[i-2] - bananas[i-1]
			c3 := bananas[i-3] - bananas[i-2]
			c4 := bananas[i-4] - bananas[i-3]
			key := fmt.Sprintf("%d,%d,%d,%d", c4, c3, c2, c1)
			if _, e := thisChanges[key]; !e {
				thisChanges[key] = bananas[i]
			}
		}

		for k, v := range thisChanges {
			changes[k] += v
		}
	}

	highest := 0
	for _, v := range changes {
		highest = max(highest, v)
	}

	return highest
}

func nextSecretNumber(n int) int {
	n = mix(n, n*64)
	n = prune(n)

	n = mix(n, n/32)
	n = prune(n)

	n = mix(n, n*2048)
	n = prune(n)
	return n
}

func mix(n, r int) int {
	return r ^ n
}

func prune(n int) int {
	return n % 16777216
}
