package chapter06

func BreadthFirstSearch(graph map[string][]string, start string, needle string) bool {
	if start == needle {
		return true
	}

	queue := make([]string, 0)
	queue = append(queue, graph[start]...)
	searched := make(map[string]bool)

	for i := 0; i < len(queue); i++ {
		node := queue[i]

		if searched[node] {
			continue
		}

		searched[node] = true

		if node == needle {
			return true
		}

		if neighbors, ok := graph[node]; ok {
			queue = append(queue, neighbors...)
		}
	}

	return false
}
