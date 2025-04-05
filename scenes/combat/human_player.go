package combat

import (
	"image"
	"time"

	"github.com/ebitenui/ebitenui/widget"
	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/controls"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/styles"
)

type humanPlayer struct {
	data *player

	unit *unitNode

	focusedOption *actionOption

	actionTooltipText *widget.Text
	actionTooltip     *widget.Window

	options []actionOption
}

type actionOption struct {
	kind   actionKind
	pos    gmath.Vec
	cell   dat.CellPos
	sprite *graphics.Sprite
}

type actionKind int

const (
	actionMove actionKind = iota
	actionAttack
	actionGuard
	actionShoot
)

func (p *humanPlayer) SetUnit(u *unitNode) {
	p.unit = u

	p.options = p.options[:0]

	{
		p.actionTooltipText = widget.NewText(
			widget.TextOpts.Text("", assets.FontTiny, styles.NormalTextColor.Color()),
			widget.TextOpts.ProcessBBCode(true),
		)
		tt := widget.NewToolTip(
			widget.ToolTipOpts.Content(game.G.UI.NewTooltip(p.actionTooltipText)),
			widget.ToolTipOpts.Delay(time.Second/3),
		)
		content := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.MinSize(32, 32),
				widget.WidgetOpts.ToolTip(tt),
			),
		)
		p.actionTooltip = widget.NewWindow(
			widget.WindowOpts.Contents(content),
			widget.WindowOpts.CloseMode(widget.NONE),
		)
		p.data.sceneState.uiRoot.AddWindow(p.actionTooltip)
	}

	p.options = append(p.options, actionOption{
		pos:  u.spritePos,
		cell: u.pos,
		kind: actionGuard,
	})
	for _, offset := range reachableNeighbors(p.data.sceneState, u) {
		dstPos := u.pos.Add(offset)
		u2 := p.data.sceneState.unitByCell[dstPos]
		if u2 == nil {
			p.options = append(p.options, actionOption{
				pos:  dstPos.ToVecPos(true),
				cell: dstPos,
				kind: actionMove,
			})
			continue
		}
		if u2.team != 0 {
			p.options = append(p.options, actionOption{
				pos:  dstPos.ToVecPos(true),
				cell: dstPos,
				kind: actionAttack,
			})
		}
	}

	if u.data.Stats.MaxRange > 0 {
		reachableRangedTargets(u.sceneState, u, func(target *unitNode) {
			p.options = append(p.options, actionOption{
				pos:  target.pos.ToVecPos(true),
				cell: target.pos,
				kind: actionShoot,
			})
		})
	}

	for i := range p.options {
		o := &p.options[i]
		var imgID resource.ImageID
		switch o.kind {
		case actionMove:
			imgID = assets.ImageActionMove
		case actionAttack:
			imgID = assets.ImageActionAttack
		case actionGuard:
			imgID = assets.ImageActionGuard
		case actionShoot:
			imgID = assets.ImageActionShoot
		}
		if imgID != assets.ImageNone {
			spr := game.G.NewSprite(imgID)
			spr.Pos.Base = &o.pos
			o.sprite = spr
			o.sprite.SetAlpha(0.5)
			game.G.Camera.AddGraphics(spr, 2)
		}
	}
}

func (p *humanPlayer) finishTurn() {
	for _, o := range p.options {
		if o.sprite != nil {
			o.sprite.Dispose()
			o.sprite = nil
		}
	}

	p.focusedOption = nil

	p.data.EventDone.Emit(gsignal.Void{})
}

func (p *humanPlayer) Update(delta float64) {
	cursorPos := game.G.Camera.ToWorldPos(game.G.Input.CursorPos())
	cellPos := dat.CellPos{
		X: int(cursorPos.X) / 32,
		Y: int(cursorPos.Y) / 32,
	}
	found := false
	for i := range p.options {
		o := &p.options[i]
		if o.cell != cellPos {
			continue
		}
		if p.focusedOption != nil {
			p.focusedOption.sprite.SetAlpha(0.5)
		}
		found = true
		p.focusedOption = o
		p.focusedOption.sprite.SetAlpha(1)
		break
	}
	if !found {
		if p.focusedOption != nil {
			p.focusedOption.sprite.SetAlpha(0.5)
			p.focusedOption = nil
		}
	}

	if p.focusedOption == nil {
		p.actionTooltip.SetLocation(image.Rect(-32, -32, 0, 0))
		return
	}

	worldRect := p.focusedOption.pos.BoundsRect(32, 32)
	loc := game.G.Camera.ToScreenRect(worldRect).ToStd()
	p.actionTooltip.SetLocation(loc)

	clicked := game.G.Input.ActionIsJustPressed(controls.ActionClick)
	switch p.focusedOption.kind {
	case actionMove:
		p.actionTooltipText.Label = "LMB: Move"
		if clicked {
			p.unit.MoveTo(p.focusedOption.cell)
			p.finishTurn()
		}
	case actionGuard:
		p.actionTooltipText.Label = "LMB: Guard"
		if clicked {
			p.unit.Guard()
			p.finishTurn()
		}
	case actionShoot:
		p.actionTooltipText.Label = "LMB: Shoot"
		if clicked {
			p.data.EventRangedAttack.Emit(meleeAttackEvent{
				Attacker: p.unit,
				Defender: p.data.sceneState.unitByCell[p.focusedOption.cell],
			})
			p.finishTurn()
		}
	case actionAttack:
		p.actionTooltipText.Label = "LMB: Attack"
		if clicked {
			p.data.EventMeleeAttack.Emit(meleeAttackEvent{
				Attacker: p.unit,
				Defender: p.data.sceneState.unitByCell[p.focusedOption.cell],
			})
			p.finishTurn()
		}
	}

	// if game.G.Input.ActionIsJustPressed(controls.ActionPanRight) {
	// 	pos := p.unit.pos
	// 	pos.X++
	// 	p.unit.MoveTo(pos)
	// 	p.finishTurn()
	// }
}
