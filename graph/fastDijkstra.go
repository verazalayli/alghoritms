package graph

import (
	"container/heap"
	"math"
)

/* usage example
func main() {
	graph := map[int][]Edge{
		0: {{to: 1, weight: 4}, {to: 2, weight: 1}},
		1: {{to: 3, weight: 1}},
		2: {{to: 1, weight: 2}, {to: 3, weight: 5}},
		3: {},
	}

	start := 0
	distances := fastDijkstra(graph, start)
	for node, dist := range distances {
		fmt.Printf("Minimal distance from %d to %d = %d\n", start, node, dist)
	}
}
*/

// ===== Priority quenue =====
type Item struct {
	node     int
	priority int
	index    int // for updating position in heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func fastDijkstra(graph map[int][]Edge, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt64
	}
	dist[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		current := item.node
		currentDist := item.priority

		if currentDist > dist[current] {
			continue
		}

		for _, edge := range graph[current] {
			newDist := dist[current] + edge.weight
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(&pq, &Item{node: edge.to, priority: newDist})
			}
		}
	}

	return dist
}
