package common

import (
	m "aoc2023/utils/math"
	"strconv"
)

type PartNumber struct {
	Value       int
	NumberStart m.Vec2
	NumberEnd   m.Vec2
}

type Schematic struct {
	Symbols map[m.Vec2]string
	Numbers []PartNumber
}

type Neighbour struct {
	Pos m.Vec2
	Val string
}

func addNum(num *string, partNum *PartNumber, result *Schematic) {
	val, err := strconv.Atoi(*num)
	if err != nil {
		panic(err)
	}
	partNum.Value = val
	result.Numbers = append(result.Numbers, *partNum)
	*num = ""
	*partNum = PartNumber{}
}

func Parse(lines []string) Schematic {
	result := Schematic{Symbols: map[m.Vec2]string{}}
	width := len(lines[0])

	for y, line := range lines {
		partNum := PartNumber{}
		num := ""

		for x, ch := range line {
			if m.IsDigit(ch) {
				if len(num) == 0 {
					partNum.NumberStart = m.Vec2{X: x, Y: y}
				}
				num += string(ch)
			} else {
				if ch != '.' {
					symPos := m.Vec2{X: x, Y: y}
					result.Symbols[symPos] = string(ch)
				}

				if len(num) > 0 {
					partNum.NumberEnd = m.Vec2{X: x - 1, Y: y}
					addNum(&num, &partNum, &result)
				}
			}

			if x == width-1 && len(num) > 0 {
				partNum.NumberEnd = m.Vec2{X: width - 1, Y: y}
				addNum(&num, &partNum, &result)
			}
		}
	}

	return result
}

func (this Schematic) NeighborSymbols(number PartNumber) []Neighbour {
	startX := number.NumberStart.X - 1
	startY := number.NumberStart.Y - 1
	endX := number.NumberEnd.X + 1
	endY := number.NumberEnd.Y + 1

	result := []Neighbour{}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if x >= number.NumberStart.X && x <= number.NumberEnd.X && y >= number.NumberStart.Y && y <= number.NumberEnd.Y {
				continue
			}

			pos := m.Vec2{X: x, Y: y}
			val, exists := this.Symbols[pos]
			if exists {
				result = append(result, Neighbour{Pos: pos, Val: val})
			}
		}
	}

	return result
}
