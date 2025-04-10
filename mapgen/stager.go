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
			Reward:          140,
			ItemReward:      true,
		}

	case 1:
		var enemy EnemyKind
		switch game.G.SelectedArmy {
		case dat.FactionUndead:
			enemy = EnemyMercenaries
		case dat.FactionHuman:
			enemy = EnemyUndead
		case dat.FactionHorde:
			if game.G.Rand.Chance(0.6) {
				enemy = EnemyMercenaries
			} else {
				enemy = EnemyUndead
			}
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
			Reward:          150,
		}

	case 2:
		cfg = Config{
			Width:           4,
			Height:          12,
			Mission:         dat.MissionKillAll,
			Enemy:           EnemyHorde,
			EnemyBudget:     (2 * dat.OrcWarriors.SquadPrice()) + dat.GoblinWarriors.SquadPrice(),
			EnemyPlacement:  EnemyPlacementRandomSpread,
			PlayerPlacement: PlayerPlacementCorner,
			ForestRatio:     0.0,
			SwampRatio:      0.2,
			Reward:          170,
			ItemReward:      true,
			MandatoryEnemies: []*dat.Unit{
				dat.GoblinWarriors.CreateUnit(),
			},
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
			Reward:          235,
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
			Reward:          280,
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
		case dat.FactionHorde:
			enemy = EnemyHorde
		default:
			panic("TODO")
		}
		gold := 225
		itemReward := true
		if game.G.Rand.Chance(0.5) {
			gold = 350
			itemReward = false
		}
		cfg = Config{
			Width:           20,
			Height:          20,
			Mission:         dat.MissionKillAll,
			Enemy:           enemy,
			EnemyBudget:     6 * dat.MercenarySwords.SquadPrice(),
			EnemyPlacement:  EnemyPlacementCenter,
			PlayerPlacement: PlayerPlacementEdges,
			ForestRatio:     0.1,
			SwampRatio:      0.05,
			Reward:          gold,
			ItemReward:      itemReward,
		}

	case 6:
		var units []*dat.Unit
		units = append(units, dat.Troll.CreateUnit())
		for i := 0; i < 7; i++ {
			units = append(units, dat.GoblinWarriors.CreateUnit())
		}
		cfg = Config{
			Width:            14,
			Height:           16,
			Mission:          dat.MissionKillAll,
			Enemy:            EnemyBrigands,
			EnemyBudget:      0,
			EnemyPlacement:   EnemyPlacementRandomSpread,
			PlayerPlacement:  PlayerPlacementCorner,
			ForestRatio:      0.05,
			SwampRatio:       0.6,
			Reward:           220,
			MandatoryEnemies: units,
			ItemReward:       true,
		}

	case 7:
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
			Reward:          350,
		}

	default:
		cfg.Width = int(gmath.CeilN(float64(game.G.Rand.IntRange(10, 24)), 2))
		cfg.Height = int(gmath.CeilN(float64(game.G.Rand.IntRange(10, 18)), 2))
		cfg.Mission = dat.MissionKillAll
		cfg.Enemy = gmath.RandElem(&game.G.Rand, []EnemyKind{
			EnemyHorde, EnemyBrigands, EnemyMercenaries, EnemyUndead,
		})
		cfg.EnemyBudget = (6 + (2 * (game.G.Stage - 8))) * dat.MercenarySwords.SquadPrice()
		for try := 0; try <= 3; try++ {
			cfg.EnemyPlacement = gmath.RandElem(&game.G.Rand, []EnemyPlacementKind{
				EnemyPlacementCorner, EnemyPlacementCenter, EnemyPlacementEdges, EnemyPlacementRandomSpread,
			})
			cfg.PlayerPlacement = gmath.RandElem(&game.G.Rand, []PlayerPlacementKind{
				PlayerPlacementCenter, PlayerPlacementCorner, PlayerPlacementEdges,
			})
			isWeird := (cfg.EnemyPlacement == EnemyPlacementEdges && cfg.PlayerPlacement == PlayerPlacementEdges)
			if isWeird {
				continue
			}
			if cfg.EnemyPlacement == EnemyPlacementCorner && cfg.PlayerPlacement == PlayerPlacementCenter {
				cfg.Width = gmath.ClampMin(cfg.Width-4, 12)
				cfg.Height = gmath.ClampMin(cfg.Height-4, 10)
			}
			break
		}
		cfg.ForestRatio = game.G.Rand.FloatRange(0, 0.3)
		cfg.SwampRatio = game.G.Rand.FloatRange(0, 0.3)
		goldRoll := game.G.Rand.FloatRange(0.4, 1.0)
		cfg.ItemReward = goldRoll < 0.65
		cfg.Reward = 50 + gmath.Scale(400, goldRoll)
	}

	cfg.Stage = game.G.Stage
	game.G.Stage++

	return Generate(cfg)
}
