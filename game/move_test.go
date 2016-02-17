package game_test

import (
	"testing"

	"bradseiler/quoridor/game"
)

func TestMovePawn(t *testing.T) {
	testCases := []struct {
		direction        game.Direction
		expectedResult   bool
		expectedPosition game.Position
	}{
		{
			direction:        game.UP,
			expectedResult:   true,
			expectedPosition: game.NewPosition(1, 4),
		},
		{
			direction:        game.DOWN,
			expectedResult:   false,
			expectedPosition: game.NewPosition(0, 4),
		},
		{
			direction:        game.LEFT,
			expectedResult:   true,
			expectedPosition: game.NewPosition(0, 3),
		},
		{
			direction:        game.RIGHT,
			expectedResult:   true,
			expectedPosition: game.NewPosition(0, 5),
		},
	}

	for n, tc := range testCases {
		g := game.NewGame()
		move := game.MovePawn{
			Direction: tc.direction,
		}
		if tc.expectedResult != g.Move(move) {
			t.Errorf("(%v) Expected move %v to have result %v", n, move, tc.expectedResult)
		}
		var player *game.Player
		if tc.expectedResult {
			player = g.WaitingPlayer
		} else {
			player = g.ActivePlayer
		}
		if actualPosition := player.PawnPos; tc.expectedPosition != actualPosition {
			t.Errorf("(%v) Expected move %v to end at %v, not %v", n, move, tc.expectedPosition, actualPosition)
		}
	}
}

func TestMoveWall(t *testing.T) {
	moves := []game.MoveWall{
		{
			Wall: game.NewWall(3, 0, game.HORIZONTAL),
		},
		{
			Wall: game.NewWall(3, 2, game.HORIZONTAL),
		},
		{
			Wall: game.NewWall(3, 4, game.HORIZONTAL),
		},
		{
			Wall: game.NewWall(2, 4, game.VERTICAL),
		},
		{
			Wall: game.NewWall(0, 4, game.VERTICAL),
		},
	}

	g := game.NewGame()
	for i, move := range moves {
		expected := i != len(moves)-1
		if expected != g.Move(move) {
			t.Errorf("(%v) Expected move %v to be %v", i, move, expected)
		}
	}
}
