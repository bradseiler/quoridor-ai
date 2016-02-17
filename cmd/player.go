package cmd

import (
	"bufio"
	"fmt"

	"bradseiler/quoridor/game"
)

type humanPlayer struct {
	scanner *bufio.Scanner
}

var _ Agent = &humanPlayer{}

func NewHumanPlayer(scanner *bufio.Scanner) Agent {
	return &humanPlayer{
		scanner: scanner,
	}
}

func (hp *humanPlayer) NextMove(g *game.Game) game.Move {
	fmt.Print("Input move: ")
	for hp.scanner.Scan() {
		line := hp.scanner.Text()
		var dir game.Direction
		n, err := fmt.Sscanf(line, "MOVE %v", &dir)
		if n == 1 && err == nil {
			return game.MovePawn{
				Direction: dir,
			}
		}
		var r, c int
		var o string
		n, err = fmt.Sscanf(line, "WALL %d %d %s", &r, &c, &o)
		if n == 3 && err == nil {
			switch o {
			case "v":
				return game.MoveWall{
					Position:    game.NewPosition(r, c),
					Orientation: game.VERTICAL,
				}
			case "h":
				return game.MoveWall{
					Position:    game.NewPosition(r, c),
					Orientation: game.HORIZONTAL,
				}
			}
		}
		fmt.Println("Invalid input. Try again.")
	}
	return game.MoveQuit{}
}
