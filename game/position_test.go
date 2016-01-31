package game_test

import (
	"testing"

	"bradseiler/quoridor/game"
	"fmt"
)

func TestScanDirection(t *testing.T) {
	testCases := []struct {
		input    string
		expected game.Direction
	}{
		{
			input:    "up",
			expected: game.UP,
		},
	}

	for _, tc := range testCases {
		var dir game.Direction
		n, err := fmt.Sscanf(tc.input, "%v", &dir)
		if err != nil {
			t.Error(err)
		}
		if n != 1 {
			t.Errorf("Expected one value, got %d", n)
		}
	}
}

func TestValid(t *testing.T) {
	testCases := []struct {
		position game.Position
		expected bool
	}{
		{
			position: game.NewPosition(0, 0),
			expected: true,
		},
		{
			position: game.NewPosition(0, -1),
			expected: false,
		},
		{
			position: game.NewPosition(1, 0),
			expected: false,
		},
		{
			position: game.NewPosition(0, 1),
			expected: false,
		},
		{
			position: game.NewPosition(-1, 0),
			expected: false,
		},
	}

	for n, tc := range testCases {
		if actual := tc.position.Valid(1); actual != tc.expected {
			t.Errorf("(%v) Expected position %v to have validity %v", n, tc.position, tc.expected)
		}
	}
}

func TestMove(t *testing.T) {
	testCases := []struct {
		dir      game.Direction
		expected game.Position
	}{
		{
			dir:      game.UP,
			expected: game.NewPosition(2, 1),
		},
		{
			dir:      game.DOWN,
			expected: game.NewPosition(0, 1),
		},
		{
			dir:      game.RIGHT,
			expected: game.NewPosition(1, 2),
		},
		{
			dir:      game.LEFT,
			expected: game.NewPosition(1, 0),
		},
	}

	for n, tc := range testCases {
		p := game.NewPosition(1, 1)
		actual := p.Move(tc.dir)
		if actual != tc.expected {
			t.Errorf("(%v) Expected moving %v from %v to go to %v, got %v", n, tc.dir, p, tc.expected, actual)
		}
	}
}
