package viewport

import (
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/gsignal"
)

type Camera struct {
	c         *graphics.Camera
	WorldSize gmath.Vec

	scene *gscene.Scene

	cameraToggleProgress float64
	cameraToggleTime     float64
	cameraToggleFrom     gmath.Vec
	cameraToggleTarget   gmath.Vec

	numLayers int

	EventMoved gsignal.Event[gsignal.Void]
}

type CameraConfig struct {
	Scene *gscene.Scene

	NumLayers int

	Rect      gmath.Rect
	WorldSize gmath.Vec

	Shader *graphics.Shader
}

func NewCamera(config CameraConfig) *Camera {
	c := graphics.NewCamera()
	if !config.Rect.IsZero() {
		c.SetViewportRect(config.Rect)
	}

	c.SetBounds(gmath.Rect{Max: config.WorldSize})

	return &Camera{
		c:         c,
		WorldSize: config.WorldSize,
		scene:     config.Scene,
		numLayers: config.NumLayers,
	}
}

func (c *Camera) ToScreenRect(worldRect gmath.Rect) gmath.Rect {
	return gmath.Rect{
		Min: c.ToScreenPos(worldRect.Min),
		Max: c.ToScreenPos(worldRect.Max),
	}
}

func (c *Camera) GetBounds() gmath.Rect {
	return c.c.GetBounds()
}

func (c *Camera) GetViewportRect() gmath.Rect {
	return c.c.GetViewportRect()
}

func (c *Camera) GetPos() gmath.Vec {
	return c.c.GetOffset().Sub(c.c.GetViewportRect().Min)
}

func (c *Camera) GetOffset() gmath.Vec {
	return c.c.GetOffset()
}

func (c *Camera) ToScreenPos(worldPos gmath.Vec) gmath.Vec {
	return worldPos.Sub(c.c.GetOffset())
}

func (c *Camera) ToWorldPos(screenPos gmath.Vec) gmath.Vec {
	return screenPos.Add(c.c.GetOffset())
}

func (c *Camera) SetOffset(pos gmath.Vec) {
	c.cameraToggleTarget = gmath.Vec{}
	if c.c.SetOffset(pos) {
		c.EventMoved.Emit(gsignal.Void{})
	}
}

func (c *Camera) Pan(delta gmath.Vec) {
	if c.c.Pan(delta) {
		c.EventMoved.Emit(gsignal.Void{})
	}
}

func (c *Camera) Center() gmath.Vec {
	return c.c.GetCenterOffset()
}

func (c *Camera) CenterOn(pos gmath.Vec) {
	c.cameraToggleTarget = gmath.Vec{}
	c.centerOn(pos)
}

func (c *Camera) centerOn(pos gmath.Vec) {
	if c.c.SetCenterOffset(pos) {
		c.EventMoved.Emit(gsignal.Void{})
	}
}

func (c *Camera) ToggleTo(pos gmath.Vec, t float64) {
	c.cameraToggleFrom = c.Center()
	c.cameraToggleTarget = pos
	c.cameraToggleProgress = 0
	c.cameraToggleTime = t
}

func (c *Camera) IsDisposed() bool {
	return false
}

func (c *Camera) IsToggling() bool {
	return !c.cameraToggleTarget.IsZero()
}

func (c *Camera) Update(delta float64) {
	if c.cameraToggleTarget.IsZero() {
		return
	}

	const cameraToggleSpeed = 1.0
	c.cameraToggleProgress += delta * cameraToggleSpeed
	centerTarget := c.cameraToggleFrom.LinearInterpolate(c.cameraToggleTarget, c.cameraToggleProgress/c.cameraToggleTime)
	if c.cameraToggleProgress >= c.cameraToggleTime {
		centerTarget = c.cameraToggleTarget
		c.cameraToggleFrom = gmath.Vec{}
		c.cameraToggleTarget = gmath.Vec{}
		c.cameraToggleTime = 0
	}
	c.centerOn(centerTarget)
}

func (c *Camera) WithLayerMask(mask uint64, f func()) {
	oldMask := c.c.GetLayerMask()
	c.c.SetLayerMask(mask)
	f()
	c.c.SetLayerMask(oldMask)
}

func (c *Camera) AddGraphics(o gscene.Graphics, layer int) {
	c.scene.AddGraphics(o, layer)
}
