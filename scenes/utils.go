package scenes

import (
	"fmt"

	"github.com/ebitenui/ebitenui/widget"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
)

func unitIconWithInfo(u *dat.UnitStats, exp float64, level int) *widget.Graphic {
	icon := widget.NewGraphic(widget.GraphicOpts.Image(game.G.Loader.LoadImage(u.Banner).Data))

	info := eui.NewPanelRows()

	info.AddChild(game.G.UI.NewText(eui.TextConfig{
		Text: u.Name,
		Font: assets.FontTiny,
	}))

	{
		pairs := widget.NewContainer(
			widget.ContainerOpts.Layout(
				widget.NewGridLayout(
					widget.GridLayoutOpts.Columns(2),
					widget.GridLayoutOpts.Spacing(2, 2),
					widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
				),
			),
		)

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			MinWidth:  100,
			Text:      "Class",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       u.Class.String(),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:      "Level",
			Font:      assets.FontTiny,
			AlignLeft: true,
		}))
		pairs.AddChild(game.G.UI.NewText(eui.TextConfig{
			Text:       fmt.Sprintf("%d (%d%%)", level+1, int(100*exp)),
			Font:       assets.FontTiny,
			AlignRight: true,
		}))

		info.AddChild(pairs)
	}

	icon.GetWidget().ToolTip = widget.NewToolTip(
		widget.ToolTipOpts.Content(
			game.G.UI.NewTooltipEx(info),
		),
	)
	return icon
}
