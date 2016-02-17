package game

type Player struct {
	PawnPos   Position
	WallsLeft int
	Color     PlayerColor
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
		Color:     color,
	}
}

func (p Player) Winner(b *Board) bool {
	switch p.Color {
	case WHITE:
		return p.PawnPos.Row == b.Size-1
	case BLACK:
		return p.PawnPos.Row == 0
	}
	return false
}

func (p Player) Copy() *Player {
	return &p
}
