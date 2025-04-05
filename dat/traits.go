package dat

type Trait int

const (
	TraitUnknown Trait = iota

	TraitFlankingImmune
	TraitUnbreakable
	TraitCauseFear
	TraitArrowResist
	TraitArrowVulnerability
	TraitMobile
)
