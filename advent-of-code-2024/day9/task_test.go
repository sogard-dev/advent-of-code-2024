package day9

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 1928, part1(testInput))
	require.Equal(t, 6471961544878, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 2858, part2(testInput))
	require.Equal(t, 6511178035564, part2(utils.GetInput(t, "input.txt")))
}

const testInput = `2333133121414131402`
