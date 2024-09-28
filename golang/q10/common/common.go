package common

import "aoc2023/utils/math"

type Tile = rune

type Board struct {
	Tiles  [][]Tile
	Start  math.Vec2
	Width  int
	Height int
}

const (
	Vertical    Tile = '|'
	Horizontal  Tile = '-'
	NorthToEast Tile = 'L'
	NorthToWest Tile = 'J'
	SouthToEast Tile = 'F'
	SouthToWest Tile = '7'
	Ground      Tile = '.'
	Start       Tile = 'S'
)

func TileFromRune(ch rune) Tile {
	switch ch {
	case Vertical:
		return Vertical
	case Horizontal:
		return Horizontal
	case NorthToEast:
		return NorthToEast
	case NorthToWest:
		return NorthToWest
	case SouthToEast:
		return SouthToEast
	case SouthToWest:
		return SouthToWest
	case Ground:
		return Ground
	case Start:
		return Start
	default:
		panic("Invalid tile")
	}
}

func Parse(lines []string) Board {
	board := Board{
		Height: len(lines),
		Width:  len(lines[0]),
	}
	tiles := [][]Tile{}

	for y, line := range lines {
		row := []Tile{}
		for x, ch := range line {
			tile := TileFromRune(ch)
			if tile == Start {
				board.Start = math.Vec2{X: x, Y: y}
			}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	board.Tiles = tiles

	return board
}
