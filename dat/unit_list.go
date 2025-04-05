package dat

import "github.com/quasilyte/ld57-game/assets"

var MercenarySwords = &UnitStats{
	// Banner: assets.ImageSkeletalWarriorsBanner,

	Cost: 10,

	MeleeAccuracy: 0.6,
	Attack:        4,
	Defense:       4,
	MaxCount:      15,
	Life:          2,
	Morale:        7,
	Speed:         2,

	Traits: []Trait{},
}

var SkeletalWarriors = &UnitStats{
	Banner: assets.ImageSkeletalWarriorsBanner,

	Cost: 8,

	MeleeAccuracy: 0.5,
	Attack:        3,
	Defense:       3,
	MaxCount:      20,
	Life:          1,
	Morale:        0,
	Speed:         2,

	Traits: []Trait{
		TraitUnbreakable,
		TraitCauseFear,
		TraitArrowResist,
	},
}
