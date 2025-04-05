package scenes

import (
	"os"
	"runtime"

	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/mapgen"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
	"github.com/quasilyte/ld57-game/styles"
)

type mainMenuController struct {
}

func NewMainMenuController() *mainMenuController {
	return &mainMenuController{}
}

func (c *mainMenuController) Init(ctx gscene.InitContext) {
	topRows := eui.NewTopLevelRows()

	// root := topRows
	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	topRows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text:  "Super Game",
		Color: styles.ColorOrange.Color(),
		Font:  assets.Font2,
	}))

	p := game.G.UI.NewPanel(eui.PanelConfig{})
	p.AddChild(root)
	topRows.AddChild(p)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text:     "PLAY",
		MinWidth: 200,
		OnClick: func() {
			m := mapgen.Generate(mapgen.Config{
				Width:       8,
				Height:      8,
				ForestRatio: 0.2,
				SwampRatio:  0.4,
			})
			m.Units = []dat.DeployedUnit{
				{
					Pos:  dat.CellPos{X: 3, Y: 3},
					Team: 0,
					Unit: &dat.Unit{
						Count: 15,
						Stats: dat.MercenarySwords,
					},
				},
				// {
				// 	Pos:  dat.CellPos{X: 1, Y: 0},
				// 	Team: 0,
				// 	Unit: &dat.Unit{
				// 		Count: 10,
				// 		Stats: dat.MercenaryCavalry,
				// 	},
				// },
				// {
				// 	Pos:  dat.CellPos{X: 2, Y: 0},
				// 	Team: 0,
				// 	Unit: &dat.Unit{
				// 		Count: 15,
				// 		Stats: dat.MercenarySwords,
				// 	},
				// },
				// {
				// 	Pos:  dat.CellPos{X: 1, Y: 1},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 15,
				// 		Stats: dat.MercenarySwords,
				// 	},
				// },
				// {
				// 	Pos:  dat.CellPos{X: 2, Y: 1},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 15,
				// 		Stats: dat.MercenarySwords,
				// 	},
				// },

				// {
				// 	Pos:  dat.CellPos{X: 4, Y: 5},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 15,
				// 		Stats: dat.SkeletalArchers,
				// 	},
				// },
				{
					Pos:  dat.CellPos{X: 4, Y: 6},
					Team: 1,
					Unit: &dat.Unit{
						Count: 5,
						Stats: dat.Mummies,
					},
				},
				// {
				// 	Pos:  dat.CellPos{X: 4, Y: 4},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 20,
				// 		Stats: dat.SkeletalWarriors,
				// 	},
				// },
				// {
				// 	Pos:  dat.CellPos{X: 3, Y: 5},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 10,
				// 		Stats: dat.Zombies,
				// 	},
				// },
				// {
				// 	Pos:  dat.CellPos{X: 4, Y: 4},
				// 	Team: 1,
				// 	Unit: &dat.Unit{
				// 		Count: 15,
				// 		Stats: dat.SkeletalWarriors,
				// 	},
				// },
			}
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
			// game.G.SceneManager.ChangeScene(NewNewGameController())
		},
	}))

	{
		settings := game.G.UI.NewButton(eui.ButtonConfig{
			Text: "SETTINGS",
			OnClick: func() {
			},
		})
		root.AddChild(settings)
	}

	{
		settings := game.G.UI.NewButton(eui.ButtonConfig{
			Text: "CREDITS",
			OnClick: func() {
			},
		})
		root.AddChild(settings)
	}

	if runtime.GOARCH != "wasm" {
		root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

		root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
			Text: "EXIT",
			OnClick: func() {
				os.Exit(0)
			},
		}))
	}

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	topRows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text:  "Ludum Dare 57 compo build 1",
		Font:  assets.FontTiny,
		Color: styles.ColorOrange.Color(),
	}))
	topRows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text:  "Made with Ebitengine",
		Font:  assets.FontTiny,
		Color: styles.ColorOrange.Color(),
	}))

	game.G.UI.Build(ctx.Scene, topRows)
}

func (c *mainMenuController) Update(delta float64) {}
