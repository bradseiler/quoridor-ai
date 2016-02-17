package game

type Board struct {
	Walls       []Wall
	Size        int
	connections map[Position]positionSlice
}

func NewBoard(size int) *Board {
	b := &Board{
		Size: size,
	}

	b.updateConnections()
	return b
}

func (b *Board) updateConnections() {
	connections := map[Position]positionSlice{}

	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			p := NewPosition(row, col)
			movePs := make([]Position, 0, 4)
			for _, dir := range []Direction{UP, DOWN, LEFT, RIGHT} {
				moveP := p.Move(dir)
				if moveP.Valid(b.Size) {
					movePs = append(movePs, moveP)
				}
			}
			connections[p] = movePs
		}
	}

	for _, wall := range b.Walls {
		for _, posPairs := range wall.blockedPositions() {
			connections[posPairs[0]].remove(posPairs[1])
			connections[posPairs[1]].remove(posPairs[0])
		}
	}
	b.connections = connections
}

func (b Board) Copy() *Board {
	ret := &Board{
		Walls: make([]Wall, len(b.Walls)),
		Size:  b.Size,
	}
	for i, wall := range b.Walls {
		ret.Walls[i] = wall
	}
	ret.updateConnections()
	return ret
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

func (b *Board) PlaceWall(wall Wall) bool {
	if !wall.Valid(b.Size) {
		return false
	}

	for _, existingWall := range b.Walls {
		if wall.Overlaps(existingWall) {
			return false
		}
	}

	for _, block := range wall.blockedPositions() {
		b.connections[block[0]] = b.connections[block[0]].remove(block[1])
		b.connections[block[1]] = b.connections[block[1]].remove(block[0])
	}

	b.Walls = append(b.Walls, wall)
	return true
}

func (b *Board) RemoveWall(wall Wall) bool {
	if !wall.Valid(b.Size) {
		return false
	}

	for i, existingWall := range b.Walls {
		if existingWall == wall {
			b.Walls = append(b.Walls[:i], b.Walls[i+1:len(b.Walls)]...)
			for _, block := range wall.blockedPositions() {
				b.connections[block[0]] = b.connections[block[0]].add(block[1])
				b.connections[block[1]] = b.connections[block[1]].add(block[0])
			}
		}
	}

	return true
}
