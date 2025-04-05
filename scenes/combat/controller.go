package combat

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/gameinput"
	"github.com/quasilyte/ld57-game/viewport"
)

type Controller struct {
	back gscene.Controller

	cam *gameinput.CameraManager

	m *dat.Map

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
			X: float64(32 * c.m.Width),
			Y: float64(32 * c.m.Height),
		},
		NumLayers: len(layers),
	})
	ctx.SetDrawer(viewport.NewDrawerWithLayers(game.G.Camera, layers))

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

	c.cam = gameinput.NewCameraManager(gameinput.CameraManagerConfig{
		Camera: game.G.Camera,
		Input:  game.G.Input,
	})

	c.initUI()
}

func (c *Controller) initUI() {
	topRows := eui.NewTopLevelRows()

	game.G.UI.BuildAt(c.scene, topRows, 4)
}

func (c *Controller) Update(delta float64) {
	c.handleInput(delta)
}

func (c *Controller) handleInput(delta float64) {
	c.cam.HandleInput(delta)
}
