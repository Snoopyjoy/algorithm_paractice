/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	store := make(map[int]struct{}, len(nums))
	for _,v := range nums  {
		if _,ok := store[v]; ok {
			return true
		}
		store[v] = struct{}{}
	}
	return false
}
// @lc code=end

