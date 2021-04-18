package main

import "fmt"

var n int
var graph [100][100]int

const max = 0x3f3f3f

var visited [100]bool
var dist [100]int

func dijkstra(src int) {
	dist[src] = 0
	for i := 0; i < n; i ++ {
		min := max

		u := -1
		for j := 0; j < n; j ++ {
			if !visited[j] && dist[j] < min {
				min, u = dist[j], j
			}
		}
		visited[u] = true

		for j := 0; j < n; j ++ {
			if !visited[j] && graph[u][j] != 0 && dist[u] != max && dist[u] + graph[u][j] < dist[j] {
				dist[j] = dist[u] + graph[u][j]
			}
		}
	}
}

func main() {
	n = 3
	//for i := 0; i < n; i ++ {
	//	for j := 0; j < n; j ++ {
	//		graph[i][j] = 0
	//	}
	//}
	dijkstra(0)
	fmt.Println(dist[:])
}
