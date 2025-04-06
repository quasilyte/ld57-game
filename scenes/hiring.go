package scenes

import (
	"fmt"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/mapgen"
	"github.com/quasilyte/ld57-game/scenes/combat"
	"github.com/quasilyte/ld57-game/scenes/sceneutil"
)

type hiringController struct {
}

func NewHiringController() *hiringController {
	return &hiringController{}
}

func (c *hiringController) Init(ctx gscene.InitContext) {
	topRows := eui.NewTopLevelRows()

	root := eui.NewTopLevelRows()

	ctx.Scene.AddGraphics(sceneutil.NewBackgroundImage(), 0)

	p := game.G.UI.NewPanel(eui.PanelConfig{})
	p.AddChild(root)
	topRows.AddChild(p)

	topRows.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text: "Recruit New Troops",
		Font: assets.FontTiny,
	}))

	goldLabel := game.G.UI.NewText(eui.TextConfig{
		Text: "0",
		Font: assets.FontTiny,
	})
	root.AddChild(goldLabel)

	updateGoldLabel := func() {
		goldLabel.Label = fmt.Sprintf("Gold: %d", game.G.Gold)
	}
	updateGoldLabel()

	unitPicker := gmath.NewRandPicker[*dat.UnitStats](&game.G.Rand)
	unitPicker.AddOption(dat.Brigands, 0.1)
	unitPicker.AddOption(dat.Assassins, 0.1)

	switch game.G.SelectedArmy {
	case dat.FactionHorde:
		if game.G.Stage >= 6 {
			unitPicker.AddOption(dat.OrcWarriors, 0.55)
			unitPicker.AddOption(dat.OrcCavalry, 0.3)
			unitPicker.AddOption(dat.GoblinWarriors, 0.3)
			unitPicker.AddOption(dat.Ogres, 0.3)
		} else {
			unitPicker.AddOption(dat.OrcWarriors, 0.45)
			unitPicker.AddOption(dat.OrcCavalry, 0.15)
			unitPicker.AddOption(dat.GoblinWarriors, 0.6)
			unitPicker.AddOption(dat.Ogres, 0.2)
		}

	case dat.FactionHuman:
		unitPicker.AddOption(dat.MercenarySwords, 0.6)
		unitPicker.AddOption(dat.MercenaryHalberds, 0.45)
		unitPicker.AddOption(dat.MercenaryArchers, 0.5)
		unitPicker.AddOption(dat.MercenaryCavalry, 0.4)

	case dat.FactionUndead:
		unitPicker.AddOption(dat.SkeletalWarriors, 0.6)
		unitPicker.AddOption(dat.Zombies, 0.45)
		unitPicker.AddOption(dat.SkeletalArchers, 0.4)
		unitPicker.AddOption(dat.UnholyKnights, 0.5)
		unitPicker.AddOption(dat.Mummies, 0.35)
		if game.G.Stage >= 6 {
			unitPicker.AddOption(dat.Reapers, 0.3)
		}
	}

	unitGrid := widget.NewContainer(
		widget.ContainerOpts.Layout(
			widget.NewGridLayout(
				widget.GridLayoutOpts.Columns(3),
				widget.GridLayoutOpts.Spacing(8, 4),
			),
		),
	)

	for i := 0; i < 2; i++ {
		u := unitPicker.Pick()
		roll := game.G.Rand.FloatRange(0.6, 1.3)
		price := gmath.Scale(u.SquadPrice(), roll)

		level := 0
		if game.G.Rand.Chance(0.4) {
			level = game.G.Rand.IntRange(1, 2)
			price = gmath.Scale(price, 1.0+(0.25*float64(level)))
		}

		icon := unitIconWithInfo(u, 0, level)
		unitGrid.AddChild(icon)

		var hireButton *widget.Button
		hireButton = game.G.UI.NewButton(eui.ButtonConfig{
			Small:     true,
			MinWidth:  40,
			MinHeight: 22,
			Text:      "HIRE",
			Font:      assets.FontTiny,
			OnClick: func() {
				if game.G.Gold < price {
					return
				}
				game.G.Gold -= price
				updateGoldLabel()
				hired := u.CreateUnit()
				hired.Level = level
				game.G.Units = append(game.G.Units, hired)
				hireButton.Text().Label = "✓"
				hireButton.GetWidget().Disabled = true
			},
		})
		unitGrid.AddChild(hireButton)
		unitGrid.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      fmt.Sprintf("%d coins (%d%%)", price, int(100*roll)),
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
	}

	root.AddChild(unitGrid)

	root.AddChild(game.G.UI.NewText(eui.TextConfig{Text: ""}))

	root.AddChild(game.G.UI.NewButton(eui.ButtonConfig{
		Text: "CONTINUE",
		Font: assets.FontTiny,
		OnClick: func() {
			m := mapgen.NextStage()
			game.G.CurrentMap = m
			game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
				Map: m,
			}))
		},
	}))

	game.G.UI.Build(ctx.Scene, topRows)
}

func (c *hiringController) Update(delta float64) {}
