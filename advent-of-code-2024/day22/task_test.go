package day22

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 37327623, part1(`1
10
100
2024`))
	require.Equal(t, 16953639210, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 23, part2(`1
2
3
2024`))
	require.Equal(t, 1863, part2(utils.GetInput(t, "input.txt")))
}
