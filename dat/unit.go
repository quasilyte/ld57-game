package dat

import resource "github.com/quasilyte/ebitengine-resource"

type Unit struct {
	Stats *UnitStats
}

type UnitStats struct {
	Banner resource.ImageID

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
