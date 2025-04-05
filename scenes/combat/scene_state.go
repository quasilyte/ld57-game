package combat

import (
	"math"

	"github.com/ebitenui/ebitenui"
	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/styles"
)

type sceneState struct {
	units []*unitNode

	scene *gscene.Scene

	uiRoot *ebitenui.UI

	m *dat.Map

	logs []string

	unitByCell map[dat.CellPos]*unitNode

	selectorRotation gmath.Rad

	currentUnitSelector *graphics.Circle
}

func newSceneState() *sceneState {
	state := &sceneState{
		unitByCell: make(map[dat.CellPos]*unitNode),
	}

	r := 18.0
	numDashes := 3.0
	dashLen := 8.0
	state.currentUnitSelector = graphics.NewCircle(r)
	state.currentUnitSelector.SetOutlineColorScale(styles.ColorOrange)
	state.currentUnitSelector.SetVisibility(false)
	state.currentUnitSelector.Rotation = &state.selectorRotation
	state.currentUnitSelector.SetOutlineWidth(1)
	state.currentUnitSelector.SetOutlineDash(dashLen, (r/numDashes*math.Pi)+(r/numDashes*math.Pi-dashLen))

	return state
}

func (s *sceneState) Update(delta float64) {
	s.selectorRotation += gmath.Rad(0.5 * delta)
}
