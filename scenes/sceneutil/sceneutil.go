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
	s := graphics.NewSprite()
	s.SetImage(screenBg)
	s.SetScaleX(game.G.WindowSize.X)
	s.SetScaleY(game.G.WindowSize.Y)
	return s
}
