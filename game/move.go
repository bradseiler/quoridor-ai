package game

import "fmt"

type Move interface {
	apply(*Game) bool
}

type MovePawn struct {
	Direction Direction
}

var _ Move = MovePawn{}

func (m MovePawn) apply(g *Game) bool {
	curPos := g.ActivePlayer.PawnPos
	newPos := curPos.Move(m.Direction)
	if !g.Board.Adjacent(curPos, newPos) {
		fmt.Printf("Not adjacent: %v, %v\n", curPos, newPos)
		return false
	}
	if newPos == g.WaitingPlayer.PawnPos {
		return false
	}
	g.ActivePlayer.PawnPos = newPos
	return true
}

type MoveWall struct {
	Wall Wall
}

var _ Move = MoveWall{}

func (m MoveWall) apply(g *Game) bool {
	if g.ActivePlayer.WallsLeft == 0 {
		return false
	}
	success := g.Board.PlaceWall(m.Wall)
	if !success {
		return false
	}
	for _, player := range g.Players() {
		distances := Distances(g.Board, player.Color)
		if _, hasDistance := distances[player.PawnPos]; !hasDistance {
			removed := g.Board.RemoveWall(m.Wall)
			if !removed {
				panic("Failed to remove wall that was just placed.")
			}
			return false
		}
	}
	return true
}

type MoveQuit struct{}

func (m MoveQuit) apply(_ *Game) bool {
	return false
}

var _ Move = MoveQuit{}
