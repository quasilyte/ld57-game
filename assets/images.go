package assets

import (
	resource "github.com/quasilyte/ebitengine-resource"

	_ "image/png"
)

func registerImageResources(loader *resource.Loader) {
	resources := map[resource.ImageID]resource.ImageInfo{
		ImageUIButtonIdle:    {Path: "image/ui/button_idle.png"},
		ImageUIButtonHover:   {Path: "image/ui/button_hover.png"},
		ImageUIButtonPressed: {Path: "image/ui/button_pressed.png"},
		ImageUIPanel:         {Path: "image/ui/panel.png"},
		ImageUITooltip:       {Path: "image/ui/tooltip.png"},

		ImageTileGrass: {Path: "image/tiles/grass.png"},

		ImageActionMove:   {Path: "image/ui/action_move.png"},
		ImageActionGuard:  {Path: "image/ui/action_guard.png"},
		ImageActionAttack: {Path: "image/ui/action_attack.png"},

		ImageFacingRight: {Path: "image/ui/facing_right.png"},
		ImageFacingDown:  {Path: "image/ui/facing_down.png"},
		ImageFacingLeft:  {Path: "image/ui/facing_left.png"},
		ImageFacingUp:    {Path: "image/ui/facing_up.png"},

		ImageZombiesBanner:          {Path: "image/banner/undead_zombies.png"},
		ImageSkeletalWarriorsBanner: {Path: "image/banner/undead_skeletal_warriors.png"},
		ImageUnholyKnights:          {Path: "image/banner/undead_knights.png"},
	}

	for id, info := range resources {
		loader.ImageRegistry.Set(id, info)
		loader.LoadImage(id)
	}
}

const (
	ImageNone resource.ImageID = iota

	ImageUIButtonIdle
	ImageUIButtonHover
	ImageUIButtonPressed

	ImageUIPanel
	ImageUITooltip

	ImageActionMove
	ImageActionAttack
	ImageActionGuard

	ImageFacingRight
	ImageFacingDown
	ImageFacingLeft
	ImageFacingUp

	ImageTileGrass

	ImageZombiesBanner
	ImageSkeletalWarriorsBanner
	ImageUnholyKnights
)
