package leetcode.46

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) == 0 {
		return ret
	}
	if len(nums) == 1{
		return [][]int{nums}
	}

	f := nums[0]
	children := permute(nums[1:])
	for _, v := range children {
		for i := 0; i < len(v) +1; i++ {
			temp := make([]int, 0, len(v)+1)
			temp = append(temp, v[:i]...)
			temp = append(temp, f)
			temp = append(temp, v[i:]...)
			ret = append(ret, temp)
		}
	}
	return ret
}

// @lc code=end
