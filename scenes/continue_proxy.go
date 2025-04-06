package scenes

import (
	"fmt"
	"strconv"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/gslices"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type continueProxyController struct {
}

func NewContinueProxyController() *continueProxyController {
	return &continueProxyController{}
}

func (c *continueProxyController) Init(ctx gscene.InitContext) {
	var itemFound *dat.ItemStats

	if game.G.Victory {
		game.G.Gold += game.G.CurrentMap.Reward
		game.G.GoldTotal += game.G.CurrentMap.Reward

		if game.G.CurrentMap.ItemReward {
			if len(game.G.ItemLootList) > 0 {
				i := gmath.RandIndex(&game.G.Rand, game.G.ItemLootList)
				itemFound = game.G.ItemLootList[i]
				game.G.ItemLootList = gslices.DeleteAt(game.G.ItemLootList, i)
				game.G.Items = append(game.G.Items, itemFound)
			}
		}
	}

	topRows := eui.NewTopLevelRows()

	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	p := game.G.UI.NewPanel(eui.PanelConfig{})
	p.AddChild(root)
	topRows.AddChild(p)

	pairs := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(2),
				widget.GridLayoutOpts.Spacing(16, 4),
				widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
			),
		),
	)

	if game.G.Victory {
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			MinWidth:  120,
			Text:      "Mission reward",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       fmt.Sprintf("%d coins", game.G.CurrentMap.Reward),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Gold coffers",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       fmt.Sprintf("%d coins", game.G.Gold),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		if itemFound != nil {
			pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
				Text:      "Item found",
				Font:      assets.FontTiny,
				AlignLeft: true,
			}))
			pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
				Text:       itemFound.Name,
				Font:       assets.FontTiny,
				AlignRight: true,
			}))
		}
	} else {
		root.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text: "Game Over",
		}))

		army := "Mercenaries"
		switch game.G.SelectedArmy {
		case dat.FactionUndead:
			army = "Undead"
		case dat.FactionHorde:
			army = "Horde"
		}
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Army",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       army,
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Missions completed",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       strconv.Itoa(game.G.Stage - 1),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Gold collected",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       fmt.Sprintf("%d coins", game.G.GoldTotal),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))
	}

	root.AddChild(pairs)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	if game.G.Victory {
		root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
			Text: "CONTINUE",
			Font: assets.FontTiny,
			OnClick: func() {
				if len(game.G.Units) < 13 {
					game.ChangeScene(NewHiringController())
					return
				}
				game.ChangeScene(NewRosterController())
			},
		}))
	} else {
		root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
			Text: "RETRY",
			Font: assets.FontTiny,
			OnClick: func() {
				m := map[*dat.Unit]*dat.Unit{}
				for i, u := range game.G.CurrentMap.Units {
					if u.Team == 0 {
						m[u.Unit] = game.G.SavedUnits[i]
					}
					game.G.CurrentMap.Units[i].Unit = game.G.SavedUnits[i]
				}
				for i, u := range game.G.Units {
					game.G.Units[i] = m[u]
				}
				game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
					Map: game.G.CurrentMap,
				}))
			},
		}))

		root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
			Text: "GIVE UP",
			Font: assets.FontTiny,
			OnClick: func() {
				game.G.Reset()
				game.ChangeScene(NewMainMenuController())
			},
		}))
	}

	game.G.UI.Build(ctx.Scene, topRows)
}

func (c *continueProxyController) Update(delta float64) {}
