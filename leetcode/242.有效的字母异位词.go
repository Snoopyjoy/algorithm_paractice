package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cache := make(map[rune]int, 26)
	for _, v := range s {
		cache[v]++
	}
	for _, v := range t {
		cache[v]--
	}
	for _, v := range cache {
		if v != 0 {
			return false
		}
	}
	return true

}

// @lc code=end'
