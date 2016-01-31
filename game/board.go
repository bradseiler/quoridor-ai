package game

type Board struct {
	walls       map[Position]WallDirection
	connections map[Position]positionSlice
	Size        int
}

func NewBoard(size int) *Board {
	b := &Board{
		walls:       map[Position]WallDirection{},
		connections: map[Position]positionSlice{},
		Size:        size,
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			p := NewPosition(row, col)
			movePs := make([]Position, 0, 4)
			for _, dir := range []Direction{UP, DOWN, LEFT, RIGHT} {
				moveP := p.Move(dir)
				if moveP.Valid(b.Size) {
					movePs = append(movePs, moveP)
				}
			}
			b.connections[p] = movePs
		}
	}

	return b
}

func (b Board) Adjacents(pos Position) []Position {
	return b.connections[pos]
}

func (b Board) Adjacent(pos1 Position, pos2 Position) bool {
	for _, p := range b.connections[pos1] {
		if p == pos2 {
			return true
		}
	}
	return false
}

type positionSlice []Position

func (ps positionSlice) add(posToAdd Position) positionSlice {
	return append(ps, posToAdd)
}

func (ps positionSlice) remove(posToRemove Position) positionSlice {
	last := len(ps) - 1
	for i, pos := range ps {
		if pos == posToRemove {
			if i != last {
				ps[i] = ps[last]
			}
			ps = ps[:last]
		}
	}
	return ps
}

func (b *Board) PlaceWall(pos Position, dir WallDirection) bool {
	if !pos.Valid(b.Size - 1) {
		return false
	}

	if _, found := b.walls[pos]; found {
		return false
	}

	adjacentWalls := b.adjacentWalls(pos, dir)
	for _, pos := range adjacentWalls {
		if adj, found := b.walls[pos]; found && adj == dir {
			return false
		}
	}

	for _, block := range blockedConnections(pos, dir) {
		b.connections[block[0]] = b.connections[block[0]].remove(block[1])
		b.connections[block[1]] = b.connections[block[1]].remove(block[0])
	}

	b.walls[pos] = dir
	return true
}

func (b *Board) RemoveWall(pos Position, dir WallDirection) bool {
	if !pos.Valid(b.Size - 1) {
		return false
	}

	if w, found := b.walls[pos]; !found || w != dir {
		return false
	}

	delete(b.walls, pos)

	for _, block := range blockedConnections(pos, dir) {
		b.connections[block[0]] = b.connections[block[0]].add(block[1])
		b.connections[block[1]] = b.connections[block[1]].add(block[0])
	}

	return true
}

func (b Board) adjacentWalls(pos Position, dir WallDirection) positionSlice {
	switch dir {
	case HORIZONTAL:
		return []Position{
			pos.Move(LEFT),
			pos.Move(RIGHT),
		}
	case VERTICAL:
		return []Position{
			pos.Move(UP),
			pos.Move(DOWN),
		}
	}
	return nil
}

func blockedConnections(pos Position, dir WallDirection) []positionSlice {
	var blocksByOffset [][]offset
	switch dir {
	case HORIZONTAL:
		blocksByOffset = [][]offset{
			[]offset{
				newOffset(0, 0),
				newOffset(1, 0),
			},
			[]offset{
				newOffset(0, 1),
				newOffset(1, 1),
			},
		}
	case VERTICAL:
		blocksByOffset = [][]offset{
			[]offset{
				newOffset(0, 0),
				newOffset(0, 1),
			},
			[]offset{
				newOffset(1, 0),
				newOffset(1, 1),
			},
		}
	default:
		panic(dir)
	}

	ret := make([]positionSlice, 2)
	for i, block := range blocksByOffset {
		ret[i] = []Position{
			pos.addOffset(block[0]),
			pos.addOffset(block[1]),
		}
	}
	return ret
}
