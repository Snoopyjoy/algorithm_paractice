package leetcode.46

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}

	ret := [][]int{}
	v := nums[0]
	childs := permute(nums[1:])
	for _, p := range childs {
		for j:= 0; j < len(nums);j++{
			a := make([]int,0,len(nums))
			a = append(a, p[:j]...)
			a = append(a, v)
			a = append(a, p[j:]...)
			ret = append(ret, a)
		}
	}
	return ret
}

// @lc code=end
