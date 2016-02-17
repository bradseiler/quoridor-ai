package cmd

import (
	"fmt"

	"bradseiler/quoridor/game"
)

type GameRunner struct {
	Game   *game.Game
	Agents [2]Agent
}

func NewGameRunner(whiteAgent, blackAgent Agent) *GameRunner {
	return &GameRunner{
		Game:   game.NewGame(),
		Agents: [2]Agent{whiteAgent, blackAgent},
	}
}

func (gp *GameRunner) Play() {
	gp.PrintGamePrompt()
	for !gp.Game.Complete() {
		agent := gp.Agents[gp.Game.ActivePlayer.Color]
		move := agent.NextMove(gp.Game.Copy())
		if _, ok := move.(game.MoveQuit); ok {
			fmt.Println("Game terminated.")
			break
		}
		if !gp.Game.Move(move) {
			fmt.Printf("Invalid move: %v\n", move)
			panic(nil)
		} else {
			gp.PrintGamePrompt()
		}
	}
	fmt.Printf("Winner: %v\n", gp.Game.Winner())
	fmt.Println(gp.Game)
}

func (gp *GameRunner) PrintGamePrompt() {
	fmt.Println(gp.Game)
}

type Agent interface {
	NextMove(*game.Game) game.Move
}
