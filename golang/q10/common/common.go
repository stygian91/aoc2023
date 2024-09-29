package common

import (
	m "aoc2023/utils/math"
	"fmt"
	"slices"
)

type Tile = rune

type Board struct {
	Tiles  [][]Tile
	Start  m.Vec2
	Width  int
	Height int
}

type TilePos struct {
	Pos  m.Vec2
	Tile Tile
}

func (this TilePos) String() string {
	return fmt.Sprintf("{X: %d, Y: %d, Tile: %c}", this.Pos.X, this.Pos.Y, this.Tile)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

const (
	NilTile     Tile = ' '
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

func (this Board) GetTile(pos m.Vec2) (Tile, bool) {
	if !this.IsInBound(pos) {
		return NilTile, false
	}

	return this.Tiles[pos.Y][pos.X], true
}

func (this Board) GetNeighbours(pos m.Vec2) []TilePos {
	res := []TilePos{}

	tile, exists := this.GetTile(pos)
	if !exists {
		return res
	}

	offsets := GetNeighbourOffsetsByTile(tile)

	for _, offset := range offsets {
		nPos := pos.Add(offset)
		nTile, exists := this.GetTile(nPos)
		if !exists {
			continue
		}
		res = append(res, TilePos{Pos: nPos, Tile: nTile})
	}

	return res
}

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

func GetNeighbourOffsetsByTile(tile Tile) []m.Vec2 {
	switch tile {
	case Vertical:
		return []m.Vec2{{X: 0, Y: -1}, {X: 0, Y: 1}}
	case Horizontal:
		return []m.Vec2{{X: 1, Y: 0}, {X: -1, Y: 0}}
	case NorthToEast:
		return []m.Vec2{{X: 0, Y: -1}, {X: 1, Y: 0}}
	case NorthToWest:
		return []m.Vec2{{X: 0, Y: -1}, {X: -1, Y: 0}}
	case SouthToEast:
		return []m.Vec2{{X: 0, Y: 1}, {X: 1, Y: 0}}
	case SouthToWest:
		return []m.Vec2{{X: 0, Y: 1}, {X: -1, Y: 0}}
	case Ground:
		return []m.Vec2{}
	case Start:
		return []m.Vec2{}
	case NilTile:
		panic("GetNeighbour of NilTile")

	default:
		panic("Unknown tile type")
	}
}
