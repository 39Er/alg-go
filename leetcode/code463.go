package leetcode

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	gridLength := len(grid)
	for i, j := range grid {
		itemLength := len(j)
		for k, l := range j {
			if l == 0 {
				continue
			}

			value := 4

			if i-1 >= 0 && grid[i-1][k] == 1 {
				value--
			}
			if i+1 < gridLength && grid[i+1][k] == 1 {
				value--
			}

			if k-1 >= 0 && j[k-1] == 1 {
				value--
			}

			if k+1 < itemLength && j[k+1] == 1 {
				value--
			}
			perimeter += value

		}
	}
	return perimeter
}

func getCellValue(i, j int, grid [][]int) int {
	if i < 0 || i >= len(grid) {
		return 0
	}
	item := grid[i]
	if j < 0 || j >= len(item) {
		return 0
	}

	return item[j]
}
