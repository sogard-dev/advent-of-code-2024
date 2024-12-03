package day3

import (
	"regexp"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3}\))`)

	sum := 0
	for _, m := range r.FindAllString(input, -1) {
		nums := utils.GetAllNumbers(m)
		sum += nums[0] * nums[1]
	}

	return sum
}

func part2(input string) int {
	do, _ := regexp.Compile(`do\(\)`)
	dont, _ := regexp.Compile(`don't\(\)`)

	enableIdx := do.FindAllStringIndex(input, -1)
	disableIdx := dont.FindAllStringIndex(input, -1)
	enabled := map[int]bool{}

	for _, v := range enableIdx {
		enabled[v[0]] = true
	}

	for _, v := range disableIdx {
		for k := v[0]; k < len(input); k++ {
			if _, exists := enabled[k]; !exists {
				enabled[k] = false
			} else if enabled[k] {
				break
			}
		}
	}

	sum := 0
	mul, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3}\))`)
	for _, m := range mul.FindAllStringIndex(input, -1) {
		enabled, exists := enabled[m[0]]
		t := input[m[0]:m[1]]
		if !exists || enabled {
			nums := utils.GetAllNumbers(t)
			sum += nums[0] * nums[1]
		}
	}

	return sum
}
