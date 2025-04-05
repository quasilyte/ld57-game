package dat

import "github.com/quasilyte/gmath"

type CellPos struct {
	X int
	Y int
}

type Tile int

const (
	TileGrass Tile = iota
	TileForest
	TileSwamp
	TileVoid
)

func (p CellPos) Add(other CellPos) CellPos {
	return CellPos{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p CellPos) IsZero() bool {
	return p == CellPos{}
}

func (p CellPos) ToVecPos(center bool) gmath.Vec {
	pos := gmath.Vec{
		X: float64(p.X * 32),
		Y: float64(p.Y * 32),
	}
	if center {
		pos.X += 16
		pos.Y += 16
	}
	return pos
}

type Map struct {
	Width  int
	Height int
	Units  []DeployedUnit
	Tiles  [][]Tile
}

type DeployedUnit struct {
	Team int
	Pos  CellPos
	Unit *Unit
}
