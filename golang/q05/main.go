package main

import (
	"aoc2023/utils/files"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type MapEntry struct {
	Src, Dest, Len int
}

func parseMap(group []string) (string, []MapEntry) {
	name := strings.Split(group[0], " ")[0]
	entries := []MapEntry{}

	for i := 1; i < len(group); i++ {
		line := group[i]
		entries = append(entries, parseEntry(line))
	}

	return name, entries
}

func parseEntry(line string) MapEntry {
	partsStr := strings.Split(line, " ")
	parts := []int{}

	for _, partStr := range partsStr {
		part, err := strconv.Atoi(partStr)
		if err != nil {
			panic(err)
		}

		parts = append(parts, part)
	}

	return MapEntry{
		Src:  parts[1],
		Dest: parts[0],
		Len:  parts[2],
	}
}

func parseSeeds(group []string) []int {
	result := []int{}
	line := group[0]
	parts := strings.Split(line, ": ")

	for _, seedStr := range strings.Split(parts[1], " ") {
		num, err := strconv.Atoi(seedStr)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}

	return result
}

func lookup(src int, entries []MapEntry) int {
	for _, entry := range entries {
		if src >= entry.Src && src < entry.Src+entry.Len {
			diff := src - entry.Src
			return entry.Dest + diff
		}
	}

	return src
}

func seedToLocation(seed int, maps map[string][]MapEntry) int {
	soil := lookup(seed, maps["seed-to-soil"])
	fertilizer := lookup(soil, maps["soil-to-fertilizer"])
	water := lookup(fertilizer, maps["fertilizer-to-water"])
	light := lookup(water, maps["water-to-light"])
	temp := lookup(light, maps["light-to-temperature"])
	humid := lookup(temp, maps["temperature-to-humidity"])
	return lookup(humid, maps["humidity-to-location"])
}

func main() {
	// lines, err := files.ReadLines("./data/demo.txt")
	lines, err := files.ReadLines("./data/input.txt")
	if err != nil {
		panic(err)
	}

	groups := [][]string{}
	current := []string{}

	for i, line := range lines {
		if len(line) == 0 || i == len(lines)-1 {
			groups = append(groups, current)
			current = []string{}
			continue
		}

		current = append(current, line)
	}

	seeds := parseSeeds(groups[0])
	maps := map[string][]MapEntry{}

	for i := 1; i < len(groups); i++ {
		group := groups[i]
		name, entries := parseMap(group)
		maps[name] = entries
	}

	seedLoc := map[int]int{}
	minLoc := math.MaxInt
	for _, seed := range seeds {
		loc := seedToLocation(seed, maps)
		seedLoc[seed] = loc

		if minLoc > loc {
			minLoc = loc
		}
	}

	fmt.Println("Part 1:", minLoc)
	// ---------------------------
	// Part 2:
	// ---------------------------
	{
		seedEntries := [][]int{}
		chunk := []int{}
		minLoc := math.MaxInt

		for i, v := range seeds {
			chunk = append(chunk, v)

			if i%2 == 1 {
				seedEntries = append(seedEntries, chunk)
				chunk = []int{}
			}
		}

		for _, seedEntry := range seedEntries {
			for seed := seedEntry[0]; seed < seedEntry[0]+seedEntry[1]; seed++ {
				loc := seedToLocation(seed, maps)

				if minLoc > loc {
					minLoc = loc
				}
			}
		}

		fmt.Println("Part 2:", minLoc)
	}
}
