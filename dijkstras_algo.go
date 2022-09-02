package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Item struct {
	vertexIdx int
	distance  int
}
type MinHeap []Item

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Time: O((v+e)*log(v)) | Space: O(v)
func dikstrasAlgorithm(start int, edges [][][]int) []int {
	numberOfVertices := len(edges)
	distances := make([]int, 0, numberOfVertices)
	for range edges {
		distances = append(distances, math.MaxInt32)
	}
	distances[start] = 0

	minDistancesHeap := &MinHeap{}
	heap.Push(minDistancesHeap, Item{vertexIdx: start, distance: 0})

	// visited := make(map[int]bool)

	for minDistancesHeap.Len() > 0 {
		item := heap.Pop(minDistancesHeap).(Item)
		vertex, currentMinDistance := item.vertexIdx, item.distance
		// if _, found := visited[vertex]; found {
		// 	continue
		// }

		// visited[vertex] = true

		for _, edge := range edges[vertex] {
			adjacentVertexIdx, distanceToAdjacentVertex := edge[0], edge[1]

			// if _, found := visited[adjacentVertexIdx]; found {
			// 	continue
			// }

			newPathDistance := currentMinDistance + distanceToAdjacentVertex
			currentAdjacentVertexDistance := distances[adjacentVertexIdx]
			if newPathDistance < currentAdjacentVertexDistance {
				distances[adjacentVertexIdx] = newPathDistance
				heap.Push(minDistancesHeap, Item{vertexIdx: adjacentVertexIdx, distance: newPathDistance})
			}
		}
	}
	return distances
}

func main() {
	edges := [][][]int{
		// for vertex 0
		{
			{1, 2}, {3, 6}, // {vertexIdx, distance} ex {1, 2} distance from vertex 0 to vertex 1 is 2
		},
		// for vertex 1
		{
			{0, 2}, {2, 5},
		},
		// for vertex 2
		{
			{1, 5}, {3, 7}, {4, 6}, {5, 9},
		},
		// for vertex 3
		{
			{0, 6}, {2, 7}, {4, 10},
		},
		// for vertex 4
		{
			{2, 6}, {3, 10}, {5, 6},
		},
		// for vertex 5
		{
			{2, 9}, {4, 6},
		},
	}

	// source/start node is 0
	distances := dikstrasAlgorithm(0, edges)
	fmt.Println("distances: ", distances)

	/* output
	distances:  [0 2 7 6 13 16]
	*/
}
