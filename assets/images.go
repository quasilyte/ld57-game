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

		ImageSkeletalWarriorsBanner: {Path: "image/banner/undead_skeletal_warriors.png"},
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

	ImageSkeletalWarriorsBanner
)
