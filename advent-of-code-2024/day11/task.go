package day11

import (
	"strconv"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type entry struct {
	stone  int
	blinks int
}

func part1(input string) int {
	return solve(input, 25)
}

func part2(input string) int {
	return solve(input, 75)
}

func solve(input string, blinks int) int {
	sum := 0
	c := map[entry]int{}
	for _, n := range utils.GetAllNumbers(input) {
		sum += recBlink(n, 0, blinks, c)
	}

	return sum
}

func recBlink(n, blink, blinks int, cache map[entry]int) int {
	if blink == blinks {
		return 1
	}
	stone := entry{
		stone:  n,
		blinks: blink,
	}
	if v, e := cache[stone]; e {
		return v
	}

	s := strconv.Itoa(n)
	v := 0
	if n == 0 {
		v = recBlink(1, blink+1, blinks, cache)
	} else if len(s)%2 == 0 {
		a, b := getHalfs(s)
		v = recBlink(a, blink+1, blinks, cache) + recBlink(b, blink+1, blinks, cache)
	} else {
		v = recBlink(n*2024, blink+1, blinks, cache)
	}
	cache[stone] = v

	return v
}

func getHalfs(s string) (int, int) {
	l := len(s) / 2
	a, _ := strconv.Atoi(s[0:l])
	b, _ := strconv.Atoi(s[l:])
	return a, b
}
