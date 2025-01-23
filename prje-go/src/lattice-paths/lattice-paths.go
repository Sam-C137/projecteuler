package lattice_paths

// Run Starting in the top left corner of a 2 times 2 grid, and only being able to move to the right and down,
// there are exactly 6 routes to the bottom right corner.
// How many such routes are there through a 20 times 20 grid?
func Run(size int) int {
	grid := getGrid(size)

	for r := range grid {
		grid[0][r] = 1
		grid[r][0] = 1
	}

	for r := range grid {
		for c := range grid[r] {
			if r > 0 && c > 0 {
				grid[r][c] = grid[r-1][c] + grid[r][c-1]
			}
		}
	}

	return grid[size][size]
}

func getGrid(size int) [][]int {
	grid := make([][]int, size+1)

	for i := range grid {
		grid[i] = make([]int, size+1)
	}

	return grid
}
