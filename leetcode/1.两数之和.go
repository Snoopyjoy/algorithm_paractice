package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int, len(nums))
	for i, v := range nums {
		v1 := target - v
		if j, ok := cache[v1]; ok {
			return []int{j, i}
		}
		cache[v] = i
	}
	return []int{}
}

// @lc code=end
