package day20

import (
	"testing"

	"github.com/sogard-dev/advent-of-code-2024/utils"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	require.Equal(t, 2, part1(testInput, 40))
	require.Equal(t, 1404, part1(utils.GetInput(t, "input.txt"), 100))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 7, part2(testInput, 74))
	require.Equal(t, 1010981, part2(utils.GetInput(t, "input.txt"), 100))
}

const testInput = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
