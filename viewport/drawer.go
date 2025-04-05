package viewport

import graphics "github.com/quasilyte/ebitengine-graphics"

func NewDrawerWithLayers(c *Camera, layers []graphics.SceneLayerDrawer) *graphics.SceneDrawer {
	d := graphics.NewSceneDrawer(layers)
	d.AddCamera(c.c)
	return d
}

func NewDrawer(c *Camera) *graphics.SceneDrawer {
	numLayers := c.numLayers

	layers := make([]graphics.SceneLayerDrawer, numLayers)
	for i := range layers {
		layers[i] = graphics.NewLayer()
	}
	return NewDrawerWithLayers(c, layers)
}
