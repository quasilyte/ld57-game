package combat

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	graphics "github.com/quasilyte/ebitengine-graphics"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/assets"
	"github.com/quasilyte/ld57-game/game"
)

type shaderNode struct {
	t         float64
	timeLimit float64
	sprite    *graphics.Sprite
	pos       gmath.Vec
}

type shaderNodeConfig struct {
	Time   float64
	Flip   bool
	Image  *ebiten.Image
	Shader resource.ShaderID
	Pos    gmath.Vec
}

func newShaderNode(config shaderNodeConfig) *shaderNode {
	n := &shaderNode{
		timeLimit: config.Time,
		pos:       config.Pos,
		sprite:    graphics.NewSprite(),
	}
	n.sprite.SetImage(config.Image)
	n.sprite.Pos.Base = &n.pos
	n.sprite.SetHorizontalFlip(config.Flip)
	n.sprite.Shader = game.G.NewShader(config.Shader)
	n.sprite.Shader.SetFloatValue("Time", float32(n.t))
	noiseOffset := image.Pt(
		game.G.Rand.IntRange(0, 64),
		game.G.Rand.IntRange(0, 64),
	)
	n.sprite.Shader.Texture1 = game.G.Loader.LoadImage(assets.ImageShaderMask).Data.SubImage(image.Rectangle{
		Min: noiseOffset,
		Max: image.Pt(n.sprite.ImageWidth(), n.sprite.ImageHeight()).Add(noiseOffset),
	}).(*ebiten.Image)
	return n
}

func (n *shaderNode) Init(scene *gscene.Scene) {
	scene.AddGraphics(n.sprite, 1)
}

func (n *shaderNode) IsDisposed() bool {
	return n.sprite.IsDisposed()
}

func (n *shaderNode) Dispose() {
	n.sprite.Dispose()
}

func (n *shaderNode) Update(delta float64) {
	n.t += delta
	if n.t >= n.timeLimit {
		n.Dispose()
		return
	}

	n.sprite.Shader.SetFloatValue("Time", float32(n.t))
}
