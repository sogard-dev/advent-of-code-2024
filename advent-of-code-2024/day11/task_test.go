package day11

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 55312, part1(`125 17`))
	require.Equal(t, 175006, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 207961583799296, part2(utils.GetInput(t, "input.txt")))
}
