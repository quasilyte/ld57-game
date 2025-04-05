package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/eui"
)

var G *GlobalContext

type GlobalContext struct {
	SceneManager *gscene.Manager

	Audio sound.System

	WindowSize gmath.Vec

	Loader *resource.Loader

	Rand gmath.Rand

	UI *eui.Builder

	SoundVolume int
}

func (ctx *GlobalContext) NewSprite(id resource.ImageID) *graphics.Sprite {
	s := graphics.NewSprite()
	img := ctx.Loader.LoadImage(id)
	s.SetImage(img.Data)
	return s
}

func (ctx *GlobalContext) PlaySound(id resource.AudioID) {
	resourceID := id
	numSamples := assets.NumSamples(id)
	if numSamples > 0 {
		resourceID += resource.AudioID(ctx.Rand.IntRange(0, numSamples-1))
	}
	ctx.Audio.PlaySound(resourceID)
}
