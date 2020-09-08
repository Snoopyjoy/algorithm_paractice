package leetcode.51
/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N皇后
 */

// @lc code=start
func fillRow(rows []int, n int) []string {
	ret := make([]string, n)
	for j, idx := range rows {
		dd := make([]byte, n)
		for i:= 0; i < n;i++ {
			if i == idx {
				dd[i] = 'Q'
			} else {
				dd[i] = '.'
			}
		}
		ret[j] =string(dd)
	}
	return ret
}

func solveNQueens(n int) [][]string {
	ret := [][]string{}

	cols := make(map[int]bool, n)
	a45 := make(map[int]bool, 2*n)
	a135 := make(map[int]bool, 2*n)
	rowRst := make([]int, n)

	var canput = func(row, col int) bool {
		return !cols[col] && !a45[row+col] && !a135[row-col]
	}
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ret = append(ret, fillRow(rowRst, n))
		}
		for i := 0; i < n; i++ {
			if !canput(row, i) {
				continue
			}
			cols[i] = true
			a45[row+i] = true
			a135[row-i] = true
			rowRst[row] = i
			dfs(row+1)
			cols[i] = false
			a45[row+i] = false
			a135[row-i] = false
			rowRst[row] = 0
		}
	}
	dfs(0)
	return ret
}

// @lc code=end

