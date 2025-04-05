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
	}
}

func (p *computerPlayer) actInfantry() {
	p.moveRandomly()
	p.finishMove()
}

func (p *computerPlayer) finishMove() {
	p.data.EventDone.Emit(gsignal.Void{})
}

func (p *computerPlayer) moveRandomly() bool {
	offset := gmath.RandIterate(&game.G.Rand, p.data.reachableNeighbors(p.unit), func(offset dat.CellPos) bool {
		dstPos := p.unit.pos.Add(offset)
		if p.data.sceneState.unitByCell[dstPos] == nil {
			p.unit.MoveTo(dstPos)
			return true
		}
		return false
	})
	return !offset.IsZero()
}
