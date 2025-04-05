package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	graphics "github.com/quasilyte/ebitengine-graphics"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/controls"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/scenes"
)

func main() {
	runner := &gameRunner{}

	game.G = &game.GlobalContext{
		SoundVolume: 3,
	}
	game.G.SceneManager = gscene.NewManager()
	game.G.WindowSize = gmath.Vec{
		X: 1920 / 4,
		Y: 1080 / 4,
	}
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	game.G.Loader = resource.NewLoader(audioContext)
	game.G.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc()
	game.G.Rand.SetSeed(time.Now().UnixNano())
	game.G.Audio.Init(audioContext, game.G.Loader)

	runner.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})
	game.G.Input = runner.inputSystem.NewHandler(0, controls.DefaultKeymap())

	game.G.UI = eui.NewBuilder(eui.Config{
		Loader: game.G.Loader,
		Audio:  &game.G.Audio,
	})

	graphics.CompileShaders()

	assets.RegisterResources(game.G.Loader)
	game.G.UI.Init()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("ld57game")

	game.G.SceneManager.ChangeScene(scenes.NewMainMenuController())

	if err := ebiten.RunGame(runner); err != nil {
		panic(err)
	}
}

type gameRunner struct {
	inputSystem input.System
}

func (r *gameRunner) Update() error {
	const delta = 1.0 / 120.0

	r.inputSystem.Update()

	if game.G.Camera != nil {
		game.G.Camera.Update(delta)
	}

	game.G.SceneManager.Update()
	return nil
}

func (r *gameRunner) Draw(screen *ebiten.Image) {
	game.G.SceneManager.Draw(screen)
}

func (g *gameRunner) Layout(_, _ int) (int, int) {
	panic("should never happen")
}

func (g *gameRunner) LayoutF(_, _ float64) (float64, float64) {
	return game.G.WindowSize.X, game.G.WindowSize.Y
}
