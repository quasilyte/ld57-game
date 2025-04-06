package combat

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/game"
)

type arrowNode struct {
	rotation       gmath.Rad
	pos            gmath.Vec
	arcFrom        gmath.Vec
	arcTo          gmath.Vec
	startPos       gmath.Vec
	targetPos      gmath.Vec
	arcScaling     float64
	arcProgression float64
	spr            *graphics.Sprite
}

type arrowNodeConfig struct {
	fireFrom   gmath.Vec
	fireTo     gmath.Vec
	flightTime float64
	arcPower   float64
}

func newArrowNode(config arrowNodeConfig) *arrowNode {
	spr := game.G.NewSprite(assets.ImageArrow)

	n := &arrowNode{
		spr:        spr,
		arcScaling: 1.0 / config.flightTime,
		pos:        config.fireFrom,
	}

	powerOffset := config.arcPower
	n.arcFrom = config.fireFrom.Add(gmath.Vec{Y: 0.5 * powerOffset})
	n.arcTo = config.fireTo.Add(gmath.Vec{Y: 0.75 * powerOffset})
	n.startPos = config.fireFrom
	n.targetPos = config.fireTo

	spr.Pos.Base = &n.pos
	spr.Rotation = &n.rotation

	return n
}

func (n *arrowNode) Init(scene *gscene.Scene) {
	scene.AddGraphics(n.spr, 1)
}

func (n *arrowNode) IsDisposed() bool {
	return n.spr.IsDisposed()
}

func (n *arrowNode) Dispose() {
	n.spr.Dispose()
}

func (n *arrowNode) Update(delta float64) {
	n.arcProgression += delta * n.arcScaling
	if n.arcProgression >= 1 {
		n.Dispose()
		return
	}
	newPos := n.startPos.CubicInterpolate(n.arcFrom, n.targetPos, n.arcTo, n.arcProgression)
	n.rotation = n.pos.AngleToPoint(newPos)
	n.pos = newPos
}
