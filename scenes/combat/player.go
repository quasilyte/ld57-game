package combat

import (
	"github.com/quasilyte/gmath"
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

func reachableRangedTargets(s *sceneState, u *unitNode, f func(target *unitNode)) {
	colFrom := gmath.ClampMin(u.pos.X-u.data.Stats.MaxRange, 0)
	rowFrom := gmath.ClampMin(u.pos.Y-u.data.Stats.MaxRange, 0)
	colTo := gmath.ClampMax(u.pos.X+u.data.Stats.MaxRange, s.m.Width-1)
	rowTo := gmath.ClampMax(u.pos.Y+u.data.Stats.MaxRange, s.m.Height-1)
	for row := rowFrom; row <= rowTo; row++ {
		for col := colFrom; col <= colTo; col++ {
			cell := dat.CellPos{X: col, Y: row}
			if dist(cell, u.pos) < 2 {
				continue
			}
			u2 := s.unitByCell[cell]
			if u2 == nil {
				continue
			}
			if u2.team == u.team {
				continue
			}
			f(u2)
		}
	}
}

func reachableNeighbors(s *sceneState, u *unitNode) []dat.CellPos {
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
		return dstPos.X >= 0 && dstPos.X < s.m.Width &&
			dstPos.Y >= 0 && dstPos.Y < s.m.Height
	})
	return candidates
}
