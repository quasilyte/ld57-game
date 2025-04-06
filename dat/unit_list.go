package dat

import "github.com/quasilyte/ld57-game/assets"

var AllUnits = []*UnitStats{
	Brigands,
	Assassins,

	GoblinWarriors,
	OrcWarriors,
	OrcCavalry,

	MercenaryHalberds,
	MercenarySwords,
	MercenaryArchers,
	MercenaryCavalry,

	SkeletalWarriors,
	SkeletalArchers,
	UnholyKnights,
	Zombies,
	Mummies,
}

var MercenaryHalberds = &UnitStats{
	Name:   "Merc. Halberds",
	Banner: assets.ImageHumanHalberdsBanner,

	Class: ClassInfantry,
	Cost:  12,

	MeleeAccuracy: 0.4,
	Attack:        5,
	Defense:       5,
	MaxCount:      12,
	Life:          2,
	Morale:        6,
	Speed:         2,

	Traits: []Trait{
		TraitChargeResist,
		TraitAntiCavalry,
	},
}

var Brigands = &UnitStats{
	Name:   "Brigands",
	Banner: assets.ImageBrigandsBanner,

	Class: ClassInfantry,
	Cost:  5,

	MeleeAccuracy: 0.3,
	Attack:        3,
	Defense:       1,
	MaxCount:      20,
	Life:          2,
	Morale:        4,
	Speed:         2,

	Traits: []Trait{
		TraitMobile,
	},
}

var Assassins = &UnitStats{
	Name:   "Assassins",
	Banner: assets.ImageAssassinsBanner,

	Class: ClassArcher,
	Cost:  22,

	MaxRange:       2,
	RangedAccuracy: 0.6,

	MeleeAccuracy: 0.65,
	RangedAttack:  6,
	Attack:        5,
	Defense:       2,
	MaxCount:      7,
	Life:          2,
	Morale:        5,
	Speed:         2,

	Traits: []Trait{
		TraitCripplingShot,
	},
}

var MercenarySwords = &UnitStats{
	Name:   "Merc. Swords",
	Banner: assets.ImageHumanWarriorsBanner,

	Class: ClassInfantry,
	Cost:  10,

	MeleeAccuracy: 0.6,
	Attack:        4,
	Defense:       4,
	MaxCount:      15,
	Life:          2,
	Morale:        8,
	Speed:         2,

	Traits: []Trait{
		TraitFlankingImmune,
	},
}

var MercenaryArchers = &UnitStats{
	Name:   "Merc. Archers",
	Banner: assets.ImageHumanArchersBanner,

	Class: ClassArcher,
	Cost:  17,

	MaxRange:       3,
	RangedAccuracy: 0.45,

	MeleeAccuracy: 0.25,
	RangedAttack:  4,
	Attack:        2,
	Defense:       1,
	MaxCount:      10,
	Life:          2,
	Morale:        4,
	Speed:         2,

	Traits: []Trait{},
}

var MercenaryCavalry = &UnitStats{
	Name:   "Merc. Cavalry",
	Banner: assets.ImageHumanKnights,

	Class: ClassCavalry,
	Cost:  19,

	MeleeAccuracy: 0.6,
	Attack:        5,
	Defense:       6,
	MaxCount:      10,
	Life:          3,
	Morale:        7,
	Speed:         3,

	Traits: []Trait{
		TraitCharge,
		TraitMobile,
	},
}

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

	MaxRange:       3,
	RangedAccuracy: 0.4,

	MeleeAccuracy: 0.3,
	RangedAttack:  3,
	Attack:        2,
	Defense:       2,
	MaxCount:      15,
	Life:          1,
	Morale:        0,
	Speed:         2,

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
	Name:   "Zombies",
	Banner: assets.ImageZombiesBanner,

	Class: ClassInfantry,
	Cost:  9,

	MeleeAccuracy: 0.4,
	Attack:        2,
	Defense:       2,
	MaxCount:      10,
	Life:          3,
	Morale:        0,
	Speed:         1,

	Traits: []Trait{
		TraitUnbreakable,
		TraitCauseFear,
		TraitStun,
		TraitArrowVulnerability,
	},
}

var Mummies = &UnitStats{
	Name:   "Mummies",
	Banner: assets.ImageHumanMummiesBanner,

	Class: ClassInfantry,
	Cost:  30,

	MeleeAccuracy: 0.85,
	Attack:        4,
	Defense:       9,
	MaxCount:      5,
	Life:          5,
	Morale:        0,
	Speed:         1,

	Traits: []Trait{
		TraitMobile,
		TraitUnbreakable,
		TraitCauseFear,
	},
}

var GoblinWarriors = &UnitStats{
	Name:   "Gob. Warriors",
	Banner: assets.ImageGoblinWarriorBanner,

	Class: ClassInfantry,
	Cost:  6,

	MeleeAccuracy: 0.3,
	Attack:        1,
	Defense:       1,
	MaxCount:      30,
	Life:          1,
	Morale:        3,
	Speed:         2,

	Traits: []Trait{
		TraitPathfinder,
	},
}

var OrcWarriors = &UnitStats{
	Name:   "Orc Warriors",
	Banner: assets.ImageOrcWarriorBanner,

	Class: ClassInfantry,
	Cost:  20,

	MeleeAccuracy: 0.7,
	Attack:        5,
	Defense:       2,
	MaxCount:      15,
	Life:          4,
	Morale:        6,
	Speed:         2,

	Traits: []Trait{
		TraitBloodlust,
	},
}

var OrcCavalry = &UnitStats{
	Name:   "Orc Boar Elite",
	Banner: assets.ImageOrcBoarEliteBanner,

	Class: ClassCavalry,
	Cost:  35,

	RangedAccuracy: 0.3,
	MeleeAccuracy:  0.75,
	Attack:         6,
	RangedAttack:   4,
	Defense:        3,
	MaxCount:       10,
	MaxRange:       2,
	Life:           4,
	Morale:         7,
	Speed:          3,

	Traits: []Trait{
		TraitBloodlust,
	},
}
