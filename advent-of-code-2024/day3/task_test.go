package day3

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 161, part1(`xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`))
	require.Equal(t, 161085926, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 48, part2(`xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`))
	require.Equal(t, 82045421, part2(utils.GetInput(t, "input.txt")))
}
