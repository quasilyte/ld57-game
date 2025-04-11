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
	Items        []*dat.ItemStats
	NewItems     []*dat.ItemStats
	ItemLootList []*dat.ItemStats
	Units        []*dat.Unit
	SavedUnits   []*dat.Unit
	Gold         int
	GoldTotal    int
	Stage        int
	SelectedArmy dat.UnitFaction

	SoundVolume int
}

func (ctx *GlobalContext) Reset() {
	ctx.Gold = 100
	ctx.GoldTotal = 0
	ctx.Stage = 0
	ctx.ItemLootList = dat.AllItems
	ctx.Items = ctx.Items[:0]
	ctx.Units = ctx.Units[:0]

	// game.G.Units = []*dat.Unit{
	// 	// dat.OrcWarriors.CreateUnit(),
	// 	// dat.GoblinWarriors.CreateUnit(),

	// 	// dat.MercenaryArchers.CreateUnit(),
	// 	// dat.MercenaryArchers.CreateUnit(),

	// 	dat.OrcCavalry.CreateUnit(),
	// 	// dat.MercenarySwords.CreateUnit(),
	// 	// dat.MercenaryCavalry.CreateUnit(),
	// }
	// m := mapgen.NextStage()
	// game.G.CurrentMap = m
	// game.G.SceneManager.ChangeScene(combat.NewController(combat.Config{
	// 	Map: m,
	// }))
}

func ChangeScene(c gscene.Controller) {
	G.Camera = nil

	G.SceneManager.ChangeScene(c)
}

func (ctx *GlobalContext) NewShader(id resource.ShaderID) *graphics.Shader {
	return graphics.NewShader(ctx.Loader.LoadShader(id).Data)
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
