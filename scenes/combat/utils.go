package combat

import "github.com/quasilyte/ld57-game/dat"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(c, other dat.CellPos) int {
	return abs(c.X-other.X) + abs(c.Y-other.Y)
}

func attackFacing(attacker, defender *unitNode) meleeAttackFacing {
	facing := meleeAttackFront
	isDiagonal := abs(attacker.pos.X-defender.pos.X) == 1 && abs(attacker.pos.Y-defender.pos.Y) == 1
	if isDiagonal {
		switch abs(attacker.facing - defender.facing) {
		case 1, 3:
			facing = meleeAttackFlank
		case 0:
			facing = meleeAttackBack
		case 2:
			// Already set to front.
		}
	} else {
		facing = meleeAttackFlank
	}
	return facing
}
