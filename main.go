package main

import (
	"bufio"
	"os"

	"bradseiler/quoridor/cmd"
	"bradseiler/quoridor/game"
	"bradseiler/quoridor/intelligence"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	g := cmd.NewGameRunner(cmd.NewHumanPlayer(scanner), intelligence.NewSimpleIntelligence(game.BLACK))
	g.Play()
}
