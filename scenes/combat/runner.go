package combat

import (
	"strconv"

	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
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
		if u.data.Count > 0 {
			u.movesLeft = u.data.Stats.Speed
			if u.data.Stats.HasTrait(dat.TraitRegen) {
				u.leftoverHP = u.maxHP()
			}
			if u.data.Experience > 1 {
				u.data.Experience = 0
				u.data.Level++
			}
			if u.broken && game.G.Rand.Chance(0.5) {
				u.movesLeft = gmath.ClampMin(u.movesLeft-1, 1)
			}
			u.steps = 0
			u.guard = false
			u.AddMorale(game.G.Rand.FloatRange(0.01, 0.03))
			// 9 => 0.045 (extra ~5% morale per turn).
			u.AddMorale(0.005 * float64(u.data.Stats.Morale))
			if u.broken {
				u.AddMorale(game.G.Rand.FloatRange(0.015, 0.025))
				// 9 => 0.081 (extra ~8% morale per turn when broken).
				u.AddMorale(0.009 * float64(u.data.Stats.Morale))
				recovered := u.data.HasItem(dat.ItemBannerOfWill) ||
					(u.morale >= 0.40 && game.G.Rand.Chance(u.morale))
				if recovered {
					u.broken = false
					u.AddMorale(0.05)
					u.updateCountLabel()
					n := NewFloatingTextNode(FloatingTextNodeConfig{
						Pos:   u.spritePos,
						Text:  "Regroup!",
						Layer: 3,
						Color: pickColor(u.team, true),
					})
					r.sceneState.scene.AddObject(n)
				}
			}
		}
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

func pickColor(team int, good bool) graphics.ColorScale {
	targetTeam := 0
	if good {
		targetTeam = 1
	}

	clr := styles.ColorRed
	if team != targetTeam {
		clr = styles.ColorTeal
	}
	return clr
}

func (r *runner) withCasualtiesCheck(melee bool, attacker, defender *unitNode, f func()) {
	initialAttackers := attacker.data.Count
	initialDefenders := defender.data.Count

	f()

	deadAttackers := initialAttackers - attacker.data.Count
	deadDefenders := initialDefenders - defender.data.Count
	if deadAttackers != 0 {
		attacker.SubMorale(game.G.Rand.FloatRange(0.05, 0.1) * float64(deadAttackers))
	}
	if deadDefenders != 0 {
		s := "-" + strconv.Itoa(deadDefenders)
		if melee {
			facing := attackFacing(attacker, defender)
			switch facing {
			case meleeAttackFlank:
				s += " Flanked"
			case meleeAttackBack:
				s += " Backstab"
			}
		}
		n := NewFloatingTextNode(FloatingTextNodeConfig{
			Pos:   defender.spritePos,
			Text:  s,
			Layer: 3,
			Color: pickColor(defender.team, false),
		})
		r.sceneState.scene.AddObject(n)
		defender.SubMorale(game.G.Rand.FloatRange(0.05, 0.1) * float64(deadDefenders))
	}
	if deadAttackers+deadDefenders > 0 {
		// TODO: it doesn't work.
		if attacker.data.Count == 0 || defender.data.Count == 0 {
			r.sceneState.pause = 1.0
		} else {
			r.sceneState.pause = 0.65
		}
	}

	if attacker.data.HasItem(dat.ItemTerrorMace) {
		if !defender.data.HasItem(dat.ItemRingOfCourage) {
			fearDmg := (0.02 * float64(deadDefenders)) + 0.05
			defender.SubMorale(fearDmg)
		}
	}
	if attacker.data.Stats.HasTrait(dat.TraitCauseFear) {
		if !defender.data.HasItem(dat.ItemRingOfCourage) {
			// 1% morale damage per unit, maxed at 15.
			// At level 8+, it's 1.5% damage per unit (so ~22% against 15%)
			fearDmg := 0.01
			switch {
			case attacker.data.Level >= 4:
				fearDmg = 0.0125
			case attacker.data.Level >= 8:
				fearDmg = 0.015
			}
			defender.SubMorale(gmath.ClampMax(fearDmg*float64(initialAttackers), 0.15))
		}
	}

	const expScaler = 0.85
	if deadDefenders > 0 {
		m := float64(defender.data.Level+1) / float64(attacker.data.Level+1)
		attacker.AddExperience(m * float64(deadDefenders) * (0.01 * (float64(defender.data.Stats.Cost) * expScaler)))
	}
	if deadAttackers > 0 {
		m := float64(attacker.data.Level+1) / float64(defender.data.Level+1)
		defender.AddExperience(m * float64(deadAttackers) * (0.01 * (float64(attacker.data.Stats.Cost) * (expScaler / 2))))
	}

	if deadAttackers > 0 && attacker.morale < 0.5 && game.G.Rand.Chance(1.0-attacker.morale) {
		attacker.broken = true
		attacker.SubMorale(0.1)
		attacker.updateCountLabel()
	}
	if !attacker.broken && deadDefenders > 0 && defender.morale < 0.5 && game.G.Rand.Chance(1.0-defender.morale) {
		defender.broken = true
		defender.SubMorale(0.1)
		defender.updateCountLabel()
	}

	if melee && deadDefenders > 0 && !attacker.broken {
		if attacker.data.Stats.HasTrait(dat.TraitBloodlust) {
			attacker.AddMorale(0.035 * float64(deadDefenders))
		}
		if attacker.data.Count < attacker.data.InitialCount && attacker.data.Stats.HasTrait(dat.TraitSoulHarvest) {
			// 3 mercenary swords of level 1 (+1) => 3 * (10+1) => 33 => ~33%
			totalValue := float64(deadDefenders * (defender.data.Stats.Cost + defender.data.Level))
			if totalValue > 15 && game.G.Rand.Chance(totalValue/100) {
				attacker.data.Count++
				deadAttackers--
				attacker.updateCountLabel()
			}
		}
	}

	if deadAttackers != 0 {
		var s string
		good := false
		if deadAttackers > 0 {
			s = "-" + strconv.Itoa(deadAttackers)
		} else {
			good = true
			s = "+" + strconv.Itoa(deadAttackers)
		}
		n := NewFloatingTextNode(FloatingTextNodeConfig{
			Pos:   attacker.spritePos,
			Text:  s,
			Layer: 3,
			Color: pickColor(attacker.team, good),
		})
		r.sceneState.scene.AddObject(n)
	}
}

func (r *runner) runRangedRound(attacker, defender *unitNode) {
	r.withCasualtiesCheck(false, attacker, defender, func() {
		for i := 0; i < attacker.data.Count; i++ {
			attackerDmg := r.runRangedAttack(attacker, defender)
			if r.damageUnit(defender, attackerDmg) {
				break
			}
		}
		if defender.favTarget == nil || game.G.Rand.Chance(0.6) {
			defender.favTarget = attacker
		}
	})
}

func (r *runner) runMeleeRound(attacker, defender *unitNode) {
	r.withCasualtiesCheck(true, attacker, defender, func() {
		facing := attackFacing(attacker, defender)

		totalAttackerDmg := 0
		totalDefenderDmg := 0

		numAttacks := attacker.data.Count
		if attacker.data.Stats.HasTrait(dat.TraitMighty) {
			numAttacks *= 2
		}

		var retaliationsLeft int
		if attacker.data.Stats.HasTrait(dat.TraitNoRetaliation) {
			retaliationsLeft = 0
		} else {
			retaliationsLeft = defender.data.Count + 1
			if defender.guard {
				retaliationsLeft += (defender.data.Count / 2)
			}
			if defender.data.HasItem(dat.ItemVengeanceHelm) {
				retaliationsLeft += gmath.ClampMin((defender.data.Count / 2), 1)
			}
			if defender.data.Stats.HasTrait(dat.TraitMighty) {
				retaliationsLeft *= 2
			}
		}

		for i := 0; i < numAttacks; i++ {
			attackerDmg := r.runMeleeAttack(false, attacker, defender, facing)
			defenderDmg := 0
			if !defender.broken && retaliationsLeft > 0 {
				retaliationsLeft--
				defenderDmg = r.runMeleeAttack(true, defender, attacker, meleeAttackFront)
			}

			totalAttackerDmg += attackerDmg
			totalDefenderDmg += defenderDmg

			if r.damageUnit(defender, attackerDmg) {
				break
			}
			if r.damageUnit(attacker, defenderDmg) {
				break
			}
		}

		switch facing {
		case meleeAttackFlank:
			defender.SubMorale(defender.morale * 0.2)
		case meleeAttackBack:
			defender.SubMorale(defender.morale * 0.5)
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
	if r.sceneState.m.Tiles[defender.pos.Y][defender.pos.X] == dat.TileForest {
		toHit *= 0.55
	}
	if attacker.data.HasItem(dat.ItemPointblankBow) {
		isDiagonal := abs(attacker.pos.X-defender.pos.X) == 1 && abs(attacker.pos.Y-defender.pos.Y) == 1
		dist := dist(attacker.pos, defender.pos)
		if isDiagonal && dist == 1 {
			toHit *= 1.4
		}
	}
	// +3% accuracy per level.
	toHit *= 1.0 + (0.03 * float64(attacker.data.Level))
	if !game.G.Rand.Chance(toHit) {
		return 0
	}

	// +3% damage block chance per level.
	dodgeChance := 0.03 * float64(defender.data.Level)
	if dodgeChance > 0 && game.G.Rand.Chance(dodgeChance) {
		return 0
	}

	atk := 0.1 * float64(attacker.data.Stats.RangedAttack)
	if defender.data.Stats.HasTrait(dat.TraitArrowVulnerability) {
		atk *= 1.5
	}
	defenseLevel := defender.data.Stats.Defense
	if defender.guard {
		defenseLevel++
	}
	def := 0.08 * (float64(defenseLevel))
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

func (r *runner) runMeleeAttack(isRetaliation bool, attacker, defender *unitNode, facing meleeAttackFacing) int {
	toHit := attacker.data.Stats.MeleeAccuracy
	if attacker.morale < 0.5 {
		toHit *= 0.75
	}
	if defender.broken {
		toHit += 0.01
		toHit *= 1.15
	}
	if defender.data.Stats.Class == dat.ClassCavalry && attacker.data.Stats.HasTrait(dat.TraitAntiCavalry) {
		toHit *= 1.1
	}
	switch facing {
	case meleeAttackFlank:
		toHit *= 1.1
		if !isRetaliation && attacker.data.HasItem(dat.ItemBackstabber) {
			toHit *= 1.4
		}
	case meleeAttackBack:
		toHit *= 1.2
	}
	// +5% accuracy per level.
	toHit *= 1.0 + (0.05 * float64(attacker.data.Level))
	if !game.G.Rand.Chance(toHit) {
		return 0
	}
	if !game.G.Rand.Chance(toHit) {
		return 0
	}

	// +4% damage block chance per level.
	dodgeChance := 0.04 * float64(defender.data.Level)
	if dodgeChance > 0 && game.G.Rand.Chance(dodgeChance) {
		return 0
	}

	atk := 0.1 * (float64(attacker.data.Stats.Attack) * attacker.morale)
	if defender.broken {
		atk *= 1.2
	}
	if !isRetaliation && attacker.steps > 0 && attacker.data.Stats.HasTrait(dat.TraitCharge) {
		if !defender.data.Stats.HasTrait(dat.TraitChargeResist) {
			if attacker.steps == 1 {
				atk *= 1.2
			} else {
				atk *= 1.5
			}
		}
	}
	defenseLevel := defender.data.Stats.Defense
	if defender.guard {
		defenseLevel++
	}
	def := 0.08 * (float64(defenseLevel))
	if attacker.data.Stats.Class == dat.ClassCavalry && defender.data.Stats.HasTrait(dat.TraitAntiCavalry) {
		def *= 1.5
	}
	if !isRetaliation && attacker.data.HasItem(dat.ItemMagicSword) {
		def *= 0.5
	}
	critChance := atk - def
	isCrit := critChance > 0 && game.G.Rand.Chance(critChance)
	dmg := 1
	if defender.data.Stats == dat.Ogres || defender.data.Stats == dat.Troll {
		if !isRetaliation && attacker.data.HasItem(dat.ItemTrollbane) {
			dmg++
		}
	}
	if defender.data.Stats.Class == dat.ClassCavalry && attacker.data.Stats.HasTrait(dat.TraitAntiCavalry) {
		dmg++
	}
	if isCrit {
		dmg *= 2
	}
	return dmg
}
