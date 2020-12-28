/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ret := 0
	j := 0
	for _, gv := range g {
		for ; j < len(s); j++ {
			if s[j] >= gv {
				ret++
				j++
				break
			}
		}
	}
	return ret
}

// @lc code=end

