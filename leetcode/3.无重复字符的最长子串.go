/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max := 0
	left := 0
	cache := map[rune]int{}
	for i, v := range s {
		if li, ok := cache[v]; ok && li >= left {
			left = cache[v] + 1
		}
		cache[v] = i
		cl := i - left + 1
		if cl > max {
			max = cl
		}
	}
	return max
}

// @lc code=end

