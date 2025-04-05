package combat

import (
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/gslices"
	"github.com/quasilyte/ld57-game/dat"
)

type player struct {
	team int

	sceneState *sceneState

	impl playerImpl

	EventDone         gsignal.Event[gsignal.Void]
	EventMeleeAttack  gsignal.Event[meleeAttackEvent]
	EventRangedAttack gsignal.Event[meleeAttackEvent]
}

type meleeAttackEvent struct {
	Attacker *unitNode
	Defender *unitNode
}

type playerImpl interface {
	SetUnit(u *unitNode)
	Update(delta float64)
}

func (p *player) reachableNeighbors(u *unitNode) []dat.CellPos {
	candidates := []dat.CellPos{
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
		{X: 0, Y: -1},
	}
	if u.data.Stats.HasTrait(dat.TraitMobile) {
		candidates = []dat.CellPos{
			{X: 1, Y: 0},
			{X: 1, Y: 1},
			{X: 0, Y: 1},
			{X: -1, Y: 1},
			{X: -1, Y: 0},
			{X: -1, Y: -1},
			{X: 0, Y: -1},
			{X: 1, Y: -1},
		}
	}

	candidates = gslices.FilterInplace(candidates, func(offset dat.CellPos) bool {
		dstPos := u.pos.Add(offset)
		return dstPos.X >= 0 && dstPos.X < p.sceneState.m.Width &&
			dstPos.Y >= 0 && dstPos.Y < p.sceneState.m.Height
	})
	return candidates
}
