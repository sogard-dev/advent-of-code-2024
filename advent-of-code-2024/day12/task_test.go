package day12

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 140, part1(testInput))
	require.Equal(t, 1461752, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 80, part2(testInput))
	require.Equal(t, 904114, part2(utils.GetInput(t, "input.txt")))
}

const testInput = `AAAA
BBCD
BBCC
EEEC`
