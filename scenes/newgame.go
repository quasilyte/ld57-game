package scenes

import (
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
	"github.com/quasilyte/ld57-game/styles"
)

type NewGameController struct {
}

func NewNewGameController() *NewGameController {
	return &NewGameController{}
}

func (c *NewGameController) Init(ctx gscene.InitContext) {
	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text:  "Choose Army",
		Color: styles.ColorOrange.Color(),
		Font:  assets.Font2,
	}))

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text:     "HUMANS",
		MinWidth: 200,
		OnClick: func() {
		},
	}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "UNDEAD",
		OnClick: func() {
		},
	}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "HORDE",
		OnClick: func() {
		},
	}))

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "BACK",
		OnClick: func() {
			game.G.SceneManager.ChangeScene(NewMainMenuController())
		},
	}))

	game.G.UI.Build(ctx.Scene, root)
}

func (c *NewGameController) Update(delta float64) {}
