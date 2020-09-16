package leetcode

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	ret := [][]int{}
	if k > n || k < 1{
		return ret
	}
	for i := n; i>=k; i-- {
		childComb := combine(i-1, k-1)
		for _,c := range childComb {
			ret = append(ret, append(c,i))
		}
		if len(childComb) == 0 {
			ret = append(ret, []int{i})
		}
	}
	return ret
}

// @lc code=end
