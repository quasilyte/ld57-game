package dat

import (
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gslices"
)

type Unit struct {
	Count int

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
	Banner resource.ImageID

	Class UnitClass

	Cost int

	MeleeAccuracy float64
	Attack        int
	Defense       int
	MaxCount      int
	Life          int
	Morale        int
	Speed         int

	Traits []Trait
}

func (stats *UnitStats) HasTrait(t Trait) bool {
	return gslices.Contains(stats.Traits, t)
}
