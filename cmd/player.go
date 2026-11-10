package cmd

import (
	"bufio"
	"fmt"

	"bradseiler/quoridor/game"
)

// humanPlayer is an Agent driven by a person typing moves at the console.
// It prints an "Input move: " prompt and reads two move formats:
//
//	MOVE <direction>   e.g. "MOVE up" — step one square in that direction
//	WALL <row> <col> <v|h>  e.g. "WALL 4 3 v" — place a wall at the given
//	                   coordinates, either vertically (v) or horizontally (h)
//
// Anything else is rejected with an error and re-prompted. If stdin is
// exhausted (EOF or input closed), the player quits the game.
type humanPlayer struct {
	scanner *bufio.Scanner
}

var _ Agent = &humanPlayer{}

// NewHumanPlayer returns an Agent that reads moves one line at a time
// from scanner.
func NewHumanPlayer(scanner *bufio.Scanner) Agent {
	return &humanPlayer{
		scanner: scanner,
	}
}

// NextMove reads from the scanner until the typed line is a valid move,
// then returns it as either a MovePawn (from "MOVE", where the direction
// names the adjacent square to step into) or a MoveWall (from "WALL",
// where "v" and "h" map to VERTICAL and HORIZONTAL orientation). Invalid
// lines are reported and ignored so the player can retry. Returns
// MoveQuit on end of input.
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
					Wall: game.Wall{
						Row:       r,
						Col:       c,
						Direction: game.VERTICAL,
					},
				}
			case "h":
				return game.MoveWall{
					Wall: game.Wall{
						Row:       r,
						Col:       c,
						Direction: game.HORIZONTAL,
					},
				}
			}
		}
		fmt.Println("Invalid input. Try again.")
	}
	return game.MoveQuit{}
}
