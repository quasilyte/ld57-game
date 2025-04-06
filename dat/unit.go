package dat

import (
	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gslices"
)

type UnitFaction int

const (
	FactionNeutral UnitFaction = iota
	FactionHuman
	FactionUndead
	FactionHorde
)

type Unit struct {
	Count int

	InitialCount int

	Experience float64
	Level      int

	Items [2]*ItemStats

	Stats *UnitStats
}

func (u *Unit) Clone() *Unit {
	cp := *u
	return &cp
}

func (u *Unit) HasItem(item *ItemStats) bool {
	for i := range u.Items {
		if u.Items[i] == item {
			return true
		}
	}
	return false
}

type UnitClass int

const (
	ClassInfantry UnitClass = iota
	ClassCavalry
	ClassArcher
	ClassCaster
)

func (class UnitClass) String() string {
	switch class {
	case ClassInfantry:
		return "infantry"
	case ClassCavalry:
		return "cavalry"
	case ClassArcher:
		return "archer"
	default:
		return "hero"
	}
}

type UnitStats struct {
	Name        string
	Banner      resource.ImageID
	AltBanner   *ebiten.Image
	ScaledImage *ebiten.Image

	Category UnitFaction

	AttackSound resource.AudioID

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

func (stats *UnitStats) SquadPrice() int {
	return stats.Cost * stats.MaxCount
}

func (stats *UnitStats) CreateUnit() *Unit {
	return &Unit{
		Count:        stats.MaxCount,
		InitialCount: stats.MaxCount,
		Level:        0, // Displayed as level 1
		Stats:        stats,
	}
}

func (stats *UnitStats) HasTrait(t Trait) bool {
	return gslices.Contains(stats.Traits, t)
}
