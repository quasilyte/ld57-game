package combat

import (
	"fmt"

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

	state  *sceneState
	runner *runner

	cam *gameinput.CameraManager

	m *dat.Map

	players      []*player
	activePlayer *player
	activeUnit   *unitNode

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

	c.updateTurnLabel()

	c.state.uiRoot = game.G.UI.BuildAt(c.scene, root, 4)
}

func (c *Controller) Update(delta float64) {
	c.handleInput(delta)
	c.state.Update(delta)
	c.activePlayer.impl.Update(delta)
}

func (c *Controller) handleInput(delta float64) {
	c.cam.HandleInput(delta)
}

func (c *Controller) nextTurn() {
	c.turn++
	c.updateTurnLabel()

	c.runner.NextTurn()
	c.onPlayerDone(gsignal.Void{})
}

func (c *Controller) onMeleeAttack(event meleeAttackEvent) {
	event.Attacker.movesLeft = 0
	event.Defender.movesLeft = 0
	event.Attacker.lookTowards(event.Defender.pos)
	c.runner.runMeleeRound(event.Attacker, event.Defender)
}

func (c *Controller) onPlayerDone(gsignal.Void) {
	if c.activeUnit != nil {
		c.activeUnit.afterTurn()
	}

	nextUnit := c.runner.NextUnit()
	if nextUnit == nil {
		c.nextTurn()
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
