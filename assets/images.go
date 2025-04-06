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

		ImageUISmallButtonDisabled: {Path: "image/ui/smallbutton_disabled.png"},
		ImageUISmallButtonIdle:     {Path: "image/ui/smallbutton_idle.png"},
		ImageUISmallButtonHover:    {Path: "image/ui/smallbutton_hover.png"},
		ImageUISmallButtonPressed:  {Path: "image/ui/smallbutton_pressed.png"},

		ImageUIPanel:   {Path: "image/ui/panel.png"},
		ImageUITooltip: {Path: "image/ui/tooltip.png"},

		ImageTileGrass:  {Path: "image/tiles/grass.png"},
		ImageTileForest: {Path: "image/tiles/forest.png"},
		ImageTileSwamp:  {Path: "image/tiles/swamp.png"},
		ImageTileVoid:   {Path: "image/tiles/void.png"},

		ImageActionMove:   {Path: "image/ui/action_move.png"},
		ImageActionGuard:  {Path: "image/ui/action_guard.png"},
		ImageActionAttack: {Path: "image/ui/action_attack.png"},
		ImageActionShoot:  {Path: "image/ui/action_shoot.png"},

		ImageFacingRight: {Path: "image/ui/facing_right.png"},
		ImageFacingDown:  {Path: "image/ui/facing_down.png"},
		ImageFacingLeft:  {Path: "image/ui/facing_left.png"},
		ImageFacingUp:    {Path: "image/ui/facing_up.png"},

		ImageBrigandsBanner:         {Path: "image/banner/brigands.png"},
		ImageAssassinsBanner:        {Path: "image/banner/assassins.png"},
		ImageHumanHalberdsBanner:    {Path: "image/banner/human_halberds.png"},
		ImageHumanWarriorsBanner:    {Path: "image/banner/human_warriors.png"},
		ImageHumanArchersBanner:     {Path: "image/banner/human_archers.png"},
		ImageHumanKnights:           {Path: "image/banner/human_knights.png"},
		ImageZombiesBanner:          {Path: "image/banner/undead_zombies.png"},
		ImageSkeletalWarriorsBanner: {Path: "image/banner/undead_skeletal_warriors.png"},
		ImageSkeletalArchersBanner:  {Path: "image/banner/undead_skeletal_archers.png"},
		ImageUnholyKnights:          {Path: "image/banner/undead_knights.png"},
		ImageHumanMummiesBanner:     {Path: "image/banner/undead_mummy.png"},
		ImageGoblinWarriorBanner:    {Path: "image/banner/goblin_warriors.png"},
		ImageOrcWarriorBanner:       {Path: "image/banner/orc_warriors.png"},
		ImageOrcBoarEliteBanner:     {Path: "image/banner/orc_archer_cavalry.png"},
		ImageOgresBanner:            {Path: "image/banner/ogres.png"},

		ImageArrow: {Path: "image/arrow.png"},

		ImageShaderMask: {Path: "image/noise.png"},
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

	ImageUISmallButtonDisabled
	ImageUISmallButtonIdle
	ImageUISmallButtonHover
	ImageUISmallButtonPressed

	ImageUIPanel
	ImageUITooltip

	ImageActionMove
	ImageActionAttack
	ImageActionGuard
	ImageActionShoot

	ImageFacingRight
	ImageFacingDown
	ImageFacingLeft
	ImageFacingUp

	ImageTileGrass
	ImageTileForest
	ImageTileSwamp
	ImageTileVoid

	ImageBrigandsBanner
	ImageAssassinsBanner
	ImageHumanHalberdsBanner
	ImageHumanWarriorsBanner
	ImageHumanArchersBanner
	ImageHumanKnights
	ImageZombiesBanner
	ImageSkeletalWarriorsBanner
	ImageSkeletalArchersBanner
	ImageUnholyKnights
	ImageHumanMummiesBanner
	ImageOrcWarriorBanner
	ImageOrcBoarEliteBanner
	ImageGoblinWarriorBanner
	ImageOgresBanner

	ImageArrow

	ImageShaderMask
)
