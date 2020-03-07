package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ret := &[]string{}
	gen(n, 0, 0, "", ret)
	return *ret
}

func gen(n, left, right int, midR string, ret *[]string) {
	if left == right && left == n {
		*ret = append(*ret, midR)
		return
	}
	if left < n {
		gen(n, left+1, right, midR+"(", ret)
	}
	if right < left {
		gen(n, left, right+1, midR+")", ret)
	}
}

// @lc code=end
