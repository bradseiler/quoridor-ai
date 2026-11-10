// Package cmd wires the game engine together: it defines the Agent
// interface that any player implementation must satisfy, and GameRunner,
// which drives a game to completion by alternating moves from two agents.
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

// Agent is anything that can choose the next move for its side of the
// board. Implementations include humanPlayer (an interactive console
// player) and the AI players in the intelligence package.
//
// NextMove returns the move the agent wants to make given the current
// game state. The state passed in is a copy the agent may keep or inspect
// freely without affecting the real game. Returning game.MoveQuit signals
// that the agent wants to end the game immediately.
type Agent interface {
	NextMove(*game.Game) game.Move
}
