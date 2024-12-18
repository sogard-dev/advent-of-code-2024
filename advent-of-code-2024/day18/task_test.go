package day18

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 22, part1(testInput, 12, 6, 6))
	require.Equal(t, 0, part1(utils.GetInput(t, "input.txt"), 1024, 70, 70))
}

func TestPart2(t *testing.T) {
	require.Equal(t, []int{6, 1}, part2(testInput, 6, 6))
	require.Equal(t, []int{46, 28}, part2(utils.GetInput(t, "input.txt"), 70, 70))
}

const testInput = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
