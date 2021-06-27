package _695_max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	var maxArea int = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if area := areaOfIsland(grid, i, j); area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func areaOfIsland(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		if grid[i][j] == 1 {
			grid[i][j] = 0
			return areaOfIsland(grid, i-1, j) + areaOfIsland(grid, i+1, j) +
				areaOfIsland(grid, i, j-1) + areaOfIsland(grid, i, j+1) + 1
		} else {
			return 0
		}
	} else {
		return 0
	}
}
