package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/eui"
	"github.com/quasilyte/ld57-game/game"
	"github.com/quasilyte/ld57-game/scenes"
)

func main() {
	game.G = &game.GlobalContext{
		SoundVolume: 3,
	}
	game.G.SceneManager = gscene.NewManager()
	game.G.WindowSize = gmath.Vec{
		X: 1920 / 2,
		Y: 1080 / 2,
	}
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	game.G.Loader = resource.NewLoader(audioContext)
	game.G.Loader.OpenAssetFunc = assets.MakeOpenAssetFunc()
	game.G.Rand.SetSeed(time.Now().UnixNano())
	game.G.Audio.Init(audioContext, game.G.Loader)
	game.G.UI = eui.NewBuilder(eui.Config{
		Loader: game.G.Loader,
		Audio:  &game.G.Audio,
	})

	assets.RegisterResources(game.G.Loader)
	game.G.UI.Init()

	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("ld57game")

	game.G.SceneManager.ChangeScene(scenes.NewMainMenuController())

	runner := &gameRunner{}
	if err := ebiten.RunGame(runner); err != nil {
		panic(err)
	}
}

type gameRunner struct{}

func (r *gameRunner) Update() error {
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
