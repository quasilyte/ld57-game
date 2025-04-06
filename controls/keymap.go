package controls

import (
	input "github.com/quasilyte/ebitengine-input"
)

const (
	ActionUnknown input.Action = iota

	ActionPanUp
	ActionPanDown
	ActionPanLeft
	ActionPanRight
	ActionPanWheel

	ActionGuard

	ActionClick
)

func DefaultKeymap() input.Keymap {
	return input.Keymap{
		ActionPanUp:    {input.KeyUp},
		ActionPanDown:  {input.KeyDown},
		ActionPanLeft:  {input.KeyLeft},
		ActionPanRight: {input.KeyRight},
		ActionPanWheel: {input.KeyMouseMiddle},
		ActionClick:    {input.KeyMouseLeft},
		ActionGuard:    {input.KeyTab},
	}
}
