package leetcode.22

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {

	ret := []string{}
	var helper func(int,int,string) 
	helper = func(l, r int, s string){
		if l == n && r == n {
			ret = append(ret, s)
			return
		}
		if l < n {
			helper(l+1, r, s + "(")
		}
		if r < l {
			helper(l, r+1, s + ")")
		}
	}
	helper(0,0,"")
	return ret
}
// @lc code=end
