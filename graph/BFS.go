package graph

import "fmt"

type Graph map[string][]string

/* example usage
func main() {
	graph := Graph{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {"F"},
		"F": {},
	}

	BFS(graph, "A")
}
*/

/*
Breadth—first search is a graph traversal or search algorithm that explores vertices "in breadth",
starting from the original vertex and exploring all its neighbors(level 1) before moving on to the neighbors of the neighbors(level 2), etc(level 3+).
*/
func BFS(graph Graph, start string) {
	visited := make(map[string]bool) // Множество посещённых вершин
	queue := []string{start}         // Очередь для BFS

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}

		visited[current] = true
		fmt.Println("Visited:", current)

		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}
}
