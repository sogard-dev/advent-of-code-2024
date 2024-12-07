package day7

import (
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func solve(input string, withConcat bool) int {
	res := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetAllNumbers(line)
		sum := nums[0]
		values := nums[1:]

		if canBeTrue(sum, values, 0, withConcat) {
			res += sum
		}
	}
	return res
}

func canBeTrue(sum int, values []int, subsum int, withConcat bool) bool {
	if len(values) == 0 {
		return sum == subsum
	} else {
		c := canBeTrue(sum, values[1:], subsum+values[0], withConcat)
		c = c || canBeTrue(sum, values[1:], subsum*values[0], withConcat)
		if withConcat {
			c = c || canBeTrue(sum, values[1:], concat(subsum, values[0]), withConcat)
		}
		return c
	}
}

func concat(cur, next int) int {
	a, _ := strconv.Atoi(strconv.Itoa(cur) + strconv.Itoa(next))
	return a
}
