package graph

import "math"

/*
Bellman-Ford algorithm finds the shortest path from one vertex (point) of the graph to all the others,
even if the edges (connections between points) have negative weights.
This is in contrast to Dijkstra's algorithm, which does not work with negative weights.
*/

/* usage example
func main() {
	edges := []Edge{
		{from: 0, to: 1, weight: 4},
		{from: 0, to: 2, weight: 5},
		{from: 1, to: 2, weight: -3},
		{from: 2, to: 3, weight: 2},
		{from: 3, to: 1, weight: 1},
	}

	vertices := 4
	start := 0
	distances, ok := bellmanFord(vertices, edges, start)
	if !ok {
		fmt.Println("Обнаружен отрицательный цикл!")
		return
	}

	for i, d := range distances {
		if d == math.MaxInt32 {
			fmt.Printf("До вершины %d нет пути\n", i)
		} else {
			fmt.Printf("Кратчайшее расстояние от %d до %d: %d\n", start, i, d)
		}
	}
}
*/

type EdgeBellmanFord struct {
	from   int
	to     int
	weight int
}

func bellmanFord(vertices int, edges []EdgeBellmanFord, start int) ([]int, bool) {
	dist := make([]int, vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	for i := 0; i < vertices-1; i++ {
		for _, edge := range edges {
			if dist[edge.from] != math.MaxInt32 && dist[edge.from]+edge.weight < dist[edge.to] {
				dist[edge.to] = dist[edge.from] + edge.weight
			}
		}
	}

	for _, edge := range edges {
		if dist[edge.from] != math.MaxInt32 && dist[edge.from]+edge.weight < dist[edge.to] {
			return dist, false
		}
	}

	return dist, true
}
