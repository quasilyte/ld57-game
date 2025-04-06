package dat

import "github.com/quasilyte/ld57-game/assets"

var AllUnits = []*UnitStats{
	Brigands,
	Assassins,

	GoblinWarriors,
	OrcWarriors,
	OrcCavalry,
	Ogres,

	MercenaryHalberds,
	MercenarySwords,
	MercenaryArchers,
	MercenaryCavalry,

	SkeletalWarriors,
	SkeletalArchers,
	UnholyKnights,
	Zombies,
	Mummies,
	Reapers,

	Troll,
}

var MercenaryHalberds = &UnitStats{
	Name:        "Merc. Halberds",
	Banner:      assets.ImageHumanHalberdsBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionHuman,

	Class:        ClassInfantry,
	Cost:         10,
	ExtraBuyCost: -5,

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
	Name:        "Brigands",
	Banner:      assets.ImageBrigandsBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionNeutral,

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
	Name:        "Assassins",
	Banner:      assets.ImageAssassinsBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionNeutral,

	Class:        ClassArcher,
	Cost:         25,
	ExtraBuyCost: 5,

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
	Name:        "Merc. Swords",
	Banner:      assets.ImageHumanWarriorsBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionHuman,

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
	Name:        "Merc. Archers",
	Banner:      assets.ImageHumanArchersBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionHuman,

	Class:        ClassArcher,
	Cost:         17,
	ExtraBuyCost: 10,

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
	Name:        "Merc. Cavalry",
	Banner:      assets.ImageHumanKnights,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionHuman,

	Class: ClassCavalry,
	Cost:  19,

	MeleeAccuracy: 0.7,
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
	Name:        "Skel. Warriors",
	Banner:      assets.ImageSkeletalWarriorsBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionUndead,

	Class: ClassInfantry,
	Cost:  7,

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
		TraitStunResist,
	},
}

var SkeletalArchers = &UnitStats{
	Name:        "Skel. Archers",
	Banner:      assets.ImageSkeletalArchersBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionUndead,

	Class:        ClassArcher,
	Cost:         12,
	ExtraBuyCost: 10,

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
	Name:        "Shadow Knights",
	Banner:      assets.ImageUnholyKnights,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionUndead,

	Class: ClassCavalry,
	Cost:  16,

	MeleeAccuracy: 0.75,
	Attack:        5,
	Defense:       4,
	MaxCount:      10,
	Life:          3,
	Morale:        0,
	Speed:         3,

	Traits: []Trait{
		TraitCauseFear,
		TraitMobile,
		TraitStunResist,
	},
}

var Zombies = &UnitStats{
	Name:        "Zombies",
	Banner:      assets.ImageZombiesBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionUndead,

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
	Name:        "Mummies",
	Banner:      assets.ImageHumanMummiesBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionUndead,

	Class:        ClassInfantry,
	Cost:         30,
	ExtraBuyCost: 5,

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

var Ogres = &UnitStats{
	Name:        "Ogres",
	Banner:      assets.ImageOgresBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionHorde,

	Class:        ClassInfantry,
	Cost:         42,
	ExtraBuyCost: 40,

	MeleeAccuracy: 0.925,
	Attack:        5,
	Defense:       7,
	MaxCount:      3,
	Life:          9,
	Morale:        7,
	Speed:         2,

	Traits: []Trait{
		TraitCauseFear,
		TraitStun,
		TraitStunResist,
	},
}

var Reapers = &UnitStats{
	Name:        "Reapers",
	Banner:      assets.ImageReapersBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionUndead,

	Class: ClassInfantry,
	Cost:  50,

	MeleeAccuracy: 0.8,
	Attack:        7,
	Defense:       3,
	MaxCount:      6,
	Life:          6,
	Morale:        9,
	Speed:         2,

	Traits: []Trait{
		TraitCauseFear,
		TraitSoulHarvest,
		TraitArrowResist,
		TraitNoRetaliation,
	},
}

var Troll = &UnitStats{
	Name:        "Troll",
	Banner:      assets.ImageTrollBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionHorde,

	Class: ClassInfantry,
	Cost:  275,

	MeleeAccuracy: 1.0,
	Attack:        9,
	Defense:       11,
	MaxCount:      1,
	Life:          17,
	Morale:        9,
	Speed:         2,

	Traits: []Trait{
		TraitStunResist,
		TraitRegen,
		TraitMighty,
		TraitPathfinder,
	},
}

var GoblinWarriors = &UnitStats{
	Name:        "Gob. Warriors",
	Banner:      assets.ImageGoblinWarriorBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionHorde,

	Class: ClassInfantry,
	Cost:  4,

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
	Name:        "Orc Warriors",
	Banner:      assets.ImageOrcWarriorBanner,
	AttackSound: assets.AudioSwordAttack1,

	Category: FactionHorde,

	Class: ClassInfantry,
	Cost:  17,

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
	Name:        "Orc Boar Elite",
	Banner:      assets.ImageOrcBoarEliteBanner,
	AttackSound: assets.AudioBluntAttack1,

	Category: FactionHorde,

	Class: ClassCavalry,
	Cost:  30,

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
