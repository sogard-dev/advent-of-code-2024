package day2

import (
	"slices"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func solve(input string, allow bool) int {
	safe := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.GetAllNumbers(line)
		if isSafe(nums) {
			safe++
		} else if allow {
			for idxToRemove := range nums {
				a := []int{}
				for k := range nums {
					if k != idxToRemove {
						a = append(a, nums[k])
					}
				}
				if isSafe(a) {
					safe++
					break
				}
			}
		}

	}
	return safe
}

func isSafe(nums []int) bool {
	if nums[0] > nums[1] {
		slices.Reverse(nums)
	}

	for i := 0; i < len(nums)-1; i++ {
		d := nums[i+1] - nums[i]

		if d < 1 || d > 3 {
			return false
		}
	}
	return true
}
