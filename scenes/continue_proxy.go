package scenes

import (
	"fmt"
	"strconv"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type continueProxyController struct {
}

func NewContinueProxyController() *continueProxyController {
	return &continueProxyController{}
}

func (c *continueProxyController) Init(ctx gscene.InitContext) {
	if game.G.Victory {
		game.G.Gold += game.G.CurrentMap.Reward
		game.G.GoldTotal += game.G.CurrentMap.Reward
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
	} else {
		root.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text: "Game Over",
		}))

		army := "Mercenaries"
		switch game.G.SelectedArmy {
		case dat.FactionUndead:
			army = "Undead"
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

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "CONTINUE",
		Font: assets.FontTiny,
		OnClick: func() {
			if game.G.Victory {
				return
			}
			game.G.Gold = 0
			game.G.Units = game.G.Units[:0]
			game.G.GoldTotal = 0
			game.G.Stage = 0
			game.ChangeScene(NewMainMenuController())
		},
	}))

	game.G.UI.Build(ctx.Scene, topRows)
}

func (c *continueProxyController) Update(delta float64) {}
