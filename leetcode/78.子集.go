package leetcode_78

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ret := [][]int{{}}
	for _, v := range nums {
		for _, s := range ret {
			temp := append([]int{}, s...)
			temp = append(temp, v)
			ret = append(ret, temp)
		}
	}
	return ret
}

// @lc code=end
