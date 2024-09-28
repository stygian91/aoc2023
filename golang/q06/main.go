package main

import (
	"fmt"
)

type Race struct {
	Duration, Record int
}

func (this Race) CalcDistance(buttonDuration int) int {
	if buttonDuration >= this.Duration || buttonDuration <= 0 {
		return 0
	}

	travelDuration := this.Duration - buttonDuration
	return travelDuration * buttonDuration
}

// func uniqueAdd(set *[]int, item int) *[]int {
// 	if !slices.Contains(*set, item) {
// 		*set = append(*set, item)
// 	}
//
// 	return set
// }

func (this Race) AboveRecord() []int {
	result := []int{}

	for i := 1; i < this.Duration; i++ {
		currentDistance := this.CalcDistance(i)
		if currentDistance > this.Record {
			result = append(result, i)
		}
	}

	return result

	// for {
	// 	currentDuration += offset
	// 	currentDistance = this.CalcDistance(currentDuration)
	//
	// 	if currentDistance < maxDistance || currentDuration > this.Duration {
	// 		result = *uniqueAdd(&result, maxDuration)
	// 		break
	// 	}
	//
	// 	if currentDistance > maxDistance {
	// 		maxDistance = currentDistance
	// 		maxDuration = currentDuration
	// 		continue
	// 	}
	//
	// 	if currentDistance == maxDistance {
	// 		result = *uniqueAdd(&result, maxDuration)
	// 		result = *uniqueAdd(&result, currentDuration)
	// 	}
	// }
}

func getDemoRaces() []Race {
	races := []Race{
		{Duration: 7, Record: 9},
		{Duration: 15, Record: 40},
		{Duration: 30, Record: 200},
	}

	return races
}

func getRacesPart1() []Race {
	races := []Race{
		{Duration: 46, Record: 208},
		{Duration: 85, Record: 1412},
		{Duration: 75, Record: 1257},
		{Duration: 82, Record: 1410},
	}

	return races
}

func getRacePart2() Race {
	return Race{
		Duration: 46857582,
		Record:   208141212571410,
	}
}

func part1() {
	// races := getDemoRaces()
	races := getRacesPart1()
	product := 1

	for _, race := range races {
		aboveRecord := race.AboveRecord()
		if race.CalcDistance(aboveRecord[0]) <= race.Record {
			panic("Found maximum was <= than record")
		}

		product *= len(aboveRecord)
	}

	fmt.Println("Part 1:", product)
}

func part2() {
	above := getRacePart2().AboveRecord()

	fmt.Println("Part 2:", len(above))
}

func main() {
	part1()
	part2()
}
