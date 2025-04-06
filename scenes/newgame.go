package scenes

import (
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/mapgen"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type NewGameController struct {
}

func NewNewGameController() *NewGameController {
	return &NewGameController{}
}

func (c *NewGameController) Init(ctx gscene.InitContext) {
	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	p := game.G.UI.NewPanel(eui.PanelConfig{})

	panelRows := eui.NewTopLevelRows()
	p.AddChild(panelRows)

	root.AddChild(p)

	panelRows.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	panelRows.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text:     "HUMANS",
		MinWidth: 200,
		OnClick: func() {
			game.G.SelectedArmy = dat.FactionHuman
			game.G.Units = []*dat.Unit{
				dat.MercenarySwords.CreateUnit(),
				dat.MercenaryCavalry.CreateUnit(),
			}
			m := mapgen.NextStage()
			game.G.CurrentMap = m
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
		},
	}))

	panelRows.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "UNDEAD",
		OnClick: func() {
			game.G.SelectedArmy = dat.FactionUndead
			game.G.Units = []*dat.Unit{
				dat.Zombies.CreateUnit(),
				dat.SkeletalArchers.CreateUnit(),
			}
			m := mapgen.NextStage()
			game.G.CurrentMap = m
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
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
