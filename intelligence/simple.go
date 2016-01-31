package intelligence

import (
	"bradseiler/quoridor/cmd"
	"bradseiler/quoridor/game"
)

type simpleIntelligence struct {
	color game.PlayerColor
}

func NewSimpleIntelligence(color game.PlayerColor) cmd.Agent {
	return &simpleIntelligence{
		color: color,
	}
}

func (si *simpleIntelligence) NextMove(g *game.Game) game.Move {
	distances := game.Distances(g.Board)[si.color]
	curPos := g.Players[si.color].PawnPos
	minDirection := game.NO_DIRECTION
	minDist := 0
	for _, direction := range []game.Direction{game.UP, game.DOWN, game.LEFT, game.RIGHT} {
		newPos := curPos.Move(direction)
		if !g.Board.Adjacent(curPos, newPos) {
			continue
		}
		if minDirection == game.NO_DIRECTION || minDist > distances[newPos] {
			minDirection = direction
			minDist = distances[newPos]
		}
	}

	return &game.MovePawn{
		Direction: minDirection,
	}
}
