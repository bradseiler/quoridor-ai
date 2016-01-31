package game

import (
	"bytes"
	"fmt"
)

type Position struct {
	Row int
	Col int
}

func NewPosition(row, col int) Position {
	return Position{
		Row: row,
		Col: col,
	}
}

type Direction int

const (
	NO_DIRECTION Direction = iota
	UP
	LEFT
	RIGHT
	DOWN
)

func (d *Direction) Scan(state fmt.ScanState, verb rune) error {
	bs, err := state.Token(true, nil)
	if err != nil {
		return err
	}

	bs = bytes.ToLower(bs)
	switch {
	case bytes.Compare(bs, []byte("up")) == 0:
		*d = UP
	case bytes.Compare(bs, []byte("down")) == 0:
		*d = DOWN
	case bytes.Compare(bs, []byte("left")) == 0:
		*d = LEFT
	case bytes.Compare(bs, []byte("right")) == 0:
		*d = RIGHT
	default:
		return fmt.Errorf("Invalid direction: %s", bs)
	}
	return nil
}

type offset struct {
	rowPlus int
	colPlus int
}

func newOffset(up, right int) offset {
	return offset{
		rowPlus: up,
		colPlus: right,
	}
}

func (p Position) addOffset(o offset) Position {
	return Position{
		Row: p.Row + o.rowPlus,
		Col: p.Col + o.colPlus,
	}
}

func (p Position) Move(dir Direction) Position {
	switch dir {
	case UP:
		return p.addOffset(newOffset(1, 0))
	case DOWN:
		return p.addOffset(newOffset(-1, 0))
	case RIGHT:
		return p.addOffset(newOffset(0, 1))
	case LEFT:
		return p.addOffset(newOffset(0, -1))
	default:
		return p
	}
}

func (p Position) Valid(size int) bool {
	return p.Row >= 0 && p.Col >= 0 && p.Row < size && p.Col < size
}
