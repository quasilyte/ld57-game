package dat

import "github.com/quasilyte/ld57-game/assets"

var AllUnits = []*UnitStats{
	// MercenarySwords,

	SkeletalWarriors,
	SkeletalArchers,
	UnholyKnights,
	Zombies,
}

// var MercenarySwords = &UnitStats{
// 	// Banner: assets.ImageSkeletalWarriorsBanner,

// 	Class: ClassInfantry,
// 	Cost:  10,

// 	MeleeAccuracy: 0.6,
// 	Attack:        4,
// 	Defense:       4,
// 	MaxCount:      15,
// 	Life:          2,
// 	Morale:        7,
// 	Speed:         2,

// 	Traits: []Trait{},
// }

var SkeletalWarriors = &UnitStats{
	Name:   "Skel. Warriors",
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

var SkeletalArchers = &UnitStats{
	Name:   "Skel. Archers",
	Banner: assets.ImageSkeletalArchersBanner,

	Class: ClassArcher,
	Cost:  12,

	RangedAccuracy: 0.4,
	MeleeAccuracy:  0.3,
	Attack:         2,
	Defense:        2,
	MaxCount:       15,
	Life:           1,
	Morale:         0,
	Speed:          2,

	Traits: []Trait{
		TraitUnbreakable,
		TraitArrowResist,
	},
}

var UnholyKnights = &UnitStats{
	Name:   "Shadow Knights",
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
