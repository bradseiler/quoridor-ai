package intelligence_test

import (
	"testing"

	"bradseiler/quoridor/game"
	"bradseiler/quoridor/intelligence"
)

func TestMoves(t *testing.T) {
	testCases := []struct {
		curPos            game.Position
		expectedDirection game.Direction
	}{
		{
			curPos:            game.NewPosition(5, 5),
			expectedDirection: game.UP,
		},
	}

	for n, tc := range testCases {
		g := game.NewGame()
		g.ActivePlayer.PawnPos = tc.curPos
		cpy := g.Copy()
		intel := intelligence.NewSimpleIntelligence(g.ActivePlayer.Color)
		move := intel.NextMove(cpy)
		movePawn, ok := move.(*game.MovePawn)
		if !ok {
			t.Errorf("(%d) Expected move pawn, got %v", n, move)
		}
		if movePawn.Direction != tc.expectedDirection {
			t.Errorf("(%d) Expected direction %v, got %v", n, tc.expectedDirection, movePawn.Direction)
		}
	}
}
