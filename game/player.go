package game

type Player struct {
	PawnPos   Position
	WallsLeft int
}

func NewPlayer(color PlayerColor, size int) *Player {
	var pawnPos Position
	switch color {
	case WHITE:
		pawnPos = NewPosition(0, size/2)
	case BLACK:
		pawnPos = NewPosition(size-1, size/2)
	}
	return &Player{
		PawnPos:   pawnPos,
		WallsLeft: size + 1,
	}
}
