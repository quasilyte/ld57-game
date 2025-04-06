package mapgen

import (
	"github.com/quasilyte/gmath"
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
			EnemyBudget:     6 * dat.Brigands.SquadPrice(),
			EnemyPlacement:  EnemyPlacementRandomSpread,
			PlayerPlacement: PlayerPlacementCorner,
			ForestRatio:     0.6,
			SwampRatio:      0.05,
			Reward:          190,
		}

	case 4:
		ogre := dat.Ogres.CreateUnit()
		ogre.Level = 1
		cfg = Config{
			Width:           8,
			Height:          10,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyHorde,
			EnemyBudget:     3 * dat.GoblinWarriors.SquadPrice(),
			EnemyPlacement:  EnemyPlacementCorner,
			PlayerPlacement: PlayerPlacementCenter,
			ForestRatio:     0.1,
			Reward:          110,
			MandatoryEnemies: []*dat.Unit{
				ogre,
			},
		}

	case 5:
		var enemy EnemyKind
		switch game.G.SelectedArmy {
		case dat.FactionUndead:
			enemy = EnemyUndead
		case dat.FactionHuman:
			enemy = EnemyMercenaries
		default:
			panic("TODO")
		}
		cfg = Config{
			Width:           20,
			Height:          20,
			Mission:         dat.MissionKillAll,
			Enemy:           enemy,
			EnemyBudget:     7 * dat.MercenarySwords.SquadPrice(),
			EnemyPlacement:  EnemyPlacementCenter,
			PlayerPlacement: PlayerPlacementEdges,
			ForestRatio:     0.1,
			SwampRatio:      0.05,
			Reward:          210,
		}

	case 6:
		cfg = Config{
			Width:           20,
			Height:          20,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyHorde,
			EnemyBudget:     9 * dat.MercenarySwords.SquadPrice(),
			EnemyPlacement:  EnemyPlacementEdges,
			PlayerPlacement: PlayerPlacementCenter,
			ForestRatio:     0.1,
			SwampRatio:      0.05,
			Reward:          150,
		}

	default:
		cfg.Width = int(gmath.CeilN(float64(game.G.Rand.IntRange(10, 24)), 2))
		cfg.Height = int(gmath.CeilN(float64(game.G.Rand.IntRange(10, 24)), 2))
		cfg.Mission = dat.MissionKillAll
		cfg.Enemy = gmath.RandElem(&game.G.Rand, []EnemyKind{
			EnemyHorde, EnemyBrigands, EnemyMercenaries, EnemyUndead,
		})
		cfg.EnemyBudget = (7 + (2 * (game.G.Stage - 5))) * dat.MercenarySwords.SquadPrice()
		cfg.EnemyPlacement = gmath.RandElem(&game.G.Rand, []EnemyPlacementKind{
			EnemyPlacementCorner, EnemyPlacementCenter, EnemyPlacementEdges, EnemyPlacementRandomSpread,
		})
		cfg.PlayerPlacement = gmath.RandElem(&game.G.Rand, []PlayerPlacementKind{
			PlayerPlacementCenter, PlayerPlacementCorner, PlayerPlacementEdges,
		})
		cfg.ForestRatio = game.G.Rand.FloatRange(0, 0.3)
		cfg.SwampRatio = game.G.Rand.FloatRange(0, 0.3)
		cfg.Reward = 50 + game.G.Rand.IntRange(0, 250)
	}

	cfg.Stage = game.G.Stage
	game.G.Stage++

	return Generate(cfg)
}
