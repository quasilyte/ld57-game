package combat

import (
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
)

type computerPlayer struct {
	data *player

	delay float64

	unit *unitNode
}

func (p *computerPlayer) SetUnit(u *unitNode) {
	p.unit = u
	p.delay = 0.2
}

func (p *computerPlayer) Update(delta float64) {
	p.delay -= delta
	if p.delay <= 0 {
		p.act()
	}
}

func (p *computerPlayer) act() {
	switch p.unit.data.Stats.Class {
	case dat.ClassInfantry:
		p.actInfantry()
	case dat.ClassCavalry:
		// TODO
		p.actInfantry()
	case dat.ClassArcher:
		p.actArcher()
	}
}

func (p *computerPlayer) actArcher() {
	if game.G.Rand.Chance(0.09) {
		if p.moveRandomly() {
			p.finishMove()
			return
		}
	}

	if game.G.Rand.Chance(0.85) {
		if target := p.findRangedTarget(); target != nil {
			p.data.EventRangedAttack.Emit(meleeAttackEvent{
				Attacker: p.unit,
				Defender: target,
			})
			p.finishMove()
			return
		}
	}

	if p.updateFavTarget() {
		return
	}

	if game.G.Rand.Chance(0.6) && p.moveRandomly() {
		p.finishMove()
		return
	}

	p.unit.Guard()
	p.finishMove()
}

func (p *computerPlayer) actInfantry() {
	if game.G.Rand.Chance(0.075) {
		if p.moveRandomly() {
			p.finishMove()
			return
		}
	}

	if game.G.Rand.Chance(0.85) {
		if target := p.findMeleeTarget(); target != nil {
			p.data.EventMeleeAttack.Emit(meleeAttackEvent{
				Attacker: p.unit,
				Defender: target,
			})
			p.finishMove()
			return
		}
	}

	if p.updateFavTarget() {
		return
	}

	p.unit.Guard()
	p.finishMove()
}

func (p *computerPlayer) updateFavTarget() bool {
	if p.unit.favTarget == nil {
		if game.G.Rand.Chance(0.95) {
			p.maybeFindFavTarget()
		}
	} else {
		if game.G.Rand.Chance(0.15) {
			p.unit.favTarget = nil
		}
	}

	if p.unit.favTarget != nil {
		if p.unit.favTarget.data.Count == 0 {
			p.unit.favTarget = nil
		}
		if p.unit.favTarget != nil {
			if p.moveTowards(p.unit.favTarget.pos) {
				p.finishMove()
				return true
			}
		}
	}

	return false
}

func (p *computerPlayer) moveTowards(dst dat.CellPos) bool {
	currentDist := dist(dst, p.unit.pos)
	offset := gmath.RandIterate(&game.G.Rand, reachableNeighbors(p.data.sceneState, p.unit), func(offset dat.CellPos) bool {
		newPos := p.unit.pos.Add(offset)
		return p.data.sceneState.unitByCell[newPos] == nil && dist(dst, newPos) < currentDist
	})
	if !offset.IsZero() {
		p.unit.MoveTo(p.unit.pos.Add(offset))
		return true
	}
	return false
}

func (p *computerPlayer) maybeFindFavTarget() {
	p.unit.favTarget = gmath.RandIterate(&game.G.Rand, p.data.sceneState.units, func(u *unitNode) bool {
		return u.team != p.unit.team && u.data.Count > 0
	})
}

func (p *computerPlayer) finishMove() {
	p.data.EventDone.Emit(gsignal.Void{})
}

func (p *computerPlayer) findRangedTarget() *unitNode {
	var bestCandidate *unitNode
	bestScore := 0.0

	reachableRangedTargets(p.data.sceneState, p.unit, func(u2 *unitNode) {
		score := 10.0
		score *= game.G.Rand.FloatRange(0.7, 1.3)
		if dist(p.unit.pos, u2.pos) <= 2 {
			score *= 1.15
		}
		if u2 == p.unit.favTarget {
			score *= 1.05
		}
		if u2.data.Stats.Defense < p.unit.data.Stats.RangedAttack {
			score *= 1.35
		}
		if u2.morale < 0.5 {
			score *= game.G.Rand.FloatRange(1.05, 1.25)
		}
		if u2.data.Stats.HasTrait(dat.TraitArrowResist) {
			score *= 0.8
		}
		if u2.data.Count < u2.data.Stats.MaxCount {
			score *= 1.1
		}
		if score > bestScore {
			bestScore = score
			bestCandidate = u2
		}
	})

	return bestCandidate
}

func (p *computerPlayer) findMeleeTarget() *unitNode {
	var bestCandidate *unitNode
	bestScore := 0.0

	for _, offset := range reachableNeighbors(p.data.sceneState, p.unit) {
		dstPos := p.unit.pos.Add(offset)
		u2 := p.data.sceneState.unitByCell[dstPos]
		if u2 == nil {
			continue
		}
		if u2.team == p.unit.team {
			continue
		}

		score := 10.0
		score *= game.G.Rand.FloatRange(0.8, 1.2)
		if u2 == p.unit.favTarget {
			score *= 1.1
		}
		if u2.data.Stats.Defense < p.unit.data.Stats.Attack {
			score *= 1.2
		}
		if u2.morale < 0.5 {
			score *= game.G.Rand.FloatRange(1.05, 1.25)
		}
		if u2.data.Count < u2.data.Stats.MaxCount {
			score *= 1.1
		}
		switch attackFacing(p.unit, u2) {
		case meleeAttackFlank:
			score *= 1.1
		case meleeAttackBack:
			score *= 1.2
		}
		if score > bestScore {
			bestScore = score
			bestCandidate = u2
		}
	}

	return bestCandidate
}

func (p *computerPlayer) moveRandomly() bool {
	return moveUnitRandomly(p.data.sceneState, p.unit)
}

func moveUnitRandomly(s *sceneState, u *unitNode) bool {
	offset := gmath.RandIterate(&game.G.Rand, reachableNeighbors(s, u), func(offset dat.CellPos) bool {
		dstPos := u.pos.Add(offset)
		if s.unitByCell[dstPos] == nil {
			u.MoveTo(dstPos)
			return true
		}
		return false
	})
	return !offset.IsZero()
}
