package day5

import (
	"slices"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type rule struct {
	a, b int
}

func part1(input string) int {
	return solve(input, false)
}

func part2(input string) int {
	return solve(input, true)
}

func fixOrdering(rules []rule, pages []int) {
	changed := true
	for changed {
		changed = false
		for _, rule := range rules {
			aIdx := slices.Index(pages, rule.a)
			bIdx := slices.Index(pages, rule.b)
			if aIdx != -1 && bIdx != -1 && bIdx < aIdx {
				pages[aIdx], pages[bIdx] = pages[bIdx], pages[aIdx]
				changed = true
			}
		}
	}
}

func solve(input string, part2 bool) int {
	blocks := strings.Split(input, "\n\n")
	rulesNumbers := utils.GetAllNumbers(blocks[0])
	rules := []rule{}
	for i := 0; i < len(rulesNumbers); i += 2 {
		rules = append(rules, rule{
			a: rulesNumbers[i],
			b: rulesNumbers[i+1],
		})
	}

	sum := 0
	for _, line := range strings.Split(blocks[1], "\n") {
		pages := utils.GetAllNumbers(line)
		initiallyValid := checkIsValid(rules, pages)
		if !part2 {
			if initiallyValid {
				middlePage := pages[len(pages)/2]
				sum += middlePage
			}
		} else {
			if !initiallyValid {
				fixOrdering(rules, pages)
				middlePage := pages[len(pages)/2]
				sum += middlePage
			}
		}
	}

	return sum
}

func checkIsValid(rules []rule, pages []int) bool {
	valid := true
	for _, rule := range rules {
		aIdx := slices.Index(pages, rule.a)
		bIdx := slices.Index(pages, rule.b)
		if aIdx != -1 && bIdx != -1 && bIdx < aIdx {
			valid = false
			break
		}
	}
	return valid
}
