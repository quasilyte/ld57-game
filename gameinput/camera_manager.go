package gameinput

import (
	input "github.com/quasilyte/ebitengine-input"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/ld57-game/controls"
	"github.com/quasilyte/ld57-game/viewport"
)

type CameraManager struct {
	camera *viewport.Camera
	input  *input.Handler

	cameraPanStartPos gmath.Vec
	cameraPanDragPos  gmath.Vec
}

type CameraManagerConfig struct {
	Camera *viewport.Camera
	Input  *input.Handler
}

func NewCameraManager(config CameraManagerConfig) *CameraManager {
	return &CameraManager{
		camera: config.Camera,
		input:  config.Input,
	}
}

func (m *CameraManager) HandleInput(delta float64) {
	h := m.input
	cameraPanSpeed := 512.0 * delta
	cameraPanBoundary := 2.0
	cam := m.camera

	var cameraPan gmath.Vec
	if h.ActionIsPressed(controls.ActionPanRight) {
		cameraPan.X += cameraPanSpeed
	}
	if h.ActionIsPressed(controls.ActionPanDown) {
		cameraPan.Y += cameraPanSpeed
	}
	if h.ActionIsPressed(controls.ActionPanLeft) {
		cameraPan.X -= cameraPanSpeed
	}
	if h.ActionIsPressed(controls.ActionPanUp) {
		cameraPan.Y -= cameraPanSpeed
	}

	if cameraPan.IsZero() {
		if info, ok := h.JustPressedActionInfo(controls.ActionPanWheel); ok {
			m.cameraPanDragPos = cam.GetOffset()
			m.cameraPanStartPos = info.Pos
		} else if info, ok := h.PressedActionInfo(controls.ActionPanWheel); ok {
			posDelta := m.cameraPanStartPos.Sub(info.Pos)
			newPos := m.cameraPanDragPos.Add(posDelta)
			cam.SetOffset(newPos)
		}
	}

	if cameraPan.IsZero() && cameraPanBoundary != 0 {
		// Mouse cursor can pan the camera too.
		cursor := h.CursorPos()
		windowSize := cam.GetViewportRect()
		if cursor.X >= windowSize.Width()-cameraPanBoundary {
			cameraPan.X += cameraPanSpeed
		}
		if cursor.Y >= windowSize.Height()-cameraPanBoundary {
			cameraPan.Y += cameraPanSpeed
		}
		if cursor.X < cameraPanBoundary {
			cameraPan.X -= cameraPanSpeed
		}
		if cursor.Y < cameraPanBoundary {
			cameraPan.Y -= cameraPanSpeed
		}
	}

	if !cameraPan.IsZero() {
		cam.Pan(cameraPan)
	}
}
