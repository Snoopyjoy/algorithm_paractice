package leetcode.51
/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	ret := [][]string{}
	angle45 := make(map[int]bool, 2*n)
	angle135 := make(map[int]bool, 2*n)
	colData := make([]bool, n)

	canPut := func(row, col int)bool{
		return !angle45[row+col] && !angle135[row-col] && !colData[col]
	}

	getLine := func(col int)string {
		bts := []byte{}
		for i:=0; i < n;i++ {
			if col == i {
				bts = append(bts, 'Q')
			} else {
				bts = append(bts, '.')
			}
		}
		return string(bts)
	}

	var helper func(row int, grids []string)
	helper = func(row int, grids []string) {
		if row == n {
			ret = append(ret, grids)
			return
		}

		for col:=0; col < n; col++ {
			if !canPut(row, col) {
				continue
			}
			colData[col] = true
			angle45[row+col] = true
			angle135[row-col] = true
			cgrids := make([]string, 0, len(grids)+1)
			cgrids = append(cgrids, grids...)
			cgrids = append(cgrids, getLine(col))
			helper(row+1, cgrids)
			colData[col] = false
			angle45[row+col] = false
			angle135[row-col] = false
		}
	}

	helper(0, []string{})

	return ret
}
// @lc code=end
