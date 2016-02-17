package game

type WallDirection int

const (
	VERTICAL WallDirection = iota
	HORIZONTAL
)

type PlayerColor int

const (
	NONE PlayerColor = iota
	WHITE
	BLACK
)

func (color PlayerColor) String() string {
	switch color {
	case WHITE:
		return "White"
	case BLACK:
		return "Black"
	default:
		return ""
	}
}

func OtherPlayer(c PlayerColor) PlayerColor {
	switch c {
	case WHITE:
		return BLACK
	case BLACK:
		return WHITE
	default:
		panic(c)
	}
}
