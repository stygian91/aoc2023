package common

import (
	m "aoc2023/utils/math"
	"slices"
)

type Tile = rune

type Board struct {
	Tiles  [][]Tile
	Start  m.Vec2
	Width  int
	Height int
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

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
				board.Start = m.Vec2{X: x, Y: y}
			}
			row = append(row, tile)
		}
		tiles = append(tiles, row)
	}

	board.Tiles = tiles

	return board
}

var validNorth = []Tile{Vertical, SouthToEast, SouthToWest}
var validSouth = []Tile{Vertical, NorthToEast, NorthToWest}
var validEast = []Tile{Horizontal, NorthToWest, SouthToWest}
var validWest = []Tile{Horizontal, NorthToEast, SouthToEast}

func (this Board) IsInBound(pos m.Vec2) bool {
	return m.IsInBound(pos.X, 0, this.Width-1) && m.IsInBound(pos.Y, 0, this.Height-1)
}

func (this Board) IsValid(validTiles []Tile, pos m.Vec2) bool {
	return this.IsInBound(pos) && slices.Contains(validTiles, this.Tiles[pos.Y][pos.X])
}

func (this Board) StartingPipes() []m.Vec2 {
	res := []m.Vec2{}

	northPos := m.Vec2{X: this.Start.X, Y: this.Start.Y - 1}
	if this.IsValid(validNorth, northPos) {
		res = append(res, northPos)
	}

	eastPos := m.Vec2{X: this.Start.X + 1, Y: this.Start.Y}
	if this.IsValid(validEast, eastPos) {
		res = append(res, eastPos)
	}

	southPos := m.Vec2{X: this.Start.X, Y: this.Start.Y + 1}
	if this.IsValid(validSouth, southPos) {
		res = append(res, southPos)
	}

	westPos := m.Vec2{X: this.Start.X - 1, Y: this.Start.Y}
	if this.IsValid(validWest, westPos) {
		res = append(res, westPos)
	}

	return res
}
