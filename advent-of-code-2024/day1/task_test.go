package day1

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

const testInput = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	require.Equal(t, 11, part1(testInput))
	require.Equal(t, 1603498, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 31, part2(testInput))
	require.Equal(t, 0, part2(utils.GetInput(t, "input.txt")))
}
