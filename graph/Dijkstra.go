package graph

import (
	"math"
)

/* usage example
func main() {
	graph := Graph{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"C", 2}, {"D", 5}},
		"C": {{"D", 1}},
		"D": {},
	}

	distances := Dijkstra(graph, "A")
	for node, dist := range distances {
		fmt.Printf("Distance from A to %s = %d\n", node, dist)
	}
}
*/

type GraphDijkstra map[int][]Edge

type Edge struct {
	to     int
	weight int
}

func dijkstraSimple(graph GraphDijkstra, start int) map[int]int {
	dist := make(map[int]int)
	visited := make(map[int]bool)
	for node := range graph {
		dist[node] = math.MaxInt64
	}
	dist[start] = 0

	for len(visited) < len(graph) {
		current := -1
		currentDist := math.MaxInt64
		for node, d := range dist {
			if !visited[node] && d < currentDist {
				current = node
				currentDist = d
			}
		}
		if current == -1 {
			break
		}

		for _, edge := range graph[current] {
			if dist[current]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[current] + edge.weight
			}
		}
		visited[current] = true
	}
	return dist
}
