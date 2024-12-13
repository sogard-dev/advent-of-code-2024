package day13

import (
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

func part1(input string) int {
	return solve(input, 0)
}

func part2(input string) int {
	return solve(input, 10000000000000)

}

func solve(input string, add int) int {
	blocks := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n\n")

	tokens := 0
	for _, b := range blocks {
		nums := utils.GetAllNumbers(b)
		tokens += solveMachine(nums, add)
	}

	return tokens
}

func solveMachine(nums []int, add int) int {
	targetX := nums[4] + add
	targetY := nums[5] + add
	ax := nums[0]
	ay := nums[1]
	bx := nums[2]
	by := nums[3]

	ci := 0
	cj := min(targetX/bx, targetY/by)

	calcX := func() int {
		return ci*ax + cj*bx
	}
	calcY := func() int {
		return ci*ay + cj*by
	}

	for {
		cx := calcX()
		cy := calcY()
		dx := cx - targetX
		dy := cy - targetY
		if dx == 0 && dy == 0 {
			return ci*3 + cj
		}

		rmJ := -1
		adI := -1

		if dx > 0 {
			rmJ = max(rmJ, dx/bx)
		}
		if dy > 0 {
			rmJ = max(rmJ, dy/by)
		}
		if rmJ >= 0 {
			cj -= max(rmJ, 1)
		}

		if dx < 0 {
			adI = max(adI, -dx/ax)
		}
		if dy < 0 {
			adI = max(adI, -dy/ay)
		}
		if adI >= 0 {
			ci += max(adI, 1)
		}

		if cj < 0 {
			return 0
		}
	}
}
