package common

import (
	"aoc2023/utils/slices"
	"strconv"
	"strings"
)

func Parse(lines []string) [][]int {
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

func DeriveNext(nums []int) []int {
	if len(nums) <= 1 {
		panic("Nums too short")
	}

	res := []int{}
	for i := 1; i < len(nums); i++ {
		res = append(res, nums[i]-nums[i-1])
	}

	return res
}

func DeriveAll(nums []int) [][]int {
	res := [][]int{}
	current := nums

	for {
		res = append(res, current)

		if IsAllZero(current) {
			break
		}

		current = DeriveNext(current)
	}

	return res
}

func IsAllZero(nums []int) bool {
	return slices.All(func(_ int, num int) bool { return num == 0 }, nums)
}

