package assets

import (
	"github.com/quasilyte/bitsweetfont"
)

var (
	FontTiny = bitsweetfont.New1()
	Font1    = bitsweetfont.New1_3()
	Font2    = bitsweetfont.Scale(Font1, 2)
	Font3    = bitsweetfont.Scale(Font1, 3)
)
