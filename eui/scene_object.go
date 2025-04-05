package eui

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/quasilyte/gscene"
)

type SceneObject struct {
	ui *ebitenui.UI
}

func (b *Builder) newSceneObject(root *widget.Container) *SceneObject {
	ui := &ebitenui.UI{
		Container:           root,
		DisableDefaultFocus: true,
	}
	b.currentObject = &SceneObject{
		ui: ui,
	}
	return b.currentObject
}

func (o *SceneObject) IsDisposed() bool { return false }

func (o *SceneObject) Init(scene *gscene.Scene) {}

func (o *SceneObject) Update(delta float64) {
	o.ui.Update()
}

func (o *SceneObject) Draw(dst *ebiten.Image) {
	o.ui.Draw(dst)
}
