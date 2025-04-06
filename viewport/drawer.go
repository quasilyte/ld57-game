package viewport

import graphics "github.com/quasilyte/ebitengine-graphics"

func NewDrawerWithLayers(c *Camera, layers []graphics.SceneLayerDrawer) *graphics.SceneDrawer {
	d := graphics.NewSceneDrawer(layers)
	d.AddCamera(c.c)
	return d
}
