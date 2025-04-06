package combat

import (
	"math"

	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/game"
)

type slashNode struct {
	progress float64
	pos      gmath.Vec
	from     gmath.Vec
	to       gmath.Vec
	spr      *graphics.Sprite
	rotation gmath.Rad
}

type slashNodeConfig struct {
	fireFrom gmath.Vec
	fireTo   gmath.Vec
}

func newSlashNode(config slashNodeConfig) *slashNode {
	spr := game.G.NewSprite(assets.ImageSlash)

	dir := config.fireTo.DirectionTo(config.fireFrom).Mulf(40)

	n := &slashNode{
		spr:      spr,
		to:       config.fireTo.Add(dir),
		from:     config.fireFrom,
		rotation: config.fireFrom.AngleToPoint(config.fireTo),
	}

	spr.Pos.Base = &n.pos
	spr.Rotation = &n.rotation

	return n
}

func (n *slashNode) Init(scene *gscene.Scene) {
	scene.AddGraphics(n.spr, 1)
}

func (n *slashNode) IsDisposed() bool {
	return n.spr.IsDisposed()
}

func (n *slashNode) Dispose() {
	n.spr.Dispose()
}

func (n *slashNode) Update(delta float64) {
	n.progress += 2 * delta
	n.rotation += gmath.Rad(2 * delta)
	if n.progress > 0.5 {
		if n.progress >= 1 {
			n.Dispose()
			return
		}
		t := powInterpolate(0, 1, n.progress+0.1, 1.5)
		newPos := n.to.LinearInterpolate(n.from, t)
		n.pos = newPos
	} else {
		t := powInterpolate(0, 1, n.progress+0.1, 3.5)
		newPos := n.from.LinearInterpolate(n.to, t)
		n.pos = newPos
	}
}

func powInterpolate(from, to, t, exp float64) float64 {
	return gmath.LerpClamped(from, to, math.Pow(t, exp))
}
