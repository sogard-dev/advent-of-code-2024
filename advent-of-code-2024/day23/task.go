package day23

import (
	"slices"
	"strings"
)

func part1(input string) int {
	connections := parse(input)
	clusters := findSetsOfThree(connections)

	sum := 0
	for _, v := range clusters {
		hasT := false
		for _, l := range v {
			if l[0] == 't' {
				hasT = true
				break
			}
		}

		if hasT {
			sum += 1
		}
	}

	return sum
}

func part2(input string) string {
	connections := parse(input)

	longest := ""

	for initK := range connections {
		currentCluster := map[string]bool{initK: true}

		changed := true
		for changed {
			changed = false
			for k := range currentCluster {
				for _, otherK := range connections[k] {
					if !currentCluster[otherK] {
						connectedToAll := true

						for testK := range currentCluster {
							if !slices.Contains(connections[otherK], testK) {
								connectedToAll = false
								break
							}
						}

						if connectedToAll {
							currentCluster[otherK] = true
							changed = true
						}
					}

				}
			}
		}

		sl := []string{}
		for k := range currentCluster {
			sl = append(sl, k)
		}
		slices.Sort(sl)
		str := strings.Join(sl, ",")
		if len(str) > len(longest) {
			longest = str
		}
	}

	return longest
}

func parse(input string) map[string][]string {
	connections := map[string][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		a, b := line[0:2], line[3:5]
		if _, e := connections[a]; !e {
			connections[a] = []string{}
		}
		if _, e := connections[b]; !e {
			connections[b] = []string{}
		}
		connections[a] = append(connections[a], b)
		connections[b] = append(connections[b], a)
	}
	return connections
}

func findSetsOfThree(m map[string][]string) map[string][]string {
	clusters := map[string][]string{}

	for k1, c := range m {
		for _, k2 := range c {
			for _, k3 := range m[k2] {
				for _, k4 := range m[k3] {
					if k4 == k1 {
						s := []string{k1, k2, k3}
						slices.Sort(s)
						k := strings.Join(s, ",")
						clusters[k] = s
					}
				}
			}
		}
	}

	return clusters
}
