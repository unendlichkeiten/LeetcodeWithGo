package _934_shortest_bridge

type point struct {
	x int
	y int
}

var directions = []point{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

func shortestBridge(grid [][]int) int {
	// 使用 dfs 遍历第一个岛屿
	r, c := len(grid), len(grid[0])
	skip := false
	queue, level := make([]*point, 0), 0
	for i := 0; i < r; i++ {
		if skip {
			break
		}
		for j := 0; j < c; j++ {
			if grid[i][j] == 1 {
				dfs(grid, &queue, r, c, i, j)
				skip = true
				break
			}
		}
	}

	// 使用 bfs 计算两个岛屿之间的最小值
	x, y := 0, 0
	for len(queue) > 0 {
		level++
		for nPoints := len(queue); nPoints > 0; nPoints-- {
			p := queue[0]
			queue = queue[1:]
			for _, direction := range directions {
				x, y = p.x+direction.x, p.y+direction.y
				if x >= 0 && y >= 0 && x < r && y < c {
					if grid[x][y] == 2 {
						continue
					}
					if grid[x][y] == 1 {
						return level
					}
					queue = append(queue, &point{x, y})
					grid[x][y] = 2
				}
			}
		}
	}
	return level
}

func dfs(grid [][]int, queue *[]*point, r, c, i, j int) {
	if i == r || j == c || i < 0 || j < 0 || grid[i][j] == 2 {
		return
	}

	if grid[i][j] == 0 {
		grid[i][j] = 2
		*queue = append(*queue, &point{i, j})
		return
	}

	grid[i][j] = 2

	dfs(grid, queue, r, c, i+1, j)
	dfs(grid, queue, r, c, i-1, j)
	dfs(grid, queue, r, c, i, j+1)
	dfs(grid, queue, r, c, i, j-1)
}
