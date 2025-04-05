package mapgen

import (
	"math"

	"github.com/quasilyte/gmath"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
)

type EnemyKind int

const (
	EnemyBrigands EnemyKind = iota
	EnemyMercenaries
	EnemyUndead
)

type EnemyPlacementKind int

const (
	EnemyPlacementRandomSpread EnemyPlacementKind = iota
	EnemyPlacementOneArmy
	EnemyPlacementTwoArmies
	EnemyPlacementAroundPlayer
	EnemyPlacementNearPlayer
	EnemyPlacementEdges
)

type PlayerPlacementKind int

const (
	PlayerPlacementCenter PlayerPlacementKind = iota
	PlayerPlacementOppositeSide
)

type Config struct {
	Width  int
	Height int

	Mission dat.MissionKind

	Enemy           EnemyKind
	EnemyBudget     int
	EnemyPlacement  EnemyPlacementKind
	PlayerPlacement PlayerPlacementKind

	ForestRatio float64
	SwampRatio  float64
}

func Generate(config Config) *dat.Map {
	m := &dat.Map{}

	const (
		minWidth  = 12
		minHeight = 10
	)

	padWidth := 12 - config.Width
	padHeight := 10 - config.Height

	padOffsetX := 0
	padOffsetY := 0
	if padWidth > 0 {
		padOffsetX = padWidth / 2
	}
	if padHeight > 0 {
		padOffsetY = padHeight / 2
	}

	m.Width = gmath.ClampMin(config.Width, minWidth)
	m.Height = gmath.ClampMin(config.Height, minHeight)
	m.Tiles = make([][]dat.Tile, m.Height)
	for i := range m.Tiles {
		m.Tiles[i] = make([]dat.Tile, m.Width)
	}

	{
		i := 0
		for padWidth > 0 {
			for j := 0; j < m.Height; j++ {
				m.Tiles[j][i] = dat.TileVoid
				m.Tiles[j][m.Width-i-1] = dat.TileVoid
			}
			padWidth -= 2
			i++
		}
	}
	{
		i := 0
		for padHeight > 0 {
			for j := 0; j < m.Width; j++ {
				m.Tiles[i][j] = dat.TileVoid
				m.Tiles[m.Height-i-1][j] = dat.TileVoid
			}
			padHeight -= 2
			i++
		}
	}

	tilePicker := gmath.NewRandPicker[dat.Tile](&game.G.Rand)
	weight := 1.0
	{
		tilePicker.AddOption(dat.TileSwamp, config.SwampRatio)
		weight -= config.SwampRatio
	}
	{
		tilePicker.AddOption(dat.TileForest, config.ForestRatio)
		weight -= config.ForestRatio
	}
	{
		tilePicker.AddOption(dat.TileGrass, weight)
	}
	for row := range m.Tiles {
		for col, t := range m.Tiles[row] {
			if t == dat.TileVoid {
				continue
			}
			m.Tiles[row][col] = tilePicker.Pick()
		}
	}

	occupiedCells := map[dat.CellPos]bool{}

	tmpCells := make([]dat.CellPos, 0, 128)

	switch config.PlayerPlacement {
	case PlayerPlacementCenter:
		numUnits := len(game.G.Units)
		placementSize := int(math.Sqrt(float64(numUnits))/2) + 1
		colFrom := ((m.Width / 2) - 1) - placementSize
		colTo := ((m.Width / 2) - 1) + placementSize
		rowFrom := ((m.Height / 2) - 1) - placementSize
		rowTo := ((m.Height / 2) - 1) + placementSize
		for row := rowFrom; row <= rowTo; row++ {
			for col := colFrom; col <= colTo; col++ {
				tmpCells = append(tmpCells, dat.CellPos{
					X: col,
					Y: row,
				})
			}
		}
		gmath.Shuffle(&game.G.Rand, tmpCells)
		for i := range game.G.Units {
			cell := tmpCells[i]
			m.Units = append(m.Units, dat.DeployedUnit{
				Team: 0,
				Pos:  cell,
				Unit: game.G.Units[i],
			})
			occupiedCells[cell] = true
		}
	default:
		panic("TODO")
	}

	budget := config.EnemyBudget

	unitKindPicker := gmath.NewRandPicker[*dat.UnitStats](&game.G.Rand)
	switch config.Enemy {
	case EnemyBrigands:
		unitKindPicker.AddOption(dat.Brigands, 1)
	default:
		panic("TODO")
	}

	enemyUnits := make([]*dat.Unit, 0, 8)
	for {
		ok := false
		for try := 0; try < 5 && budget > 0; try++ {
			candidate := unitKindPicker.Pick()
			price := candidate.SquadPrice()
			if gmath.Scale(price, 0.8) > budget {
				continue
			}
			budget -= price
			enemyUnits = append(enemyUnits, candidate.CreateUnit())
			ok = true
			break
		}
		if !ok {
			break
		}
	}

	tmpCells = tmpCells[:0] // Re-use them

	switch config.EnemyPlacement {
	case EnemyPlacementEdges:
		for row := padOffsetY; row < m.Height-padOffsetY; row++ {
			for col := padOffsetX; col < m.Width-padOffsetX; col++ {
				cell := dat.CellPos{X: col, Y: row}
				if occupiedCells[cell] {
					continue
				}
				ok := (row == padOffsetY || row == m.Height-padOffsetY-1) ||
					(col == padOffsetX || col == m.Width-padOffsetX-1)
				if ok {
					tmpCells = append(tmpCells, cell)
				}
			}
		}
	default:
		panic("TODO")
	}

	gmath.Shuffle(&game.G.Rand, tmpCells)
	for i, u := range enemyUnits {
		m.Units = append(m.Units, dat.DeployedUnit{
			Team: 1,
			Pos:  tmpCells[i],
			Unit: u,
		})
	}

	return m
}
