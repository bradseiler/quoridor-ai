package game

type Move interface {
	apply(*Game) bool
}

type MovePawn struct {
	Direction Direction
}

var _ Move = MovePawn{}

func (m MovePawn) apply(g *Game) bool {
	curPos := g.CurrentPlayer().PawnPos
	newPos := curPos.Move(m.Direction)
	if !g.Board.Adjacent(curPos, newPos) {
		return false
	}
	if newPos == g.OtherPlayer().PawnPos {
		return false
	}
	g.CurrentPlayer().PawnPos = newPos
	return true
}

type MoveWall struct {
	Position    Position
	Orientation WallDirection
}

var _ Move = MoveWall{}

func (m MoveWall) apply(g *Game) bool {
	if g.CurrentPlayer().WallsLeft == 0 {
		return false
	}
	success := g.Board.PlaceWall(m.Position, m.Orientation)
	if !success {
		return false
	}
	distances := Distances(g.Board)
	for _, player := range []PlayerColor{WHITE, BLACK} {
		if _, hasDistance := distances[player][g.Players[player].PawnPos]; !hasDistance {
			removed := g.Board.RemoveWall(m.Position, m.Orientation)
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
