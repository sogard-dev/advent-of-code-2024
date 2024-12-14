package day14

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 12, part1(`p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`, 11, 7))
	require.Equal(t, 225943500, part1(utils.GetInput(t, "input.txt"), 101, 103))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 6377, part2(utils.GetInput(t, "input.txt"), 101, 103))
}
