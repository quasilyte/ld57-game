package dat

import "github.com/quasilyte/ld57-game/assets"

var MercenarySwords = &UnitStats{
	// Banner: assets.ImageSkeletalWarriorsBanner,

	Class: ClassInfantry,
	Cost:  10,

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

	Class: ClassInfantry,
	Cost:  8,

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

var UnholyKnights = &UnitStats{
	Banner: assets.ImageUnholyKnights,

	Class: ClassCavalry,
	Cost:  16,

	MeleeAccuracy: 0.65,
	Attack:        5,
	Defense:       4,
	MaxCount:      10,
	Life:          3,
	Morale:        0,
	Speed:         3,

	Traits: []Trait{
		TraitUnbreakable,
		TraitCauseFear,
		TraitMobile,
	},
}

var Zombies = &UnitStats{
	Banner: assets.ImageZombiesBanner,

	Class: ClassInfantry,
	Cost:  5,

	MeleeAccuracy: 0.4,
	Attack:        2,
	Defense:       2,
	MaxCount:      10,
	Life:          2,
	Morale:        0,
	Speed:         1,

	Traits: []Trait{
		TraitUnbreakable,
		TraitCauseFear,
	},
}
