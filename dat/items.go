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
	ItemRingOfFortitude,
	ItemTrollbane,
	ItemBannerOfWill,
	ItemPointblankBow,
	ItemBracerOfLearning,
	ItemTerrorMace,
	ItemDaggerOfPrecision,
}

var ItemPointblankBow = &ItemStats{
	Name: "Point Blank Bow",
	Hint: "Increased accuracy when attacking a 1-tile dist diagonal target",
	Icon: assets.ImageItemMagicBow,
}

var ItemTrollbane = &ItemStats{
	Name: "Trollbane",
	Hint: "Doubled damage against Trolls and Ogres, but not during retaliation",
	Icon: assets.ImageItemMagicAxe,
}

var ItemBannerOfWill = &ItemStats{
	Name: "Banner of Will",
	Hint: "Bearer always regroups after 1 turn of being broken (routed)",
	Icon: assets.ImageItemMagicBanner,
}

var ItemBracerOfLearning = &ItemStats{
	Name: "Bracer of Learning",
	Hint: "Wearing this makes you earn experience much faster, just don't ask how",
	Icon: assets.ImageItemMagicBracer,
}

var ItemRingOfFortitude = &ItemStats{
	Name: "Ring of Foritude",
	Hint: "Grants +1 HP per unit (+1 effective CON rating)",
	Icon: assets.ImageItemMagicRing2,
}

var ItemRingOfCourage = &ItemStats{
	Name: "Ring of Courage",
	Hint: "Grants complete resist against Cause Fear effect",
	Icon: assets.ImageItemMagicRing,
}

var ItemDaggerOfPrecision = &ItemStats{
	Name: "Dagger of Precision",
	Hint: "Adds 10% melee chance to hit (+1 ACC)",
	Icon: assets.ImageItemMagicDagger,
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

var ItemTerrorMace = &ItemStats{
	Name: "Terror Mace",
	Hint: "Increases morale damage in melee, but not during retaliation",
	Icon: assets.ImageItemMagicMace,
}

var ItemBackstabber = &ItemStats{
	Name: "Backstabber",
	Hint: "Improves attack from behind",
	Icon: assets.ImageItemBackstabber,
}
