/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	numLen := len( nums )
    if numLen <= 1{
		return numLen
	}
	cursor := 0
	for index := 1; index < numLen; index++ {
		if nums[cursor] != nums[index] {
			cursor++
			if cursor != index {
				nums[cursor] = nums[index]
			}
		}
	}  
	return cursor+1
}
// @lc code=end

