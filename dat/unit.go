package dat

import (
	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gslices"
)

type Unit struct {
	Count int

	Level int

	Stats *UnitStats
}

type UnitClass int

const (
	ClassInfantry UnitClass = iota
	ClassCavalry
	ClassArcher
	ClassCaster
)

type UnitStats struct {
	Name        string
	Banner      resource.ImageID
	ScaledImage *ebiten.Image

	Class UnitClass

	Cost int

	MaxRange       int
	RangedAccuracy float64
	MeleeAccuracy  float64
	Attack         int
	RangedAttack   int
	Defense        int
	MaxCount       int
	Life           int
	Morale         int
	Speed          int

	Traits []Trait
}

func (stats *UnitStats) HasTrait(t Trait) bool {
	return gslices.Contains(stats.Traits, t)
}
