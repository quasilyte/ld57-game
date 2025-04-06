package scenes

import (
	"fmt"
	"strconv"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/gslices"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/mapgen"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type rosterController struct {
	gridRoot       *widget.Container
	slotGrids      [2]*widget.Container
	goldLabel      *widget.Text
	boughtCounters []int
	uiRoot         *ebitenui.UI
}

func NewRosterController() *rosterController {
	return &rosterController{}
}

func (c *rosterController) updateGoldLabel() {
	c.goldLabel.Label = fmt.Sprintf("Gold: %d", game.G.Gold)
}

func (c *rosterController) Init(ctx gscene.InitContext) {
	topRows := eui.NewTopLevelRows()

	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	p := game.G.UI.NewPanel(eui.PanelConfig{})
	p.AddChild(root)
	topRows.AddChild(p)

	c.goldLabel = game.G.UI.NewText(eui.TextConfig{
		Text: "0",
		Font: assets.FontTiny,
	})
	root.AddChild(c.goldLabel)

	c.updateGoldLabel()

	c.boughtCounters = make([]int, len(game.G.Units))

	numGrids := 1
	if len(game.G.Units) > 7 {
		numGrids++
	}
	for i := 0; i < numGrids; i++ {
		c.slotGrids[i] = widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(5),
					widget.GridLayoutOpts.Spacing(8, 4),
				),
			),
		)
	}
	if numGrids == 1 {
		c.gridRoot = c.slotGrids[0]
	} else {
		c.gridRoot = widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Spacing(16, 4),
				),
			),
		)
		c.gridRoot.AddChild(c.slotGrids[0])
		c.gridRoot.AddChild(c.slotGrids[1])
	}
	root.AddChild(c.gridRoot)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "CONTINUE",
		Font: assets.FontTiny,
		OnClick: func() {
			for _, u := range game.G.Units {
				u.InitialCount = u.Count
			}
			m := mapgen.NextStage()
			game.G.CurrentMap = m
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
		},
	}))

	c.fillSlots()

	c.uiRoot = game.G.UI.Build(ctx.Scene, topRows)
}

func (c *rosterController) fillSlots() {
	for _, grid := range c.slotGrids {
		if grid != nil {
			grid.RemoveChildren()
		}
	}

	unitsGrids := c.slotGrids[:]

	for i, u := range game.G.Units {
		unitGrid := unitsGrids[0]
		if i >= 7 {
			unitGrid = unitsGrids[1]
		}

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
				cost := c.boughtCounters[i] + u.Stats.Cost + (2 * u.Level)
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
				cost := c.boughtCounters[i] + u.Stats.Cost + (2 * u.Level)
				if game.G.Gold < cost {
					return
				}
				game.G.Gold -= cost
				u.Count++
				if c.boughtCounters[i] < 15 {
					c.boughtCounters[i]++
				}
				c.updateGoldLabel()
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

		itemSlots := widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Spacing(4, 4),
				),
			),
		)
		unitGrid.AddChild(itemSlots)

		maxNumArtifacts := 1
		if u.Level >= 5 {
			maxNumArtifacts = 2
		}
		empty := dat.EmptyIcon
		for j := 0; j < 2; j++ {
			tt := game.G.UI.NewText(eui.TextConfig{
				Font:      assets.FontTiny,
				AlignLeft: true,
			})
			icon := widget.NewGraphic(widget.GraphicOpts.Image(empty))

			updateTooltip := func() {
				if j >= maxNumArtifacts {
					tt.Label = "Need unit level 5"
					return
				}
				if u.Items[j] == nil {
					tt.Label = "Empty item slot"
					return
				}
				item := u.Items[j]
				tt.Label = item.Name + "\n\n" + item.Hint
			}
			updateIcon := func() {
				if u.Items[j] == nil {
					icon.Image = empty
					return
				}
				icon.Image = game.G.Loader.LoadImage(u.Items[j].Icon).Data
			}
			updateTooltip()
			updateIcon()

			stack := widget.NewContainer(
				widget.ContainerOpts.Layout(widget.NewStackedLayout()),
			)
			b := game.G.UI.NewButton(eui.ButtonConfig{
				Small:     true,
				Font:      assets.FontTiny,
				MinWidth:  16,
				MinHeight: 22,
				OnClick: func() {
					if u.Items[j] != nil {
						item := u.Items[j]
						u.Items[j] = nil
						game.G.Items = append(game.G.Items, item)
						updateTooltip()
						updateIcon()
						return
					}
					c.openItemSelect(u, j)
				},
			})
			b.GetWidget().ToolTip = widget.NewToolTip(
				widget.ToolTipOpts.Content(
					game.G.UI.NewTooltip(tt),
				),
			)
			b.GetWidget().Disabled = j >= maxNumArtifacts
			stack.AddChild(b)
			stack.AddChild(icon)

			itemSlots.AddChild(stack)
		}

		// unitGrid.AddChild(game.G.UI.NewText(eui.TextConfig{
		// 	AlignLeft: true,
		// 	Text:      fmt.Sprintf("Level %d (%d%%)", u.Level+1, int(100*u.Experience)),
		// 	Font:      assets.FontTiny,
		// }))
	}
}

func (c *rosterController) openItemSelect(u *dat.Unit, itemSlot int) {
	panel := game.G.UI.NewPanel(eui.PanelConfig{
		MinWidth: 128,
	})

	rows := eui.NewTopLevelRows()
	panel.AddChild(rows)

	title := "Select Item"
	rows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text: title,
		Font: assets.FontTiny,
	}))

	grid := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(6),
				widget.GridLayoutOpts.Spacing(8, 4),
			),
		),
	)
	rows.AddChild(grid)

	var window *widget.Window

	for i, item := range game.G.Items {
		stack := widget.NewContainer(
			widget.ContainerOpts.Layout(widget.NewStackedLayout()),
		)
		b := game.G.UI.NewButton(eui.ButtonConfig{
			Small:     true,
			Font:      assets.FontTiny,
			MinWidth:  16,
			MinHeight: 22,
			OnClick: func() {
				u.Items[itemSlot] = item
				game.G.Items = gslices.DeleteAt(game.G.Items, i)
				c.fillSlots()
				window.Close()
			},
		})
		b.GetWidget().ToolTip = widget.NewToolTip(
			widget.ToolTipOpts.Content(
				game.G.UI.NewTooltip(game.G.UI.NewText(eui.TextConfig{
					Text: item.Name + "\n\n" + item.Hint,
					Font: assets.FontTiny,
				})),
			),
		)
		icon := widget.NewGraphic(widget.GraphicOpts.Image(game.G.Loader.LoadImage(item.Icon).Data))
		stack.AddChild(b)
		stack.AddChild(icon)
		grid.AddChild(stack)
	}

	rows.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "OK",
		Font: assets.FontTiny,
		OnClick: func() {
			window.Close()
		},
	}))

	anchor := widget.NewContainer(widget.ContainerOpts.Layout(widget.NewAnchorLayout()))
	anchor.AddChild(panel)

	window = widget.NewWindow(
		widget.WindowOpts.Modal(),
		widget.WindowOpts.Contents(anchor),
		widget.WindowOpts.Location(gmath.Rect{Max: game.G.WindowSize}.ToStd()),
	)

	c.uiRoot.AddWindow(window)
}

func (c *rosterController) Update(delta float64) {}
