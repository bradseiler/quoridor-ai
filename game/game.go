package game

import "fmt"

const (
	GAME_SIZE int = 7
)

type Game struct {
	Board   *Board
	Players [2]*Player
	Turn    PlayerColor
}

func NewGame() *Game {
	return &Game{
		Board: NewBoard(GAME_SIZE),
		Players: [2]*Player{
			NewPlayer(WHITE, GAME_SIZE),
			NewPlayer(BLACK, GAME_SIZE),
		},
		Turn: WHITE,
	}
}

func (g Game) Complete() bool {
	return g.Winner() != nil
}

func (g Game) Winner() *PlayerColor {
	var winner PlayerColor
	switch {
	case g.Players[WHITE].PawnPos.Row == g.Board.Size-1:
		winner = WHITE
	case g.Players[BLACK].PawnPos.Row == 0:
		winner = BLACK
	default:
		return nil
	}
	return &winner
}

func (g *Game) Move(m Move) bool {
	success := m.apply(g)
	if success {
		g.Turn = OtherPlayer(g.Turn)
	}
	return success
}

func (g Game) CurrentPlayer() *Player {
	return g.Players[g.Turn]
}

func (g Game) OtherPlayer() *Player {
	return g.Players[OtherPlayer(g.Turn)]
}

func (g Game) String() string {
	ret := fmt.Sprintf("Turn: %v\n", g.Turn)
	ret += g.layoutBoard(nil) + "\n"
	//	for _, color := range []PlayerColor{WHITE, BLACK} {
	//		ret += "-------------------\n\n"
	//		ret += g.layoutBoard(&color) + "\n"
	//	}
	return ret
}

func (g Game) layoutBoard(distColor *PlayerColor) string {
	ret := ""
	distances := Distances(g.Board)
	rowP := NewPosition(g.Board.Size-1, 0)
	prevWalls := make([]bool, g.Board.Size-1)
	for {
		colP := rowP
		i := 0
		for {
			if g.Players[WHITE].PawnPos == colP {
				ret += " W "
			} else if g.Players[BLACK].PawnPos == colP {
				ret += " B "
			} else if distColor == nil {
				ret += "   "
			} else if dist, found := distances[*distColor][colP]; found {
				ret += fmt.Sprintf(" %d ", dist)
			} else {
				ret += "   "
			}
			nextP := colP.Move(RIGHT)
			if !nextP.Valid(g.Board.Size) {
				break
			}
			if g.Board.Adjacent(colP, nextP) {
				ret += ":"
			} else {
				ret += "|"
				prevWalls[i] = !prevWalls[i]
			}
			colP = nextP
			i++
		}
		ret += "\n"
		colP = rowP
		nextRowP := rowP.Move(DOWN)
		if !nextRowP.Valid(g.Board.Size) {
			break
		}
		prevWall := false
		i = 0
		for {
			downP := colP.Move(DOWN)
			if g.Board.Adjacent(colP, downP) {
				ret += " . "
			} else if !prevWall {
				ret += "----"
				prevWall = true
			} else {
				ret += "---"
				prevWall = false
			}
			nextP := colP.Move(RIGHT)
			if !nextP.Valid(g.Board.Size) {
				break
			}
			if !prevWall {
				if prevWalls[i] {
					ret += "|"
				} else {
					ret += " "
				}
			}
			colP = nextP
			i++
		}
		ret += "\n"
		rowP = nextRowP
	}
	return ret
}
