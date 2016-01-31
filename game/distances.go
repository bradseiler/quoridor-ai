package game

import (
	"container/heap"
)

func Distances(b *Board) [2]map[Position]int {
	ret := [2]map[Position]int{
		map[Position]int{},
		map[Position]int{},
	}

	for _, color := range []PlayerColor{WHITE, BLACK} {
		dMap := ret[color]
		heap := newNodeHeap(b, color, dMap)
		populateDistances(b, dMap, heap)
	}
	return ret
}

func newNodeHeap(b *Board, color PlayerColor, dMap map[Position]int) (ret *nodeHeap) {
	ret = new(nodeHeap)
	heap.Init(ret)
	winRow := 0
	if color == WHITE {
		winRow = b.Size - 1
	}
	for col := 0; col < b.Size; col++ {
		p := NewPosition(winRow, col)
		dMap[p] = 0
		heap.Push(ret, newNode(0, p))
	}
	return
}

func populateDistances(b *Board, dMap map[Position]int, nh *nodeHeap) {
	for len(*nh) > 0 {
		node := heap.Pop(nh).(node)
		for _, neighbor := range b.connections[node.position] {
			if _, hasDistance := dMap[neighbor]; !hasDistance {
				dMap[neighbor] = node.distance + 1
				heap.Push(nh, newNode(node.distance+1, neighbor))
			}
		}
	}
}

type node struct {
	distance int
	position Position
}

func newNode(distance int, position Position) node {
	return node{
		distance: distance,
		position: position,
	}
}

type nodeHeap []node

var _ heap.Interface = &nodeHeap{}

func (nh nodeHeap) Len() int {
	return len(nh)
}
func (nh nodeHeap) Less(i, j int) bool {
	return nh[i].distance < nh[j].distance
}
func (nh nodeHeap) Swap(i, j int) {
	nh[i], nh[j] = nh[j], nh[i]
}

func (nh *nodeHeap) Push(x interface{}) {
	*nh = append(*nh, x.(node))
}

func (nh *nodeHeap) Pop() interface{} {
	old := *nh
	n := len(old)
	*nh = old[0 : n-1]
	return old[n-1]
}
