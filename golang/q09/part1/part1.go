package part1

import (
	"aoc2023/utils/slices"
	"fmt"
	"strconv"
	"strings"
)

func parse(lines []string) [][]int {
	return slices.Map(func(_ int, line string) []int {
		parts := strings.Fields(line)

		return slices.Map(func(_ int, part string) int {
			res, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}

			return res
		}, parts)
	}, lines)
}

func deriveNext(nums []int) []int {
	if len(nums) <= 1 {
		panic("Nums too short")
	}

	res := []int{}
	for i := 1; i < len(nums); i++ {
		res = append(res, nums[i]-nums[i-1])
	}

	return res
}

func deriveAll(nums []int) [][]int {
	res := [][]int{}
	current := nums

	for {
		res = append(res, current)

		if isAllZero(current) {
			break
		}

		current = deriveNext(current)
	}

	return res
}

func isAllZero(nums []int) bool {
	return slices.All(func(_ int, num int) bool { return num == 0 }, nums)
}

func extra(derives [][]int) int {
	currentE := 0

	for i := len(derives) - 2; i >= 0; i-- {
		currentD := derives[i]
		last := currentD[len(currentD)-1]
		currentE += last
	}

	return currentE
}

func Part1(lines []string) {
	nums := parse(lines)
	sum := 0

	for _, entry := range nums {
		sum += extra(deriveAll(entry))
	}

	fmt.Println("Part 1:", sum)
}
