package mapgen

import (
	"github.com/quasilyte/ld57-game/dat"
	"github.com/quasilyte/ld57-game/game"
)

func NextStage() *dat.Map {
	var cfg Config

	stage := game.G.Stage
	switch stage {
	case 0:
		cfg = Config{
			Width:           8,
			Height:          8,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyBrigands,
			EnemyBudget:     1 * dat.Brigands.SquadPrice(),
			EnemyPlacement:  EnemyPlacementEdges,
			PlayerPlacement: PlayerPlacementCenter,
			ForestRatio:     0.1,
			SwampRatio:      0,
			Reward:          50,
		}

	case 1:
		var enemy EnemyKind
		switch game.G.SelectedArmy {
		case dat.FactionUndead:
			enemy = EnemyMercenaries
		case dat.FactionHuman:
			enemy = EnemyUndead
		default:
			panic("TODO")
		}
		cfg = Config{
			Width:           10,
			Height:          10,
			Mission:         dat.MissionKillAll,
			Enemy:           enemy,
			EnemyBudget:     2 * dat.MercenarySwords.SquadPrice(),
			EnemyPlacement:  EnemyPlacementEdges,
			PlayerPlacement: PlayerPlacementCenter,
			ForestRatio:     0.2,
			SwampRatio:      0.1,
			Reward:          100,
		}

	default:
		panic("TODO")
	}

	cfg.Stage = game.G.Stage
	game.G.Stage++

	return Generate(cfg)
}
