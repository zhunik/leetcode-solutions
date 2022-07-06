package solutions

func numIslands(grid [][]byte) int {
	var islands int
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 49 {
				islands++
				markAllLand(grid, i, j)
			}
		}
	}

	return islands
}

func markAllLand(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 48 {
		return
	}

	grid[i][j] = 48

	markAllLand(grid, i-1, j)
	markAllLand(grid, i+1, j)
	markAllLand(grid, i, j-1)
	markAllLand(grid, i, j+1)
}
