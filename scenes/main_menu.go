package scenes

import (
	"os"
	"runtime"

	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
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
			testMap := make([][]dat.Tile, 20)
			for i := range testMap {
				testMap[i] = make([]dat.Tile, 20)
			}
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: &dat.Map{
					Width:  20,
					Height: 20,
					Tiles:  testMap,
					Units: []dat.DeployedUnit{
						{
							Pos:  dat.CellPos{X: 1, Y: 1},
							Team: 0,
							Unit: &dat.Unit{
								Count: 10,
								Stats: dat.SkeletalWarriors,
							},
						},
						{
							Pos:  dat.CellPos{X: 0, Y: 0},
							Team: 0,
							Unit: &dat.Unit{
								Count: 10,
								Stats: dat.SkeletalArchers,
							},
						},
						{
							Pos:  dat.CellPos{X: 3, Y: 3},
							Team: 0,
							Unit: &dat.Unit{
								Count: 10,
								Stats: dat.UnholyKnights,
							},
						},
						{
							Pos:  dat.CellPos{X: 2, Y: 2},
							Team: 1,
							Unit: &dat.Unit{
								Count: 15,
								Stats: dat.SkeletalWarriors,
							},
						},
						{
							Pos:  dat.CellPos{X: 5, Y: 5},
							Team: 1,
							Unit: &dat.Unit{
								Count: 15,
								Stats: dat.Zombies,
							},
						},
						{
							Pos:  dat.CellPos{X: 4, Y: 4},
							Team: 1,
							Unit: &dat.Unit{
								Count: 15,
								Stats: dat.SkeletalWarriors,
							},
						},
					},
				},
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
