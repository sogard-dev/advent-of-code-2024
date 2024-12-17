package utils

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
)

type Coordinate2D struct {
	X, Y int
}

func (c Coordinate2D) Add(o Coordinate2D) Coordinate2D {
	return Coordinate2D{c.X + o.X, c.Y + o.Y}
}

func (c Coordinate2D) Opposite() Coordinate2D {
	return Coordinate2D{-c.X, -c.Y}
}

func Get2DDirections() []Coordinate2D {
	return []Coordinate2D{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
}

func GetInput(t *testing.T, file string) string {
	goldenfile := filepath.Join("testdata", file)
	want, err := os.ReadFile(goldenfile)
	if err != nil {
		t.Fatal("error reading golden file:", err)
	}

	return string(want)
}

func GetAllNumbers(input string) []int {
	re := regexp.MustCompile("[-]?[0-9]+")
	numbers := re.FindAllString(input, -1)
	ret := []int{}
	for _, v := range numbers {
		v, err := strconv.Atoi(v)
		if err != nil {
			panic("shit")
		}
		ret = append(ret, v)
	}
	return ret
}

func NextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func GetPerm[T interface{}](orig []T, p []int) []T {
	result := append([]T{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func SetBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

func HasBit(n int, pos int) bool {
	val := n & (1 << pos)
	return (val > 0)
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}
