package chapter07

import (
	"fmt"
	"sort"
)

func doDfs(
	current string,
	graph map[string][]string,
	visited map[string]bool,
	result *[]string,
	level int,
) {
	// just for printing depth
	space := ""
	for i := 0; i < level; i++ {
		space += "  "
	}

	fmt.Println(space+"enter:", current)

	visited[current] = true

	neighbors := graph[current]
	for i := 0; i < len(neighbors); i++ {
		n := neighbors[i]

		if visited[n] == false {
			fmt.Println(space+"go to:", n)
			doDfs(n, graph, visited, result, level+1)
		} else {
			fmt.Println(space+n, "already visited")
		}
	}

	*result = append(*result, current)
	fmt.Println(space+"exit:", current, "result now:", *result)
}

func TopologicalSort(graph map[string][]string) []string {
	visited := make(map[string]bool)
	result := []string{}

	// map order is random so sort keys
	keys := []string{}
	for k := range graph {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("start topo sort")

	for i := 0; i < len(keys); i++ {
		node := keys[i]
		if visited[node] == false {
			doDfs(node, graph, visited, &result, 0)
		}
	}

	// reverse result (not sure if best way but works)
	final := []string{}
	for i := len(result) - 1; i >= 0; i-- {
		final = append(final, result[i])
	}

	fmt.Println("final result:", final)
	return final
}
