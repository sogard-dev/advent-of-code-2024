package day17

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0}, part1(`Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`))
	require.Equal(t, []int{1, 5, 3, 0, 2, 5, 2, 5, 3}, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 108107566389757, part2(utils.GetInput(t, "input.txt")))
}
