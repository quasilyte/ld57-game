package combat

import (
	"fmt"
	"strconv"

	"github.com/ebitenui/ebitenui/widget"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/gsignal"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/gameinput"
	"github.com/quasilyte/ld57-game/styles"
	"github.com/quasilyte/ld57-game/viewport"
)

const sidePanelWidth = 128

type Controller struct {
	back gscene.Controller

	turnLabel *widget.Text
	turn      int

	turnPending bool

	state  *sceneState
	runner *runner

	cam *gameinput.CameraManager

	m *dat.Map

	unitInfoRows *widget.Container

	players      []*player
	activePlayer *player
	activeUnit   *unitNode

	focusedUnit *unitNode

	scene *gscene.Scene
}

type Config struct {
	Back gscene.Controller

	Map *dat.Map
}

func NewController(config Config) *Controller {
	return &Controller{
		back: config.Back,
		m:    config.Map,
	}
}

func (c *Controller) Init(ctx gscene.InitContext) {
	c.scene = ctx.Scene

	c.state = newSceneState()
	c.state.scene = c.scene
	c.state.m = c.m
	c.runner = &runner{
		sceneState: c.state,
	}

	c.players = []*player{
		{
			team:       0,
			sceneState: c.state,
		},
		{
			team:       1,
			sceneState: c.state,
		},
	}
	c.players[0].impl = &humanPlayer{
		data: c.players[0],
	}
	c.players[1].impl = &computerPlayer{
		data: c.players[1],
	}
	for _, p := range c.players {
		p.EventDone.Connect(nil, c.onPlayerDone)
		p.EventMeleeAttack.Connect(nil, c.onMeleeAttack)
		p.EventRangedAttack.Connect(nil, c.onRangedAttack)
	}

	layers := []graphics.SceneLayerDrawer{
		graphics.NewLayer(),
		graphics.NewLayer(),
		graphics.NewLayer(),
		graphics.NewLayer(),       // World UI
		graphics.NewStaticLayer(), // UI layer
	}
	game.G.Camera = viewport.NewCamera(viewport.CameraConfig{
		Scene: ctx.Scene,
		Rect: gmath.Rect{
			Max: game.G.WindowSize,
		},
		WorldSize: gmath.Vec{
			X: float64(32*c.m.Width) + sidePanelWidth,
			Y: float64(32 * c.m.Height),
		},
		NumLayers: len(layers),
	})
	ctx.SetDrawer(viewport.NewDrawerWithLayers(game.G.Camera, layers))

	game.G.Camera.AddGraphics(c.state.currentUnitSelector, 2)

	i := 1
	for y := 0; y < c.m.Height; y++ {
		for x := 0; x < c.m.Width; x++ {
			spr := game.G.NewSprite(assets.ImageTileGrass)
			spr.SetCentered(false)
			spr.Pos.Offset.X = float64(x * 32)
			spr.Pos.Offset.Y = float64(y * 32)
			if i%2 == 0 {
				spr.SetColorScale(graphics.ColorScale{R: 0.96, G: 0.96, B: 0.96, A: 1})
			}
			i++
			game.G.Camera.AddGraphics(spr, 1)
		}
		i++
	}

	for _, u := range c.m.Units {
		n := newUnitNode(unitNodeConfig{
			Data:  u.Unit,
			Pos:   u.Pos,
			Team:  u.Team,
			State: c.state,
		})
		c.state.units = append(c.state.units, n)
		c.scene.AddObject(n)
	}

	c.cam = gameinput.NewCameraManager(gameinput.CameraManagerConfig{
		Camera: game.G.Camera,
		Input:  game.G.Input,
	})

	c.initUI()

	c.nextTurn()
}

func (c *Controller) initUI() {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionEnd,
			VerticalPosition:   widget.AnchorLayoutPositionStart,
		})),
	)

	c.turnLabel = game.G.UI.NewText(eui.TextConfig{
		Text:  "Turn 000",
		Font:  assets.FontTiny,
		Color: styles.NormalTextColor.Color(),
	})

	panel := game.G.UI.NewPanel(eui.PanelConfig{
		MinWidth:  sidePanelWidth,
		MinHeight: 1080 / 4,
	})
	root.AddChild(panel)

	panelRows := eui.NewPanelRows()
	panel.AddChild(panelRows)

	panelRows.AddChild(c.turnLabel)

	c.unitInfoRows = eui.NewPanelRows()
	panelRows.AddChild(c.unitInfoRows)

	c.updateTurnLabel()

	c.state.uiRoot = game.G.UI.BuildAt(c.scene, root, 4)
}

func (c *Controller) Update(delta float64) {
	if c.turnPending {
		c.state.pause = gmath.ClampMin(c.state.pause-delta, 0)
		if c.state.pause == 0 {
			c.nextTurn()
			c.turnPending = false
		}
	}

	c.handleInput(delta)
	c.state.Update(delta)

	if c.activePlayer != nil {
		c.activePlayer.impl.Update(delta)
	}
}

func (c *Controller) handleInput(delta float64) {
	c.cam.HandleInput(delta)

	cursorPos := game.G.Camera.ToWorldPos(game.G.Input.CursorPos())
	cellPos := dat.CellPos{
		X: int(cursorPos.X) / 32,
		Y: int(cursorPos.Y) / 32,
	}
	hovered := c.state.unitByCell[cellPos]
	if hovered == nil && c.focusedUnit != nil {
		c.focusedUnit = nil
		c.unitInfoRows.RemoveChildren()
		return
	}
	if hovered == c.focusedUnit {
		return
	}
	c.focusedUnit = hovered
	c.unitInfoRows.RemoveChildren()

	c.unitInfoRows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text:     hovered.data.Stats.Name,
		Font:     assets.FontTiny,
		MinWidth: sidePanelWidth - 32,
	}))

	c.unitInfoRows.AddChild(widget.NewGraphic(
		widget.GraphicOpts.Image(hovered.data.Stats.ScaledImage),
	))

	{
		pairs := widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Spacing(2, 2),
					widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
				),
			),
		)

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Units",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       fmt.Sprintf("%d/%d", hovered.data.Count, hovered.data.Stats.MaxCount),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Morale",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		if hovered.broken {
			pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
				Text:       "Broken!",
				Font:       assets.FontTiny,
				AlignRight: true,
			}))
		} else {
			pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
				Text:       fmt.Sprintf("%d%%", int(hovered.morale*100)),
				Font:       assets.FontTiny,
				AlignRight: true,
			}))
		}

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Level",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(hovered.data.Level + 1),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		c.unitInfoRows.AddChild(pairs)
	}

	c.unitInfoRows.AddChild(game.G.UI.NewText(eui.TextConfig{}))

	{
		pairs := widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(4),
					widget.GridLayoutOpts.Spacing(10, 2),
					widget.GridLayoutOpts.Stretch([]bool{false, true, false, true}, nil),
				),
			),
		)

		// ATK DEF
		// ACC CON
		// SPD DIS

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "ATK",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		atkString := strconv.Itoa(hovered.data.Stats.Attack)
		if hovered.data.Stats.RangedAttack > 0 {
			atkString += " [color=ffffee]" + strconv.Itoa(hovered.data.Stats.RangedAttack) + "[/color]"
		}
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       atkString,
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "DEF",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(hovered.data.Stats.Defense),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		accString := strconv.Itoa(gmath.Iround(10 * hovered.data.Stats.MeleeAccuracy))
		if hovered.data.Stats.RangedAccuracy > 0 {
			accString += " [color=ffffee]" + strconv.Itoa(gmath.Iround(10*hovered.data.Stats.RangedAccuracy)) + "[/color]"
		}
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "ACC",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       accString,
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "CON",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(hovered.data.Stats.Life),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "SPD",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(hovered.data.Stats.Speed),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "DIS",
			AlignLeft: true,
			Font:      assets.FontTiny,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(hovered.data.Stats.Morale),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		c.unitInfoRows.AddChild(pairs)
	}

	{
		var traitStrings []string
		for _, t := range hovered.data.Stats.Traits {
			switch t {
			case dat.TraitUnbreakable:
				traitStrings = append(traitStrings, "Inf. Morale")
			case dat.TraitCauseFear:
				traitStrings = append(traitStrings, "Causes Fear")
			case dat.TraitArrowResist:
				traitStrings = append(traitStrings, "Arrow Resist")
			case dat.TraitArrowVulnerability:
				traitStrings = append(traitStrings, "Arrow Weakness")
			case dat.TraitMobile:
				traitStrings = append(traitStrings, "Diag. Moves")
			case dat.TraitFlankingImmune:
				traitStrings = append(traitStrings, "Flanking Resist")
			}
		}
		if len(traitStrings) > 0 {
			c.unitInfoRows.AddChild(game.G.UI.NewText(eui.TextConfig{
				Text: "---",
				Font: assets.FontTiny,
			}))
			traitRows := eui.NewPanelRows()
			for _, s := range traitStrings {
				traitRows.AddChild(game.G.UI.NewText(eui.TextConfig{
					Text:      s,
					Font:      assets.FontTiny,
					AlignLeft: true,
				}))
			}
			c.unitInfoRows.AddChild(traitRows)
		}
	}
}

func (c *Controller) nextTurn() {
	c.turn++
	c.updateTurnLabel()

	c.runner.NextTurn()
	c.onPlayerDone(gsignal.Void{})
}

func (c *Controller) onMeleeAttack(event meleeAttackEvent) {
	event.Attacker.movesLeft = 0
	event.Defender.movesLeft = gmath.ClampMin(event.Defender.movesLeft-1, 0)
	event.Attacker.lookTowards(event.Defender.pos)
	event.Defender.afterTurn()
	c.runner.runMeleeRound(event.Attacker, event.Defender)
}

func (c *Controller) onRangedAttack(event meleeAttackEvent) {
	event.Attacker.movesLeft = 0
	event.Defender.movesLeft = 0
	event.Attacker.lookTowards(event.Defender.pos)
	c.runner.runRangedRound(event.Attacker, event.Defender)
}

func (c *Controller) onPlayerDone(gsignal.Void) {
	if c.activeUnit != nil {
		c.activeUnit.afterTurn()
	}

	nextUnit := c.runner.NextUnit()
	if nextUnit == nil {
		c.turnPending = true
		c.activePlayer = nil
		return
	}

	if nextUnit.broken {
		moveUnitRandomly(c.state, nextUnit)
		c.onPlayerDone(gsignal.Void{})
		return
	}

	c.activePlayer = c.players[nextUnit.team]
	c.activePlayer.impl.SetUnit(nextUnit)
	c.activeUnit = nextUnit

	c.state.currentUnitSelector.SetVisibility(true)
	c.state.currentUnitSelector.Pos.Base = &nextUnit.spritePos
	if nextUnit.team == 0 {
		c.state.currentUnitSelector.SetOutlineColorScale(styles.ColorTeal)
	} else {
		c.state.currentUnitSelector.SetOutlineColorScale(styles.ColorRed)
	}
}

func (c *Controller) updateTurnLabel() {
	c.turnLabel.Label = fmt.Sprintf("Turn %d", c.turn)
}
