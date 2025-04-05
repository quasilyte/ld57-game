package sceneutil

import (
	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/styles"
)

func NewBackgroundImage() *graphics.Sprite {
	screenBg := ebiten.NewImage(1, 1)
	screenBg.Fill(styles.ColorBackground.ScaleRGB(0.75).Color())
	// screenBg.Fill(graphics.RGB(0x57839C).ScaleRGB(0.4).Color())
	// screenBg.Fill(graphics.RGB(0x70579C).ScaleRGB(0.4).Color())
	s := graphics.NewSprite()
	s.SetImage(screenBg)
	s.SetScaleX(game.G.WindowSize.X)
	s.SetScaleY(game.G.WindowSize.Y)
	return s
}
