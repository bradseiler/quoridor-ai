package game

import (
	"fmt"
)

type WallDirection int

const (
	VERTICAL WallDirection = iota
	HORIZONTAL
)

type Wall struct {
	Row       int
	Col       int
	Direction WallDirection
}

func NewWall(row int, col int, dir WallDirection) Wall {
	return Wall{
		Row:       row,
		Col:       col,
		Direction: dir,
	}
}

func (w Wall) Valid(size int) bool {
	return (w.Row >= 0) && (w.Col >= 0) && (w.Row < size-1) && (w.Col < size-1)
}

func adjacent(a, b int) bool {
	return a-b == 1 || b-a == 1
}

func (w Wall) Overlaps(otherWall Wall) bool {
	if w.Row == otherWall.Row && w.Col == otherWall.Col {
		return true
	}
	if w.Row == otherWall.Row &&
		w.Direction == HORIZONTAL &&
		otherWall.Direction == HORIZONTAL &&
		adjacent(w.Col, otherWall.Col) {
		return true
	}
	if w.Col == otherWall.Col &&
		w.Direction == VERTICAL &&
		otherWall.Direction == VERTICAL &&
		adjacent(w.Row, otherWall.Row) {
		return true
	}
	return false
}

func (w Wall) blockedPositions() []positionSlice {
	switch w.Direction {
	case VERTICAL:
		return []positionSlice{
			positionSlice{
				NewPosition(w.Row, w.Col),
				NewPosition(w.Row, w.Col+1),
			},
			positionSlice{
				NewPosition(w.Row+1, w.Col),
				NewPosition(w.Row+1, w.Col+1),
			},
		}
	case HORIZONTAL:
		return []positionSlice{
			positionSlice{
				NewPosition(w.Row, w.Col),
				NewPosition(w.Row+1, w.Col),
			},
			positionSlice{
				NewPosition(w.Row, w.Col+1),
				NewPosition(w.Row+1, w.Col+1),
			},
		}
	default:
		panic(fmt.Errorf("Invalid wall direction."))
	}
}
