package scenes

import (
	"fmt"
	"strconv"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/mapgen"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type rosterController struct {
}

func NewRosterController() *rosterController {
	return &rosterController{}
}

func (c *rosterController) Init(ctx gscene.InitContext) {
	topRows := eui.NewTopLevelRows()

	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	p := game.G.UI.NewPanel(eui.PanelConfig{})
	p.AddChild(root)
	topRows.AddChild(p)

	goldLabel := game.G.UI.NewText(eui.TextConfig{
		Text: "0",
		Font: assets.FontTiny,
	})
	root.AddChild(goldLabel)

	updateGoldLabel := func() {
		goldLabel.Label = fmt.Sprintf("Gold: %d", game.G.Gold)
	}
	updateGoldLabel()

	boughtCounters := make([]int, len(game.G.Units))

	unitGrid := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(4),
				widget.GridLayoutOpts.Spacing(8, 4),
			),
		),
	)

	for i, u := range game.G.Units {
		unitGrid.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text: strconv.Itoa(i + 1),
			Font: assets.FontTiny,
		}))

		countLabel := game.G.UI.NewText(eui.TextConfig{
			Font:        assets.FontTiny,
			AlignLeft:   true,
			ForceBBCode: true,
		})
		updateCountLabel := func() {
			s := strconv.Itoa(u.Count)
			if u.Count < u.Stats.MaxCount {
				s = "[color=cd3e56]" + s + "[/color]"
			}
			s += "/" + strconv.Itoa(u.Stats.MaxCount)
			countLabel.Label = s
		}
		updateCountLabel()

		icon := unitIconWithInfo(u.Stats, u.Experience, u.Level)
		unitGrid.AddChild(icon)

		hireCostTooltip := game.G.UI.NewText(eui.TextConfig{
			Text: "?",
			Font: assets.FontTiny,
		})
		updateHireTooltip := func() {
			if u.Stats.MaxCount == u.Count {
				hireCostTooltip.Label = "Unit is full"
			} else {
				cost := boughtCounters[i] + u.Stats.Cost + (2 * u.Level)
				hireCostTooltip.Label = fmt.Sprintf("Cost: %d", cost)
			}
		}
		updateHireTooltip()

		hireButton := game.G.UI.NewButton(eui.ButtonConfig{
			Small:     true,
			Text:      "+",
			MinWidth:  22,
			MinHeight: 22,
			Font:      assets.FontTiny,
			OnClick: func() {
				if u.Count >= u.Stats.MaxCount {
					return
				}
				cost := boughtCounters[i] + u.Stats.Cost + (2 * u.Level)
				if game.G.Gold < cost {
					return
				}
				game.G.Gold -= cost
				u.Count++
				boughtCounters[i]++
				updateGoldLabel()
				updateCountLabel()
				updateHireTooltip()
			},
		})
		unitGrid.AddChild(hireButton)

		hireButton.GetWidget().ToolTip = widget.NewToolTip(
			widget.ToolTipOpts.Content(
				game.G.UI.NewTooltip(hireCostTooltip),
			),
		)

		unitGrid.AddChild(countLabel)

		// unitGrid.AddChild(game.G.UI.NewText(eui.TextConfig{
		// 	AlignLeft: true,
		// 	Text:      fmt.Sprintf("Level %d (%d%%)", u.Level+1, int(100*u.Experience)),
		// 	Font:      assets.FontTiny,
		// }))
	}

	root.AddChild(unitGrid)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "CONTINUE",
		Font: assets.FontTiny,
		OnClick: func() {
			for _, u := range game.G.Units {
				u.InitialCount = u.Count
			}
			if len(game.G.Units) < 8 {
				game.ChangeScene(NewHiringController())
				return
			}
			m := mapgen.NextStage()
			game.G.CurrentMap = m
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
		},
	}))

	game.G.UI.Build(ctx.Scene, topRows)
}

func (c *rosterController) Update(delta float64) {}
