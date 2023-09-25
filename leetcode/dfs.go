package leetcode

func GetMaximumGold(grid [][]int) int {
	maxX := len(grid)
	maxY := len(grid[0])
	result := 0
	// dfs
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			if grid[i][j] != 0 {
				cur := dfs(&grid, i, j)
				if result < cur {
					result = cur
				}
			}

		}
	}
	return result
}

func dfs(grid *[][]int, r, c int) int {
	if r < 0 || r >= len(*grid) || c < 0 || c >= len((*grid)[0]) || (*grid)[r][c] == 0 {
		return 0
	}
	sum := (*grid)[r][c]
	val := (*grid)[r][c]
	(*grid)[r][c] = 0
	sum += max(dfs(grid, r - 1, c), dfs(grid, r + 1, c), dfs(grid, r, c - 1), dfs(grid, r, c + 1))
	(*grid)[r][c] = val
	return sum
}

func dfs1(x int, y int, grid *[][]int, maxX, maxY int) int {
	if  x < 0 || x >= maxX || y < 0 || y >= maxY || (*grid)[x][y] == 0{
		return 0
	}
	sum := (*grid)[x][y]
	val := (*grid)[x][y]
	(*grid)[x][y] = 0
	var a, b, c, d int
	// add top
	//if x+1 < maxX && (*grid)[x+1][y] > 0 {
		 a = dfs1 (x+1, y, grid, maxX, maxY)
	//}
	// add bottom
	//if x-1 >= 0 && (*grid)[x-1][y] > 0 {
		b = dfs1 (x-1, y, grid, maxX, maxY)
	//}

	// add left
	//if y-1 >= 0 && (*grid)[x][y-1] > 0 {
		c = dfs1 (x, y-1, grid, maxX, maxY)
	//}

	// add right
	//if y+1 < maxY && (*grid)[x][y+1] > 0 {
		d = dfs1 (x, y+1, grid, maxX, maxY)
	//}
	sum += max(a, b, c, d)
	(*grid)[x][y] = val
	return sum

}

func max(a, b, c, d int) int {
	result := a
	if b > result {
		result = b
	}
	if c > result {
		result = c
	}
	if d > result {
		result = d
	}
	return result
}



