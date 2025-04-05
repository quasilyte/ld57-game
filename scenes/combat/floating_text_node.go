package combat

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/styles"
)

type FloatingTextNode struct {
	pos        gmath.Vec
	fadeDelay  float64
	floatDelta float64
	layer      int
	label      *graphics.Label
}

type FloatingTextNodeConfig struct {
	Pos   gmath.Vec
	Text  string
	Layer int
	Color graphics.ColorScale
}

var floatingTextBoxSize = gmath.Vec{X: 32, Y: 16}

func NewFloatingTextNode(config FloatingTextNodeConfig) *FloatingTextNode {
	l := graphics.NewLabel(assets.FontTiny)
	l.SetText(config.Text)

	if config.Color.A == 0 {
		l.SetColorScale(styles.NormalTextColor)
	} else {
		l.SetColorScale(config.Color)
	}

	n := &FloatingTextNode{
		fadeDelay:  0.6,
		label:      l,
		layer:      config.Layer,
		pos:        config.Pos.Sub(floatingTextBoxSize.Mulf(0.5)),
		floatDelta: float64(game.G.Rand.IntRange(-1, 1)),
	}

	n.label.Pos.Base = &n.pos

	return n
}

func (n *FloatingTextNode) Init(s *gscene.Scene) {
	box := n.label.BoundsRect()

	n.label.SetAlignHorizontal(graphics.AlignHorizontalCenter)
	n.label.SetAlignVertical(graphics.AlignVerticalCenter)
	n.label.SetSize(int(box.Width())+2, int(box.Height()))
	game.G.Camera.AddGraphics(n.label, n.layer)
}

func (n *FloatingTextNode) Dispose() {
	n.label.Dispose()
}

func (n *FloatingTextNode) IsDisposed() bool {
	return n.label.IsDisposed()
}

func (n *FloatingTextNode) Update(delta float64) {
	xMultiplier := n.floatDelta
	n.pos.Y -= 10 * delta
	n.pos.X -= (5 * delta) * xMultiplier

	if n.fadeDelay > 0 {
		n.fadeDelay -= delta
		return
	}

	alpha := n.label.GetAlpha()
	if alpha <= 0.08 {
		n.Dispose()
		return
	}
	n.label.SetAlpha(alpha - 1*float32(delta))
}
