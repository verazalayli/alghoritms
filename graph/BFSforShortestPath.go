package graph

/* usage example
func main() {
	graph := Graph{
		"A": {"B", "C"},
		"B": {"D", "E"},
		"C": {"F"},
		"D": {},
		"E": {"F"},
		"F": {},
	}

	path := BFSShortestPath(graph, "A", "F")
	if path != nil {
		fmt.Println("Shortest path:", path)
	} else {
		fmt.Println("No path found.")
	}
}
*/

// Optimizations applied:
// 1. Preallocate map and slice capacities to reduce runtime allocations.
// 2. Use a head index instead of slicing the queue to avoid copying.
// 3. Mark 'start' as visited immediately to ensure correctness and avoid re-visits.
func BFSShortestPath(graph Graph, start, target string) []string {
	N := len(graph) // Estimated size for capacity preallocation

	// Preallocate capacity for visited and predecessors maps to avoid resizing
	visited := make(map[string]bool, N)
	predecessors := make(map[string]string, N)

	// Preallocate queue with capacity; use head index instead of slicing
	queue := make([]string, 0, N)
	queue = append(queue, start)
	visited[start] = true // Immediately mark start as visited
	head := 0             // Pointer to the current front of the queue

	for head < len(queue) {
		current := queue[head]
		head++ // Advance the head instead of slicing

		if current == target {
			return reconstructPath(predecessors, start, target)
		}

		// Explore neighbors
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				predecessors[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	// No path found
	return nil
}

// reconstructPath builds the shortest path from 'start' to 'end'
// using the predecessors map generated during BFS.
func reconstructPath(pre map[string]string, start, end string) []string {
	path := []string{end}
	current := end

	// Walk backwards from end to start using the predecessor map
	for current != start {
		prev, ok := pre[current]
		if !ok {
			return nil // Incomplete path, should not happen
		}
		current = prev
		path = append([]string{current}, path...) // Prepend to path
	}

	return path
}
