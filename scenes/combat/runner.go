package combat

import (
	"strconv"

	"github.com/quasilyte/gslices"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/styles"
)

type meleeAttackFacing int

const (
	meleeAttackFront meleeAttackFacing = iota
	meleeAttackFlank
	meleeAttackBack
)

type runner struct {
	sceneState *sceneState

	currentUnit int
	turnOrder   []scheduledUnit
}

type scheduledUnit struct {
	t    float64
	unit *unitNode
}

func (r *runner) NextTurn() {
	r.currentUnit = 0
	r.turnOrder = r.turnOrder[:0]

	maxTime := 5.0

	r.sceneState.units = gslices.FilterInplace(r.sceneState.units, func(u *unitNode) bool {
		u.movesLeft = u.data.Stats.Speed
		u.afterTurn()
		return u.data.Count > 0
	})

	for _, u := range r.sceneState.units {
		thresholdStep := maxTime / float64(u.data.Stats.Speed)
		currentMax := thresholdStep
		currentMin := 0.0
		for i := 0; i < u.data.Stats.Speed; i++ {
			t := game.G.Rand.FloatRange(currentMin, currentMax)
			r.turnOrder = append(r.turnOrder, scheduledUnit{
				t:    t,
				unit: u,
			})
			currentMin += thresholdStep
			currentMax += thresholdStep
		}
	}

	gslices.SortFunc(r.turnOrder, func(x, y scheduledUnit) bool {
		return x.t < y.t
	})
}

func (r *runner) NextUnit() *unitNode {
	var unit *unitNode
	for r.currentUnit < len(r.turnOrder) {
		u := r.turnOrder[r.currentUnit]
		r.currentUnit++
		canAct := u.unit.data.Count > 0 && u.unit.movesLeft > 0
		if canAct {
			unit = u.unit
			break
		}
	}
	return unit
}

func (r *runner) withCasualtiesCheck(attacker, defender *unitNode, f func()) {
	initialAttackers := attacker.data.Count
	initialDefenders := defender.data.Count

	f()

	deadAttackers := initialAttackers - attacker.data.Count
	deadDefenders := initialDefenders - defender.data.Count
	if deadAttackers != 0 {
		clr := styles.ColorRed
		if attacker.team != 0 {
			clr = styles.ColorTeal
		}
		n := NewFloatingTextNode(FloatingTextNodeConfig{
			Pos:   attacker.spritePos,
			Text:  "-" + strconv.Itoa(deadAttackers),
			Layer: 3,
			Color: clr,
		})
		r.sceneState.scene.AddObject(n)
	}
	if deadDefenders != 0 {
		clr := styles.ColorRed
		if defender.team != 0 {
			clr = styles.ColorTeal
		}
		n := NewFloatingTextNode(FloatingTextNodeConfig{
			Pos:   defender.spritePos,
			Text:  "-" + strconv.Itoa(deadDefenders),
			Layer: 3,
			Color: clr,
		})
		r.sceneState.scene.AddObject(n)
	}
}

func (r *runner) runRangedRound(attacker, defender *unitNode) {
	r.withCasualtiesCheck(attacker, defender, func() {
		for i := 0; i < attacker.data.Count; i++ {
			attackerDmg := r.runRangedAttack(attacker, defender)
			if r.damageUnit(defender, attackerDmg) {
				break
			}
		}
	})
}

func (r *runner) runMeleeRound(attacker, defender *unitNode) {
	r.withCasualtiesCheck(attacker, defender, func() {
		facing := attackFacing(attacker, defender)

		totalAttackerDmg := 0
		totalDefenderDmg := 0

		for i := 0; i < attacker.data.Count; i++ {
			attackerDmg := r.runMeleeAttack(attacker, defender, facing)
			defenderDmg := r.runMeleeAttack(defender, attacker, meleeAttackFront)

			totalAttackerDmg += attackerDmg
			totalDefenderDmg += defenderDmg

			if r.damageUnit(defender, attackerDmg) {
				break
			}
			if r.damageUnit(attacker, defenderDmg) {
				break
			}
		}
	})
}

func (r *runner) damageUnit(u *unitNode, dmg int) bool {
	return u.onDamage(dmg)
}

func (r *runner) runRangedAttack(attacker, defender *unitNode) int {
	toHit := attacker.data.Stats.RangedAccuracy
	if attacker.morale < 0.5 {
		toHit *= 0.85
	}
	if !game.G.Rand.Chance(toHit) {
		return 0
	}

	atk := 0.1 * float64(attacker.data.Stats.RangedAttack)
	if defender.data.Stats.HasTrait(dat.TraitArrowVulnerability) {
		atk *= 1.5
	}
	def := 0.08 * (float64(defender.data.Stats.Defense))
	if defender.data.Stats.HasTrait(dat.TraitArrowResist) {
		def *= 1.5
	}
	critChance := atk - def
	isCrit := critChance > 0 && game.G.Rand.Chance(critChance)
	dmg := 1
	if isCrit {
		dmg *= 2
	}
	return dmg
}

func (r *runner) runMeleeAttack(attacker, defender *unitNode, facing meleeAttackFacing) int {
	toHit := attacker.data.Stats.MeleeAccuracy
	if attacker.morale < 0.5 {
		toHit *= 0.75
	}
	switch facing {
	case meleeAttackFlank:
		toHit *= 1.1
	case meleeAttackBack:
		toHit *= 1.2
	}
	if !game.G.Rand.Chance(toHit) {
		return 0
	}

	atk := 0.1 * (float64(attacker.data.Stats.Attack) * attacker.morale)
	def := 0.08 * (float64(defender.data.Stats.Defense))
	switch facing {
	case meleeAttackFlank:
		def *= 0.75
	case meleeAttackBack:
		def *= 0.25
	}
	critChance := atk - def
	isCrit := critChance > 0 && game.G.Rand.Chance(critChance)
	dmg := 1
	if isCrit {
		dmg *= 2
	}
	return dmg
}
