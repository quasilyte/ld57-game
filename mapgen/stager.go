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
			Reward:          110,
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
			Reward:          160,
		}

	case 2:
		cfg = Config{
			Width:           4,
			Height:          12,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyHorde,
			EnemyBudget:     3 * dat.OrcWarriors.SquadPrice(),
			EnemyPlacement:  EnemyPlacementRandomSpread,
			PlayerPlacement: PlayerPlacementCorner,
			ForestRatio:     0.0,
			SwampRatio:      0.2,
			Reward:          140,
		}

	case 3:
		cfg = Config{
			Width:           12,
			Height:          12,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyBrigands,
			EnemyBudget:     7 * dat.Brigands.SquadPrice(),
			EnemyPlacement:  EnemyPlacementRandomSpread,
			PlayerPlacement: PlayerPlacementCorner,
			ForestRatio:     0.6,
			SwampRatio:      0.05,
			Reward:          190,
		}

	default:
		panic("TODO")
	}

	cfg.Stage = game.G.Stage
	game.G.Stage++

	return Generate(cfg)
}
