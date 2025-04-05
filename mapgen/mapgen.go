package mapgen

import (
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
)

type Config struct {
	Width       int
	Height      int
	ForestRatio float64
	SwampRatio  float64
}

func Generate(config Config) *dat.Map {
	m := &dat.Map{}

	// w=12, h=10

	const (
		minWidth  = 12
		minHeight = 10
	)

	padWidth := 12 - config.Width
	padHeight := 10 - config.Height

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

	return m
}
