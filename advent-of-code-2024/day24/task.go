package day24

import (
	"slices"
	"strconv"
	"strings"

	"github.com/sogard-dev/advent-of-code-2024/utils"
)

type gate struct {
	in1, in2, out string
	gateType      string
}

func (g *gate) toString() string {
	return g.in1 + "_" + g.gateType + "_" + g.in2 + "_" + g.out
}

func part1(input string) int {
	blocks := strings.Split(input, "\n\n")
	wires := parseWires(blocks)
	gates := parseGates(blocks)

	for _, v := range gates {
		v.evaluate(wires, gates)
	}

	return getNumberFromWires(wires, 'z')
}

func ui(input string) {
	blocks := strings.Split(input, "\n\n")
	gates := parseGates(blocks)

	for _, g := range gates {
		if g.gateType == "XOR" {
			println(g.toString(), `[color="lightgreen",style="filled"];`)
		} else if g.gateType == "OR" {
			println(g.toString(), `[color="blue",style="filled"];`)
		} else if g.gateType == "AND" {
			println(g.toString(), `[color="red",style="filled"];`)
		}

		if gIn, e := gates[g.in1]; e {
			println(gIn.toString(), "->", g.toString())
		}

		if gIn, e := gates[g.in2]; e {
			println(gIn.toString(), "->", g.toString())
		}

	}
}

func part2(input string) string {
	blocks := strings.Split(input, "\n\n")
	wires := parseWires(blocks)
	gates := parseGates(blocks)
	evaluateGates(gates, wires)

	//Something wrong with:
	//x06 XOR y06 -> dbp SWAP y06 AND x06 -> fdv
	//rqt OR rdt -> z23  SWAP nsr XOR gsd -> kdf
	//qbw XOR fqf -> ckj SWAP x15 AND y15 -> z15
	//vbt AND vqr -> z39 SWAP vqr XOR vbt -> rpp

	z := getNumberFromWires(wires, 'z')
	a := strconv.FormatInt(int64(z), 2)

	return a
}

func evaluateGates(gates map[string]gate, wires map[string]bool) {
	for _, v := range gates {
		v.evaluate(wires, gates)
	}
}

func (g gate) evaluate(wires map[string]bool, gates map[string]gate) {
	if _, e := wires[g.in1]; !e {
		gates[g.in1].evaluate(wires, gates)
	}
	if _, e := wires[g.in2]; !e {
		gates[g.in2].evaluate(wires, gates)
	}

	if g.gateType == "AND" {
		wires[g.out] = wires[g.in1] && wires[g.in2]
	} else if g.gateType == "XOR" {
		wires[g.out] = wires[g.in1] != wires[g.in2]
	} else if g.gateType == "OR" {
		wires[g.out] = wires[g.in1] || wires[g.in2]
	} else {
		panic("Ohh noes")
	}
}

func getNumberFromWires(wires map[string]bool, prefix byte) int {
	keys := []string{}
	for k := range wires {
		if k[0] == prefix {
			keys = append(keys, k)
		}
	}
	slices.Sort(keys)

	bin := ""
	for _, k := range keys {
		v := wires[k]
		if v {
			bin = "1" + bin
		} else {
			bin = "0" + bin
		}
	}

	i, _ := strconv.ParseInt(bin, 2, 64)

	return int(i)
}

func parseGates(blocks []string) map[string]gate {
	gates := map[string]gate{}
	for _, line := range strings.Split(blocks[1], "\n") {
		spl := strings.Split(line, " ")
		g := gate{in1: spl[0], in2: spl[2], gateType: spl[1], out: spl[4]}
		gates[spl[4]] = g
	}
	return gates
}

func parseWires(blocks []string) map[string]bool {
	wires := map[string]bool{}
	for _, line := range strings.Split(blocks[0], "\n") {
		spl := strings.Split(line, ": ")
		wires[spl[0]] = utils.GetAllNumbers(spl[1])[0] == 1
	}
	return wires
}
