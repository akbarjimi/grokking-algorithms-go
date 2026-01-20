package chapter07

func DepthFirstSearchRecursive(graph map[string][]string, start string, needle string, visited map[string]bool) bool {
	if start == needle {
		return true
	}

	visited[start] = true

	for i := 0; i < len(graph[start]); i++ {
		node := graph[start][i]

		if visited[node] {
			continue
		}

		if DepthFirstSearchRecursive(graph, node, needle, visited) {
			return true
		}
	}

	return false
}
