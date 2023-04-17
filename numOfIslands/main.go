package main

func numIslands(grid [][]byte) int {
	if len(grid) < 1 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	count := 0
	for i := 0; i != n; i++ {
		for j := 0; j != m; j++ {
			if grid[i][j] == '1' {
				count++
				explore(grid, i, j, n, m)
			}
		}
	}
	return count
}

func explore(grid [][]byte, x, y, n, m int) {
	if x >= n || y >= m || x < 0 || y < 0 {
		return
	}
	if grid[x][y] == '1' {
		grid[x][y] = '0'
		explore(grid, x+1, y, n, m)
		explore(grid, x-1, y, n, m)
		explore(grid, x, y-1, n, m)
		explore(grid, x, y+1, n, m)
	}
}
