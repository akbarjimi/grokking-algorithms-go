package Chapter06

func BreadthFirstSearch(graph map[string][]string, start string, needle string) bool {
	queue := make([]string, 0)
	queue = append(queue, graph[start]...)

	for i := 0; i < len(queue); i++ {
		if queue[i] == needle {
			return true
		} else {
			queue = append(queue, graph[queue[i]]...)
		}
	}

	return false
}
