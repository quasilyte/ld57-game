package game

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	sound "github.com/quasilyte/ebitengine-sound"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/viewport"
)

var G *GlobalContext

type GlobalContext struct {
	SceneManager *gscene.Manager

	Audio sound.System

	WindowSize gmath.Vec

	Input *input.Handler

	Camera *viewport.Camera

	Loader *resource.Loader

	Rand gmath.Rand

	UI *eui.Builder

	NewContinueProxy func() gscene.Controller
	NewMainMenu      func() gscene.Controller

	Victory      bool
	CurrentMap   *dat.Map
	Units        []*dat.Unit
	Gold         int
	GoldTotal    int
	Stage        int
	SelectedArmy dat.UnitFaction

	SoundVolume int
}

func ChangeScene(c gscene.Controller) {
	G.Camera = nil

	G.SceneManager.ChangeScene(c)
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
