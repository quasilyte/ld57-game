package combat

import (
	"strconv"

	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/styles"
)

type unitNode struct {
	data            *dat.Unit
	team            int
	sceneState      *sceneState
	spritePos       gmath.Vec
	countLabel      *graphics.Label
	pos             dat.CellPos
	sprite          *graphics.Sprite
	facingIndicator *graphics.Sprite
	movesIndicator  *graphics.Label
	leftoverHP      int
	movesLeft       int
	facing          int
	morale          float64

	favTarget *unitNode
}

type unitNodeConfig struct {
	Team  int
	State *sceneState
	Data  *dat.Unit
	Pos   dat.CellPos
}

func newUnitNode(config unitNodeConfig) *unitNode {
	spr := game.G.NewSprite(config.Data.Stats.Banner)
	n := &unitNode{
		sceneState: config.State,
		sprite:     spr,
		data:       config.Data,
		pos:        config.Pos,
		team:       config.Team,
		leftoverHP: config.Data.Stats.Life,
		morale:     1.0,
		facing:     game.G.Rand.IntRange(0, 3),
	}
	if config.Team != 0 {
		spr.SetColorScale(graphics.ColorScale{R: 1.2, G: 1.2, B: 0.6, A: 1})
	}

	n.facingIndicator = graphics.NewSprite()
	n.facingIndicator.Pos.Base = &n.spritePos
	n.facingIndicator.Pos.Offset = gmath.Vec{X: 11, Y: 10}

	n.countLabel = graphics.NewLabel(assets.FontTiny)
	n.countLabel.Pos.Base = &n.spritePos
	n.countLabel.SetColorScale(styles.BrightTextColor)
	n.countLabel.SetSize(32, 20)
	n.countLabel.Pos.Offset.X -= 16
	n.countLabel.Pos.Offset.Y = 1
	n.countLabel.SetAlignHorizontal(graphics.AlignHorizontalLeft)
	n.countLabel.SetAlignVertical(graphics.AlignVerticalCenter)

	n.movesIndicator = graphics.NewLabel(assets.FontTiny)
	n.movesIndicator.SetText("....")
	n.movesIndicator.Pos.Base = &n.spritePos
	n.movesIndicator.SetColorScale(styles.BrightTextColor)
	n.movesIndicator.SetSize(32, 20)
	n.movesIndicator.Pos.Offset.X -= 16
	n.movesIndicator.Pos.Offset.Y = -28
	n.movesIndicator.SetAlignHorizontal(graphics.AlignHorizontalCenter)
	n.movesIndicator.SetAlignVertical(graphics.AlignVerticalCenter)

	spr.Pos.Base = &n.spritePos
	return n
}

func (u *unitNode) IsDisposed() bool {
	return u.sprite.IsDisposed()
}

func (u *unitNode) Dispose() {
	u.sprite.Dispose()
	u.countLabel.Dispose()
	u.facingIndicator.Dispose()
	u.movesIndicator.Dispose()
	delete(u.sceneState.unitByCell, u.pos)
}

func (u *unitNode) lookTowards(pos dat.CellPos) {
	switch {
	case pos.X > u.pos.X:
		u.facing = 0
	case pos.X < u.pos.X:
		u.facing = 2
	case pos.Y > u.pos.Y:
		u.facing = 1
	default:
		u.facing = 3
	}
	u.updateFacingIndicator()
}

func (u *unitNode) Guard() {
	u.movesLeft = 0
}

func (u *unitNode) MoveTo(pos dat.CellPos) {
	u.lookTowards(pos)
	u.movesLeft--

	delete(u.sceneState.unitByCell, u.pos)

	u.pos = pos
	u.updateSpritePos()

	u.sceneState.unitByCell[u.pos] = u
}

var moveLabels = []string{
	"",
	".",
	"..",
	"...",
	"....",
}

func (u *unitNode) afterTurn() {
	u.movesIndicator.SetText(moveLabels[u.movesLeft])
}

func (u *unitNode) onDamage(dmg int) bool {
	if u.leftoverHP < dmg {
		u.data.Count--
		u.leftoverHP = u.data.Stats.Life

	} else {
		u.leftoverHP -= dmg
	}
	if u.data.Count == 0 {
		u.Dispose()
	}
	u.updateCountLabel()
	return u.data.Count == 0
}

func (u *unitNode) Init(scene *gscene.Scene) {
	u.sceneState.unitByCell[u.pos] = u

	u.updateSpritePos()
	u.updateCountLabel()
	u.updateFacingIndicator()

	game.G.Camera.AddGraphics(u.movesIndicator, 3)
	game.G.Camera.AddGraphics(u.countLabel, 3)
	game.G.Camera.AddGraphics(u.facingIndicator, 2)
	game.G.Camera.AddGraphics(u.sprite, 1)
}

func (u *unitNode) updateCountLabel() {
	u.countLabel.SetText(strconv.Itoa(u.data.Count))
}

func (u *unitNode) updateSpritePos() {
	u.spritePos = u.pos.ToVecPos(true)
}

func (u *unitNode) updateFacingIndicator() {
	u.facingIndicator.SetImage(game.G.Loader.LoadImage(assets.ImageFacingRight + resource.ImageID(u.facing)).Data)
}

func (u *unitNode) Update(delta float64) {

}
