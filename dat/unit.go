package dat

import resource "github.com/quasilyte/ebitengine-resource"

type Unit struct {
	Stats *UnitStats
}

type UnitStats struct {
	Banner resource.ImageID
}
