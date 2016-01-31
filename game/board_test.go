package game_test

import (
	"testing"

	"bradseiler/quoridor/game"
)

func TestNewBoard(t *testing.T) {
	board := game.NewBoard(2)
	if !board.Adjacent(game.NewPosition(0, 0), game.NewPosition(0, 1)) {
		t.Errorf("Missing expected adjacency: %v", board)
	}
	if board.Adjacent(game.NewPosition(0, 0), game.NewPosition(1, 1)) {
		t.Errorf("Unexpected adjacency: %v", board)
	}
}

type wall struct {
	pos game.Position
	dir game.WallDirection
}

func newAdjacency(sr, sc, er, ec int) []game.Position {
	return []game.Position{
		game.NewPosition(sr, sc),
		game.NewPosition(er, ec),
	}
}

func TestPlaceAndRemoveWall(t *testing.T) {
	testCases := []struct {
		wallsToPlace      []wall
		adjacenciesToLose [][]game.Position
	}{
		{
			wallsToPlace: []wall{
				{
					pos: game.NewPosition(1, 1),
					dir: game.HORIZONTAL,
				},
			},
			adjacenciesToLose: [][]game.Position{
				newAdjacency(1, 1, 2, 1),
				newAdjacency(1, 2, 2, 2),
			},
		},
		{
			wallsToPlace: []wall{
				{
					pos: game.NewPosition(1, 1),
					dir: game.VERTICAL,
				},
			},
			adjacenciesToLose: [][]game.Position{
				newAdjacency(1, 1, 1, 2),
				newAdjacency(2, 1, 2, 2),
			},
		},
	}

	for _, tc := range testCases {
		board := game.NewBoard(3)
		for _, adj := range tc.adjacenciesToLose {
			if !board.Adjacent(adj[0], adj[1]) {
				t.Errorf("Expected to be adjacent: %v", adj)
			}
		}
		for _, wallToPlace := range tc.wallsToPlace {
			board.PlaceWall(wallToPlace.pos, wallToPlace.dir)
		}
		for _, adj := range tc.adjacenciesToLose {
			if board.Adjacent(adj[0], adj[1]) {
				t.Errorf("Expected not to be adjacent: %v", adj)
			}
		}
		for _, wallToPlace := range tc.wallsToPlace {
			board.RemoveWall(wallToPlace.pos, wallToPlace.dir)
		}
		for _, adj := range tc.adjacenciesToLose {
			if !board.Adjacent(adj[0], adj[1]) {
				t.Errorf("Expected to be adjacent: %v", adj)
			}
		}
	}
}
