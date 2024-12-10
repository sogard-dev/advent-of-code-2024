package day10

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 4, part1(`..90..9
...1.98
...2..7
6543456
765.987
876....
987....`))
	require.Equal(t, 841, part1(utils.GetInput(t, "input.txt")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 81, part2(`89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`))
	require.Equal(t, 1875, part2(utils.GetInput(t, "input.txt")))
}
