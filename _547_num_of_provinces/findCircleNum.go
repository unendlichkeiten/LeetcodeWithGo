package _547_num_of_provinces

func findCircleNum(isConnected [][]int) int {
	var pCount int = 0
	visited := make([]bool, 0)
	for i := 0; i < len(isConnected); i++ {
		visited = append(visited, false)
	}
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			travelCity(isConnected, i, visited)
			pCount++
		}
	}

	return pCount
}

func travelCity(isConnected [][]int, i int, visted []bool) {
	visted[i] = true
	for j := 0; j < len(isConnected[i]); j++ {
		if isConnected[i][j] == 1 && !visted[j] {
			travelCity(isConnected, j, visted)
		}
	}
}
