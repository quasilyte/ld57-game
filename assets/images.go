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
		ImageActionRotate: {Path: "image/ui/action_rotate.png"},

		ImageFacingRight: {Path: "image/ui/facing_right.png"},
		ImageFacingDown:  {Path: "image/ui/facing_down.png"},
		ImageFacingLeft:  {Path: "image/ui/facing_left.png"},
		ImageFacingUp:    {Path: "image/ui/facing_up.png"},

		ImageItemMagicSword:  {Path: "image/item/magic_sword.png"},
		ImageItemBackstabber: {Path: "image/item/backstabber.png"},
		ImageItemMagicHelmet: {Path: "image/item/magic_helmet.png"},
		ImageItemMagicRing:   {Path: "image/item/ring.png"},
		ImageItemMagicRing2:  {Path: "image/item/ring2.png"},
		ImageItemMagicAxe:    {Path: "image/item/trollbane.png"},
		ImageItemMagicBanner: {Path: "image/item/banner.png"},
		ImageItemMagicBow:    {Path: "image/item/bow.png"},
		ImageItemMagicBracer: {Path: "image/item/bracer.png"},
		ImageItemMagicMace:   {Path: "image/item/mace.png"},
		ImageItemMagicDagger: {Path: "image/item/dagger.png"},

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
		ImageReapersBanner:          {Path: "image/banner/undead_reapers.png"},
		ImageGoblinWarriorBanner:    {Path: "image/banner/goblin_warriors.png"},
		ImageOrcWarriorBanner:       {Path: "image/banner/orc_warriors.png"},
		ImageOrcBoarEliteBanner:     {Path: "image/banner/orc_archer_cavalry.png"},
		ImageOgresBanner:            {Path: "image/banner/ogres.png"},
		ImageTrollBanner:            {Path: "image/banner/troll.png"},

		ImageArrow: {Path: "image/arrow.png"},
		ImageSlash: {Path: "image/slash.png"},

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
	ImageActionRotate

	ImageFacingRight
	ImageFacingDown
	ImageFacingLeft
	ImageFacingUp

	ImageTileGrass
	ImageTileForest
	ImageTileSwamp
	ImageTileVoid

	ImageItemMagicSword
	ImageItemBackstabber
	ImageItemMagicHelmet
	ImageItemMagicRing
	ImageItemMagicRing2
	ImageItemMagicAxe
	ImageItemMagicBanner
	ImageItemMagicBow
	ImageItemMagicBracer
	ImageItemMagicMace
	ImageItemMagicDagger

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
	ImageReapersBanner
	ImageOrcWarriorBanner
	ImageOrcBoarEliteBanner
	ImageGoblinWarriorBanner
	ImageOgresBanner
	ImageTrollBanner

	ImageArrow
	ImageSlash

	ImageShaderMask
)
