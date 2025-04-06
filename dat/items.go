package dat

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/ld57-game/assets"
)

var EmptyIcon = ebiten.NewImage(1, 1)

func init() {
	EmptyIcon.Fill(color.Transparent)
}

type ItemStats struct {
	Name       string
	Hint       string
	Icon       resource.ImageID
	ScaledIcon *ebiten.Image
}

var AllItems = []*ItemStats{
	ItemMagicSword,
	ItemBackstabber,
	ItemVengeanceHelm,
	ItemRingOfCourage,
	ItemTrollbane,
}

var ItemTrollbane = &ItemStats{
	Name: "Trollbane",
	Hint: "Doubled damage against Trolls and Ogres, but not during retaliation",
	Icon: assets.ImageItemMagicAxe,
}

var ItemRingOfCourage = &ItemStats{
	Name: "Ring of Courage",
	Hint: "Grants complete resist against Cause Fear effect",
	Icon: assets.ImageItemMagicHelmet,
}

var ItemVengeanceHelm = &ItemStats{
	Name: "Helm of Vengeance",
	Hint: "Increases the number of retaliation attacks",
	Icon: assets.ImageItemMagicHelmet,
}

var ItemMagicSword = &ItemStats{
	Name: "Magic sword",
	Hint: "Ignores half enemy DEF in melee, but not during retaliation",
	Icon: assets.ImageItemMagicSword,
}

var ItemBackstabber = &ItemStats{
	Name: "Backstabber",
	Hint: "Improves attack from behind",
	Icon: assets.ImageItemBackstabber,
}
