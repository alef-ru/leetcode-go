package problems

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
// Memory Usage: 2 MB, less than 87.95% of Go online submissions for Pascal's Triangle.

//  -------------   submission start

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)
	if numRows > 0 {
		rows[0] = []int{1}
		for r := 1; r < numRows; r++ {
			rows[r] = make([]int, r+1)
			rows[r][0] = 1
			for i := 1; i <= r-1; i++ {
				rows[r][i] = rows[r-1][i-1] + rows[r-1][i]
			}
			rows[r][r] = 1
		}
	}

	return rows
}

//  -------------   submission end
