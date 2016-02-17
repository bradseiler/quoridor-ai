package game

import "fmt"

const (
	GAME_SIZE int = 7
)

type Game struct {
	Board         *Board
	ActivePlayer  *Player
	WaitingPlayer *Player
}

func NewGame() *Game {
	return &Game{
		Board:         NewBoard(GAME_SIZE),
		ActivePlayer:  NewPlayer(WHITE, GAME_SIZE),
		WaitingPlayer: NewPlayer(BLACK, GAME_SIZE),
	}
}

func (g Game) Complete() bool {
	return g.Winner() != NONE
}

func (g Game) Players() []*Player {
	return []*Player{
		g.ActivePlayer,
		g.WaitingPlayer,
	}
}

func (g Game) PlayerForColor(color PlayerColor) *Player {
	for _, p := range g.Players() {
		if p.Color == color {
			return p
		}
	}
	return nil
}

func (g Game) Winner() PlayerColor {
	for _, p := range g.Players() {
		if p.Winner(g.Board) {
			return p.Color
		}
	}
	return NONE
}

func (g *Game) Move(m Move) bool {
	success := m.apply(g)
	if success {
		g.ActivePlayer, g.WaitingPlayer = g.WaitingPlayer, g.ActivePlayer
	}
	return success
}

func (g Game) Copy() *Game {
	return &Game{
		Board:         g.Board.Copy(),
		ActivePlayer:  g.ActivePlayer.Copy(),
		WaitingPlayer: g.WaitingPlayer.Copy(),
	}
}

func (g Game) String() string {
	ret := fmt.Sprintf("Turn: %v\n", g.ActivePlayer.Color)
	ret += g.layoutBoard(NONE) + "\n"
	for _, color := range []PlayerColor{WHITE, BLACK} {
		ret += "-------------------\n\n"
		ret += g.layoutBoard(color) + "\n"
	}
	return ret
}

func (g Game) playerAtPos(pos Position) *Player {
	for _, p := range g.Players() {
		if p.PawnPos == pos {
			return p
		}
	}
	return nil
}

func (g Game) layoutBoard(distColor PlayerColor) string {
	ret := ""
	distances := Distances(g.Board, distColor)
	rowP := NewPosition(g.Board.Size-1, 0)
	prevWalls := make([]bool, g.Board.Size-1)
	for {
		colP := rowP
		i := 0
		for {
			if p := g.playerAtPos(colP); p != nil {
				switch p.Color {
				case WHITE:
					ret += " W "
				case BLACK:
					ret += " B "
				}
			} else if distColor == NONE {
				ret += "   "
			} else if dist, found := distances[colP]; found {
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
